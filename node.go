package pegen

import (
	"errors"
	"fmt"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn"
	"strconv"
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

type node struct {
	name       string
	scan       bool
	expression []*ast.Node
}

func (g *Generator) nodeName(s string) string {
	// Check whether the node has an alias.
	if c, ok := g.config.NodeAliases[s]; ok {
		s = c
	}
	return s
}

func (g *Generator) parseNode(n *ast.Node) error {
	var node node
	for _, n := range n.Children() {
		switch n.Type {
		case pegn.CommentType, pegn.EndLineType:
			// Ignore these.
			continue
		case pegn.CheckIdType:
			node.name = n.ValueString()

		// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
		case pegn.ExpressionType:
			node.expression = n.Children()
		default:
			return errors.New("unknown node child")
		}
	}
	g.nodes = append(g.nodes, node)
	return nil
}

func (g *Generator) parseScan(n *ast.Node) error {
	scan := node{
		scan: true,
	}
	for _, n := range n.Children() {
		switch n.Type {
		case pegn.CommentType, pegn.EndLineType:
			// Ignore these.
			continue
		case pegn.CheckIdType:
			scan.name = n.ValueString()
		case pegn.ExpressionType:
			// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
			scan.expression = n.Children()
		default:
			return errors.New("unknown scan child")
		}
	}
	g.nodes = append(g.nodes, scan)
	return nil
}

func (g *Generator) generateNodes() error {
	//  Write to the ast buffer.
	w := g.writers["ast"]
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
						w.w("Type: ")
						if len(node.expression) == 1 {
							if singleNestedValue(node.expression[0]) {
								w.noIndent().w(" ") // To align with 'Value: '
							}
						}
						w.noIndent().wlnf("%s,", g.typeName(g.nodeName(node.name)))
						w.w("Value: ")
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
	return nil
}

func (g *Generator) generateExpression(w *writer, expression []*ast.Node, indent bool) error {
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
			case pegn.CommentType, pegn.EndLineType:
				// Ignore these.
				continue

			// Sequence <-- Rule (Spacing Rule)*
			case pegn.SequenceType:
				if err := g.generateSequence(w, n.Children(), indent); err != nil {
					return err
				}

			default:
				return fmt.Errorf("unknown expression child: %v", pegn.NodeTypes[n.Type])
			}
		}
	}

	if 1 < size {
		w.wln("},")
	}
	return nil
}

func (g *Generator) generateSequence(w *writer, sequence []*ast.Node, indent bool) error {
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
			case pegn.CommentType, pegn.EndLineType:
				// Ignore these.
				continue
			// Plain <-- Primary Quant?
			case pegn.PlainType:
				// Plain <-- Primary Quant?
				var quant *ast.Node
				switch last := n.Children()[len(n.Children())-1]; last.Type {
				case pegn.OptionalType, pegn.MinZeroType, pegn.MinOneType, pegn.MinMaxType, pegn.CountType:
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
					case pegn.OptionalType:
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
					case pegn.MinZeroType:
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
					case pegn.MinOneType:
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
					case pegn.MinMaxType:
						min := q.Children()[0].ValueString()
						max := q.Children()[1].ValueString()
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
					case pegn.CountType:
						min := q.Children()[0].ValueString()
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
						return fmt.Errorf("unknown quant child: %v", pegn.NodeTypes[n.Type])
					}
				}
			// PosLook <-- '&' Primary Quant?
			case pegn.PosLookType:
				return fmt.Errorf("unsupported: %s", pegn.NodeTypes[n.Type])
			// NegLook <-- '!' Primary Quant?
			case pegn.NegLookType:
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
				return fmt.Errorf("unknown sequence child: %v", pegn.NodeTypes[n.Type])
			}
		}
	}

	if 1 < size {
		w.wln("},")
	}
	return nil
}

// Primary <- Simple / CheckId / '(' Expression ')'
func (g *Generator) generatePrimary(w *writer, n *ast.Node, indent bool) error {
	if !indent {
		w = w.noIndent()
		indent = true
	}

	switch n.Type {
	case pegn.CommentType, pegn.EndLineType:
		// Ignore these.
	case pegn.UnicodeType, pegn.HexadecimalType:
		v, _ := ConvertToRuneString(n.ValueString()[1:], 16)
		w.w(v)
	case pegn.BinaryType:
		v, _ := ConvertToRuneString(n.ValueString()[1:], 2)
		w.w(v)
	case pegn.OctalType:
		v, _ := ConvertToRuneString(n.ValueString()[1:], 8)
		w.w(v)
	case pegn.ClassIdType, pegn.ResClassIdType,
		pegn.TokenIdType, pegn.ResTokenIdType,
		pegn.CheckIdType:
		id, err := g.GetID(n)
		if err != nil {
			return err
		}
		w.w(id)
	case pegn.AlphaRangeType:
		// AlphaRange <-- '[' Letter '-' Letter ']'
		min := n.Children()[0].Value
		max := n.Children()[1].Value
		w.wf("parser.CheckRuneRange('%s', '%s')", min, max)
	case pegn.IntRangeType:
		// IntRange <-- '[' Integer '-' Integer ']'
		min, _ := strconv.Atoi(n.Children()[0].ValueString())
		max, _ := strconv.Atoi(n.Children()[1].ValueString())
		if min < 0 {
			return fmt.Errorf("int range is negative: [%v-%v]", min, max)
		}
		if max <= min {
			return fmt.Errorf("int range is inverted: [%v-%v]", min, max)
		}
		if 10 <= max {
			return fmt.Errorf("int range too large: [%v-%v]", min, max)
		}
		w.wf("parser.CheckRuneRange('%d', '%d')", min, max)
	case pegn.UniRangeType, pegn.HexRangeType:
		// UniRange <-- '[' Unicode '-' Unicode ']'
		// HexRange <-- '[' Hexadec '-' Hexadec ']'
		min, _ := ConvertToRuneString(n.Children()[0].ValueString()[1:], 16)
		max, _ := ConvertToRuneString(n.Children()[1].ValueString()[1:], 16)
		w.wf("parser.CheckRuneRange(%s, %s)", min, max)
	case pegn.BinRangeType:
		// BinRange <-- '[' Binary '-' Binary ']'
		min, _ := ConvertToRuneString(n.Children()[0].ValueString()[1:], 2)
		max, _ := ConvertToRuneString(n.Children()[1].ValueString()[1:], 2)
		w.wf("parser.CheckRuneRange(%s, %s)", min, max)
	case pegn.OctRangeType:
		// OctRange <-- '[' Octal '-' Octal ']'
		min, _ := ConvertToRuneString(n.Children()[0].ValueString()[1:], 8)
		max, _ := ConvertToRuneString(n.Children()[1].ValueString()[1:], 8)
		w.wf("parser.CheckRuneRange(%s, %s)", min, max)
	case pegn.StringType:
		w.wf("%q", n.Value)
	case pegn.ExpressionType:
		if err := g.generateExpression(w, n.Children(), indent); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown plain child: %v", pegn.NodeTypes[n.Type])
	}

	if n.Type != pegn.ExpressionType {
		w.noIndent().wln(",")
	}
	return nil
}
