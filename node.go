package pegen

import (
	"errors"
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
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

func (g *Generator) parseNode(n *pegn.Node) {
	var node node
	for _, n := range n.Children() {
		switch n.Type {
		case nd.CheckId:
			node.name = n.Value
		case nd.Expression:
			// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
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
	w := g.writers["ast"]
	for _, node := range g.nodes {
		w.wlnf("func %s(p *pegn.Parser) (*pegn.Node, error) {", node.name)
		{
			w := w.indent()
			w.wlnf("node := pegn.NewNode(%s, NodeTypes)", g.typeName(node.name))

			var seqCount int
			for _, n := range node.expression {
				if n.Type == nd.Sequence {
					seqCount++
				}
			}

			for _, n := range node.expression {
				// Sequence <-- Rule (Spacing Rule)*
				switch n.Type {
				case nd.Comment, nd.EndLine:
					// Ignore these.
				case nd.Sequence:
					if seqCount > 1 {
						w.wln("if n, err := func() (*pegn.Node, error) {")
					}

					{
						w := w
						if seqCount > 1 {
							w = w.indent()
						}

						w.wln("var (")
						{
							w := w.indent()
							if seqCount > 1 {
								w.wlnf("node = pegn.NewNode(%s, NodeTypes)", g.typeName(node.name))
								w.ln()
							}
							w.wln("err error")
							w.wln("n   *pegn.Node")
						}
						w.wln(")")
						w.wln("_ = err")
						w.wln("_ = n")
						w.ln()

						// Sequence <-- Rule (Spacing Rule)*
						// Rule      <- PosLook / NegLook / Plain
						for _, n := range n.Children() {
							switch n.Type {
							case nd.EndLine:
								// Ignore this.
							case nd.Plain:
								// Plain <-- Primary Quant?
								if len(n.Children()) == 1 {
									switch n := n.Children()[0]; n.Type {
									case nd.Unicode, nd.Hexadec:
										value, _ := ConvertToRuneString(n.Value[1:], 16)
										w.wlnf("if _, err = p.Expect(%s); err != nil {", value)
										w.indent().wlnf("return expected(%q, p)", n.Value)
										w.wln("}")
									case nd.Binary:
										value, _ := ConvertToRuneString(n.Value[1:], 2)
										w.wlnf("if _, err = p.Expect(%s); err != nil {", value)
										w.indent().wlnf("return expected(%q, p)", n.Value)
										w.wln("}")
									case nd.Octal:
										value, _ := ConvertToRuneString(n.Value[1:], 8)
										w.wlnf("if _, err = p.Expect(%s); err != nil {", value)
										w.indent().wlnf("return expected(%q, p)", n.Value)
										w.wln("}")
									case nd.ClassId:
										id := g.GetID(n)
										w.wlnf("if _, err = p.Expect(%s); err != nil {", g.className(id))
										w.indent().wlnf("return expected(%q, p)", id)
										w.wln("}")
									case nd.TokenId:
										id := g.GetID(n)
										w.wlnf("if _, err = p.Expect(%s); err != nil {", g.tokenName(id))
										w.indent().wlnf("return expected(%q, p)", id)
										w.wln("}")
									case nd.AlphaRange, nd.IntRange,
										nd.UniRange, nd.BinRange,
										nd.HexRange, nd.OctRange:
										g.errors = append(g.errors, errors.New("ranges not supported"))
										continue
									case nd.String:
										w.wlnf("if _, err = p.Expect(%q); err != nil {", n.Value)
										w.indent().wlnf("return expected(%q, p)", n.Value)
										w.wln("}")
									case nd.CheckId:
										w.wlnf("n, err = %s(p)", n.Value)
										w.wln("if err != nil {")
										w.indent().wlnf("return expected(%q, p)", n.Value)
										w.wln("}")
										if g.nodes.get(n.Value).scan {
											w.wln("node.AdoptFrom(n)")
										} else {
											w.wln("node.AppendChild(n)")
										}
									case nd.Expression:
										// TODO
									default:
										g.errors = append(g.errors, fmt.Errorf("unknown plain child: %v", n.Types[n.Type]))
										continue
									}
								} else {
									// TODO: Quant
								}

							case nd.PosLook, nd.NegLook:
								// PosLook <-- '&' Primary Quant?
								// NegLook <-- '!' Primary Quant?
								g.errors = append(g.errors, errors.New("unsupported: "+n.Types[n.Type]))
							default:
								g.errors = append(g.errors, fmt.Errorf("unknown sequence child: %v", n.Types[n.Type]))
							}
							w.ln()
						}

						if seqCount > 1 {
							w.wln("return node, nil")
						}
					}

					if seqCount > 1 {
						w.wln("}(); err == nil {")
						{
							w := w.indent()
							w.wln("node.AdoptFrom(n)")
							w.wln("return node, nil")
						}
						w.wln("}")
						w.ln()
					}
				default:
					g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
				}
			}

			if seqCount > 1 {
				w.wln("return nil, nil")
			} else {
				w.wln("return node, nil")
			}
		}
		w.wln("}")
		w.ln()
	}

	w.wln("func expected(value string, p *pegn.Parser) (*pegn.Node, error) {")
	w.indent().wln("return nil, fmt.Errorf(\"expected %v at %v\", value, p.Mark())")
	w.wln("}")
}
