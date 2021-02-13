package pegn

import (
	"errors"
	"fmt"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn/nd"
)

type nodes []node

func (ns nodes) get(id string) node {
	for _, node := range ns {
		if node.name == id {
			return node
		}
	}
	return node{}
}

// node is a simplified representation of a NodeDef or ScanDef.
//
// PEGN:
//	NodeDef <-- CheckId SP+ '<--' SP+ Expression
//	ScanDef <-- CheckId SP+ '<-'  SP+ Expression
type node struct {
	name       string
	scan       bool
	expression []*ast.Node
}

// nodeName returns a formatted node name.
func (g *generator) nodeName(s string) string {
	// Check whether the node has an alias.
	if c, ok := g.config.NodeAliases[s]; ok {
		s = c
	}
	return s
}

// parseNode parses the given node as a significant node and adds it to the list
// of nodes within the generator.
func (g *generator) parseNode(n *ast.Node) error {
	var node node
	for _, n := range n.Children() {
		switch n.Type {
		case nd.Comment, nd.EndLine:
			// Ignore these.
			continue
		case nd.CheckId:
			node.name = n.Value

		// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
		case nd.Expression:
			node.expression = n.Children()
		default:
			return errors.New("unknown node child")
		}
	}
	g.nodes = append(g.nodes, node)
	return nil
}

// parseScan parses the given node as a insignificant node and adds it to the
// list of nodes within the generator.
func (g *generator) parseScan(n *ast.Node) error {
	scan := node{
		scan: true,
	}
	for _, n := range n.Children() {
		switch n.Type {
		case nd.Comment, nd.EndLine:
			// Ignore these.
			continue
		case nd.CheckId:
			scan.name = n.Value
		case nd.Expression:
			// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
			scan.expression = n.Children()
		default:
			return errors.New("unknown scan child")
		}
	}
	g.nodes = append(g.nodes, scan)
	return nil
}

// generateNodes writes all the nodes to the given writer.
func (g *generator) generateNodes(w *writer) error {
	for idx, node := range g.nodes {
		w.wlnf("func %s(p *ast.Parser) (*ast.Node, error) {", g.nodeName(node.name))
		{
			w := w.indent()
			w.wln("return p.Expect(")
			if node.scan {
				if err := g.generateExpression(w.indent(), node.expression, true); err != nil {
					return err
				}
			} else {
				{
					w := w.indent()
					w.wln("ast.Capture{")
					{
						w := w.indent()
						w.wlnf("Type:        %s,", g.typeNameGenerated(g.nodeName(node.name)))
						w.wlnf("TypeStrings: %s,", g.typePrefix("NodeTypes"))
						w.w("Value: ")
						if len(node.expression) == 1 {
							if singleNestedValue(node.expression[0]) {
								w.noIndent().w("      ") // To align with Type(Strings)
							}
						}
						if err := g.generateExpression(w, node.expression, false); err != nil {
							return err
						}
					}
					w.wln("},")
				}
			}
			w.wln(")")
		}
		w.wlnf("}")
		if idx != len(g.nodes)-1 {
			w.ln()
		}
	}

	for _, dep := range g.dependencies {
		if err := dep.generateNodes(w); err != nil {
			return err
		}
	}

	return nil
}

// generateExpression is responsible for generating expressions.
func (g *generator) generateExpression(w *writer, expression []*ast.Node, indent bool) error {
	size := len(expression)
	if 1 < size {
		if !indent {
			w.noIndent().wln("op.Or{")
			indent = true
		} else {
			w.wln("op.Or{")
		}
	}

	{
		w := w
		if 1 < size {
			w = w.indent()
		}

		for _, n := range expression {
			switch n.Type {
			case nd.Comment, nd.EndLine:
				// Ignore these.
				continue

			// Sequence <-- Rule (Spacing Rule)*
			case nd.Sequence:
				if err := g.generateSequence(w, n.Children(), indent); err != nil {
					return err
				}

			default:
				return fmt.Errorf("unknown expression child: %v", nd.NodeTypes[n.Type])
			}
		}
	}

	if 1 < size {
		w.wln("},")
	}
	return nil
}

