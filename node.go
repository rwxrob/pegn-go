package pegen

import (
	"errors"
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
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
	expression []*pegn.Node
}

func (g *Generator) nodeName(s string) string {
	// Check whether the node has an alias.
	if c, ok := g.config.NodeAliases[s]; ok {
		s = c
	}
	return s
}

func (g *Generator) parseNode(n *pegn.Node) {
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
			g.errors = append(g.errors, errors.New("unknown node child"))
		}
	}
	g.nodes = append(g.nodes, node)
}

func (g *Generator) parseScan(n *pegn.Node) {
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
			g.errors = append(g.errors, errors.New("unknown scan child"))
		}
	}
	g.nodes = append(g.nodes, scan)
}

func (g *Generator) generateNodes() {
	//  Write to the ast buffer.
	w := g.writers["ast"]
	for idx, node := range g.nodes {
		w.wlnf("func %s(p *ast.Parser) (*ast.Node, error) {", g.nodeName(node.name))
		{
			w := w.indent()
			w.wln("return p.Expect(")
			if node.scan {
				g.generateExpression(w.indent(), node.expression)
			} else {
				{
					w := w.indent()
					w.wln("ast.Capture{")
					{
						w := w.indent()
						w.wlnf("Type: %s,", g.typeName(g.nodeName(node.name)))
						w.wln("Value: ")
						g.generateExpression(w, node.expression)
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
}

func (g *Generator) generateExpression(w *writer, expression []*pegn.Node) {
	size := len(expression)
	if 1 < size {
		w.wln("op.Or{")
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
				g.generateSequence(w, n.Children())

			default:
				g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
				continue
			}
		}
	}

	if 1 < size {
		w.wln("},")
	}
}

func (g *Generator) generateSequence(w *writer, sequence []*pegn.Node) {
	size := len(sequence)
	if 1 < size {
		w.wln("op.And{")
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
				var quant *pegn.Node
				switch last := n.Children()[len(n.Children())-1]; last.Type {
				case nd.Optional, nd.MinZero, nd.MinOne, nd.MinMax, nd.Count:
					quant = last
				}
				if quant == nil {
					g.generatePrimary(w, n.Children()[0])
					break
				} else {
					q := n.Children()[1]
					n := n.Children()[0]

					switch q.Type {
					case nd.Optional:
						w.wln("op.Optional(")
						g.generatePrimary(w.indent(), n)
						w.wln("),")
					case nd.MinZero:
						w.wln("op.MinZero(")
						g.generatePrimary(w.indent(), n)
						w.wln("),")
					case nd.MinOne:
						w.wln("op.MinOne(")
						g.generatePrimary(w.indent(), n)
						w.wln("),")
					case nd.MinMax:
						min := q.Children()[0].Value
						max := q.Children()[1].Value
						w.wlnf("op.MinMax(%s, %s,", min, max)
						g.generatePrimary(w.indent(), n)
						w.wln("),")
					case nd.Count:
						count := q.Value
						w.wlnf("op.Repeat(%s,", count)
						g.generatePrimary(w.indent(), n)
						w.wln("),")
					default:
						g.errors = append(g.errors, fmt.Errorf("unknown quant child: %v", n.Types[n.Type]))
						continue
					}
				}
			// PosLook <-- '&' Primary Quant?
			case nd.PosLook:
				g.errors = append(g.errors, errors.New("unsupported: "+n.Types[n.Type]))
				continue
			// NegLook <-- '!' Primary Quant?
			case nd.NegLook:
				w.wln("op.Not{")
				g.generatePrimary(w.indent(), n.Children()[0])
				w.wln("},")
			default:
				g.errors = append(g.errors, fmt.Errorf("unknown sequence child: %v", n.Types[n.Type]))
				continue
			}
		}
	}

	if 1 < size {
		w.wln("},")
	}
}

// Primary <- Simple / CheckId / '(' Expression ')'
func (g *Generator) generatePrimary(w *writer, n *pegn.Node) {
	switch n.Type {
	case nd.Comment, nd.EndLine:
		// Ignore these.
	case nd.Unicode, nd.Hexadec:
		v, _ := ConvertToRuneString(n.Value[1:], 16)
		w.w(v)
	case nd.Binary:
		v, _ := ConvertToRuneString(n.Value[1:], 2)
		w.w(v)
	case nd.Octal:
		v, _ := ConvertToRuneString(n.Value[1:], 8)
		w.w(v)
	case nd.ClassId:
		w.w(g.className(g.GetID(n)))
	case nd.TokenId:
		id := g.tokenName(g.GetID(n))
		tk := g.tokens.get(id)
		if tk.isString() {
			g.errors = append(g.errors, errors.New("token value is a string"))
			break
		}
		w.w(id)
	case nd.AlphaRange:
		// AlphaRange <-- '[' Letter '-' Letter ']'
		min := n.Children()[0].Value
		max := n.Children()[1].Value
		w.wf("parser.CheckRuneRange('%s', '%s')", min, max)
	case nd.IntRange:
		// IntRange <-- '[' Integer '-' Integer ']'
		min, _ := strconv.Atoi(n.Children()[0].Value)
		max, _ := strconv.Atoi(n.Children()[1].Value)
		if min < 0 {
			g.errors = append(g.errors, fmt.Errorf("int range is negative: [%v-%v]", min, max))
			break
		}
		if max <= min {
			g.errors = append(g.errors, fmt.Errorf("int range is inverted: [%v-%v]", min, max))
			break
		}
		if 10 <= max {
			g.errors = append(g.errors, fmt.Errorf("int range too large: [%v-%v]", min, max))
			break
		}
		w.wf("parser.CheckRuneRange('%d', '%d')", min, max)
	case nd.UniRange, nd.HexRange:
		// UniRange <-- '[' Unicode '-' Unicode ']'
		// HexRange <-- '[' Hexadec '-' Hexadec ']'
		min, _ := ConvertToRuneString(n.Children()[0].Value[1:], 16)
		max, _ := ConvertToRuneString(n.Children()[1].Value[1:], 16)
		w.wf("parser.CheckRuneRange(%s, %s)", min, max)
	case nd.BinRange:
		// BinRange <-- '[' Binary '-' Binary ']'
		min, _ := ConvertToRuneString(n.Children()[0].Value[1:], 2)
		max, _ := ConvertToRuneString(n.Children()[1].Value[1:], 2)
		w.wf("parser.CheckRuneRange(%s, %s)", min, max)
	case nd.OctRange:
		// OctRange <-- '[' Octal '-' Octal ']'
		min, _ := ConvertToRuneString(n.Children()[0].Value[1:], 8)
		max, _ := ConvertToRuneString(n.Children()[1].Value[1:], 8)
		w.wf("parser.CheckRuneRange(%s, %s)", min, max)
	case nd.String:
		w.wf("%q", n.Value)
	case nd.CheckId:
		w.w(g.nodeName(g.GetID(n)))
	case nd.Expression:
		g.generateExpression(w, n.Children())
	default:
		g.errors = append(g.errors, fmt.Errorf("unknown plain child: %v", n.Types[n.Type]))
	}

	if n.Type != nd.Expression {
		w.noIndent().wln(",")
	}
}