// generateSequence is responsible for generating sequences.
func (g *generator) generateSequence(w *writer, sequence []*ast.Node, indent bool) error {
	size := len(sequence)
	if 1 < size {
		if !indent {
			w.noIndent().wln("op.And{")
			indent = true
		} else {
			w.wln("op.And{")
		}
	}

	{
		w := w
		if 1 < size {
			w = w.indent()
		}

		for _, n := range sequence {
			switch n.Type {
			case nd.Comment, nd.EndLine:
				// Ignore these.
				continue
			// Plain <-- Primary Quant?
			case nd.Plain:
				// Plain <-- Primary Quant?
				var quant *ast.Node
				switch last := n.Children()[len(n.Children())-1]; last.Type {
				case nd.Optional, nd.MinZero, nd.MinOne, nd.MinMax, nd.Count:
					quant = last
				}
				if quant == nil {
					if err := g.generatePrimary(w, n.Children()[0], indent); err != nil {
						return err
					}
					break
				} else {
					q := n.Children()[1]
					n := n.Children()[0]

					switch q.Type {
					case nd.Optional:
						if !indent {
							w.noIndent().wln("op.Optional(")
							indent = true
						} else {
							w.wln("op.Optional(")
						}
						if err := g.generatePrimary(w.indent(), n, indent); err != nil {
							return err
						}
						w.wln("),")
					case nd.MinZero:
						if !indent {
							w.noIndent().wln("op.MinZero(")
							indent = true
						} else {
							w.wln("op.MinZero(")
						}
						if err := g.generatePrimary(w.indent(), n, indent); err != nil {
							return err
						}
						w.wln("),")
					case nd.MinOne:
						if !indent {
							w.noIndent().wln("op.MinOne(")
							indent = true
						} else {
							w.wln("op.MinOne(")
						}
						if err := g.generatePrimary(w.indent(), n, indent); err != nil {
							return err
						}
						w.wln("),")
					case nd.MinMax:
						min := q.Children()[0].Value
						max := q.Children()[1].Value
						if !indent {
							w.noIndent().wlnf("op.MinMax(%s, %s,", min, max)
							indent = true
						} else {
							w.wlnf("op.MinMax(%s, %s,", min, max)
						}
						if err := g.generatePrimary(w.indent(), n, indent); err != nil {
							return err
						}
						w.wln("),")
					case nd.Count:
						min := q.Value
						if !indent {
							w.noIndent().wlnf("op.Repeat(%s,", min)
							indent = true
						} else {
							w.wlnf("op.Repeat(%s,", min)
						}
						if err := g.generatePrimary(w.indent(), n, indent); err != nil {
							return err
						}
						w.wln("),")
					default:
						return fmt.Errorf("unknown quant child: %v", nd.NodeTypes[n.Type])
					}
				}
			// PosLook <-- '&' Primary Quant?
			case nd.PosLook:
				if !indent {
					w.noIndent().wln("op.Ensure{")
					indent = true
				} else {
					w.wln("op.Ensure{")
				}
				if err := g.generatePrimary(w.indent(), n.Children()[0], indent); err != nil {
					return err
				}
				w.wln("},")
			// NegLook <-- '!' Primary Quant?
			case nd.NegLook:
				if !indent {
					w.noIndent().wln("op.Not{")
					indent = true
				} else {
					w.wln("op.Not{")
				}
				if err := g.generatePrimary(w.indent(), n.Children()[0], indent); err != nil {
					return err
				}
				w.wln("},")
			default:
				return fmt.Errorf("unknown sequence child: %v", nd.NodeTypes[n.Type])
			}
		}
	}

	if 1 < size {
		w.wln("},")
	}
	return nil
}
