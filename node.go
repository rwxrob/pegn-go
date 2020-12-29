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
	for _, node := range g.nodes {
		w.wlnf("func %s(p *pegn.Parser) (*pegn.Node, error) {", node.name)
		{
			w := w.indent()
			w.wln("var (")
			{
				w := w.indent()
				w.wlnf("node = pegn.NewNode(%s, NodeTypes)", g.typeName(node.name))
				w.ln()
				w.wln("err error")
				w.wln("n   *pegn.Node")
			}
			w.wln(")")
			w.wln("_ = err")
			w.wln("_ = n")
			w.ln()

			// Indicates whether there are one or more possible sequences. If
			// there is only one sequence then the parser has to success since
			// there are no other alternatives.
			var count int
			for _, n := range node.expression {
				if n.Type == nd.Sequence {
					count++
				}
			}

			// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
			// Spacing     <- ComEndLine? SP+
			// ComEndLine  <- SP* ('# ' Comment)? EndLine
			for _, n := range node.expression {
				switch n.Type {
				case nd.Comment, nd.EndLine:
					// Ignore these.
					continue

				// Sequence <-- Rule (Spacing Rule)*
				case nd.Sequence:
					g.generateSequence(w, node, n, count)

				default:
					g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
				}
			}

			if count > 1 {
				w.wln("return nil, err")
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

func (g *Generator) generateSequence(w *writer, node node, n *pegn.Node, seqCount int) {
	if seqCount > 1 {
		w.wln("n, err = func() (*pegn.Node, error) {")
	}
	defer func() {
		if seqCount > 1 {
			w.wln("}()")
			w.wln("if err == nil {")
			{
				w := w.indent()
				w.wln("node.AdoptFrom(n)")
				w.wln("return node, nil")
			}
			w.wln("}")
			w.ln()
		}
	}()

	{
		w := w
		if seqCount > 1 {
			w = w.indent()
		}

		if seqCount > 1 {
			w.wln("var (")
			{
				w := w.indent()
				w.wlnf("node = pegn.NewNode(%s, NodeTypes)", g.typeName(node.name))
				w.ln()
				w.wln("err error")
				w.wln("n   *pegn.Node")
			}
			w.wln(")")
			w.wln("_ = err")
			w.wln("_ = n")
			w.ln()
		}

		// Sequence <-- Rule (Spacing Rule)*
		// Rule      <- PosLook / NegLook / Plain
		for _, n := range n.Children() {
			switch n.Type {
			case nd.EndLine:
				// Ignore this.
				continue
			case nd.Plain:
				// Plain <-- Primary Quant?
				var quant *pegn.Node
				switch last := n.Children()[len(n.Children())-1]; last.Type {
				case nd.Optional, nd.MinZero, nd.MinOne, nd.MinMax, nd.Count:
					quant = last
				}

				if quant == nil {
					g.generateSinglePlain(w, node, n.Children()[0])
					break
				}
				q := n.Children()[1]
				n := n.Children()[0]
				switch q.Type {
				case nd.Optional:
					switch n.Type {
					case nd.Unicode, nd.Hexadec:
						value, _ := ConvertToRuneString(n.Value[1:], 16)
						w.wlnf("_, _ = p.Expect(%s)", value)
					case nd.Binary:
						value, _ := ConvertToRuneString(n.Value[1:], 2)
						w.wlnf("_, _ = p.Expect(%s)", value)
					case nd.Octal:
						value, _ := ConvertToRuneString(n.Value[1:], 8)
						w.wlnf("_, _ = p.Expect(%s)", value)

					case nd.ClassId:
						id := g.GetID(n)
						w.wlnf("_, _ = p.Expect(%s)", g.className(id))

					case nd.TokenId:
						id := g.GetID(n)
						w.wlnf("_, _ = p.Expect(%s)", g.tokenName(id))

					case nd.AlphaRange, nd.IntRange,
						nd.UniRange, nd.BinRange,
						nd.HexRange, nd.OctRange:
						g.errors = append(g.errors, errors.New("ranges not supported"))
						continue

					case nd.String:
						w.wlnf("_, _ = p.Expect(%s)", n.Value)

					case nd.CheckId:
						w.cf("%s?", n.Value)
						w.wlnf("n, err = %s(p)", n.Value)
						w.wln("if err == nil {")
						if g.nodes.get(n.Value).scan {
							w.indent().wln("node.AdoptFrom(n)")
						} else {
							w.indent().wln("node.AppendChild(n)")
						}
						w.wln("}")
					case nd.Expression:
						var count int
						for _, n := range n.Children() {
							if n.Type == nd.Sequence {
								count++
							}
						}

						w.wln("_, err = func() (*pegn.Node, error) {")
						{
							w := w.indent()
							w.wln("var (")
							{
								w := w.indent()
								w.wlnf("node = pegn.NewNode(%s, NodeTypes)", g.typeName(node.name))
								w.ln()
								w.wln("err error")
								w.wln("n   *pegn.Node")
							}
							w.wln(")")
							w.ln()

							// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
							// Spacing     <- ComEndLine? SP+
							// ComEndLine  <- SP* ('# ' Comment)? EndLine
							for _, n := range n.Children() {
								switch n.Type {
								case nd.Comment, nd.EndLine:
									// Ignore these.
									continue

								// Sequence <-- Rule (Spacing Rule)*
								case nd.Sequence:
									g.generateSequence(w, node, n, count)

								default:
									g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
								}
							}

							w.wln("return node, nil")
						}
						w.wln("}()")
						w.wln("if err == nil {")
						w.indent().wln("node.AdoptFrom(n)")
						w.wln("}")
					default:
						g.errors = append(g.errors, fmt.Errorf("unknown plain child: %v", n.Types[n.Type]))
						continue
					}
				case nd.MinZero, nd.MinOne, nd.MinMax, nd.Count:
					w.wln("{")
					{
						w := w.indent()
						w.wln("var count int")
						if n.Type == nd.Expression {
							var count int
							for _, n := range n.Children() {
								if n.Type == nd.Sequence {
									count++
								}
							}

							w.wln("exp := func() (*pegn.Node, error) {")
							{
								w := w.indent()
								w.wln("var (")
								{
									w := w.indent()
									w.wlnf("node = pegn.NewNode(%s, NodeTypes)", g.typeName(node.name))
									w.ln()
									w.wln("err error")
									w.wln("n   *pegn.Node")
								}
								w.wln(")")
								w.wln("_ = err")
								w.wln("_ = n")
								w.ln()

								// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
								// Spacing     <- ComEndLine? SP+
								// ComEndLine  <- SP* ('# ' Comment)? EndLine
								for _, n := range n.Children() {
									switch n.Type {
									case nd.Comment, nd.EndLine:
										// Ignore these.
										continue

									// Sequence <-- Rule (Spacing Rule)*
									case nd.Sequence:
										g.generateSequence(w, node, n, count)

									default:
										g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
									}
								}

								w.wln("return node, nil")
							}
							w.wln("}")
						}
						var condition string
						if q.Type == nd.Count {
							condition = fmt.Sprintf("count < %s", q.Value)
						}
						w.wlnf("for beg := p.Mark(); %s; count++ {", condition)
						{
							w := w.indent()
							switch n.Type {
							case nd.Unicode, nd.Hexadec:
								value, _ := ConvertToRuneString(n.Value[1:], 16)
								w.wlnf("if _, err = p.Expect(%s); err != nil {", value)
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")
							case nd.Binary:
								value, _ := ConvertToRuneString(n.Value[1:], 2)
								w.wlnf("if _, err = p.Expect(%s); err != nil {", value)
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")
							case nd.Octal:
								value, _ := ConvertToRuneString(n.Value[1:], 8)
								w.wlnf("if _, err = p.Expect(%s); err != nil {", value)
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")

							case nd.ClassId:
								id := g.GetID(n)
								w.wlnf("if _, err = p.Expect(%s); err != nil {", g.className(id))
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")

							case nd.TokenId:
								id := g.GetID(n)
								w.wlnf("if _, err = p.Expect(%s); err != nil {", g.tokenName(id))
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")

							case nd.AlphaRange, nd.IntRange,
								nd.UniRange, nd.BinRange,
								nd.HexRange, nd.OctRange:
								g.errors = append(g.errors, errors.New("ranges not supported"))

							case nd.String:
								w.wlnf("if _, err = p.Expect(%q); err != nil {", n.Value)
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")

							case nd.CheckId:
								w.wlnf("n, err = %s(p)", n.Value)
								w.wln("if err == nil {")
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")
								if g.nodes.get(n.Value).scan {
									w.wln("node.AdoptFrom(n)")
								} else {
									w.wln("node.AppendChild(n)")
								}
							case nd.Expression:
								w.wln("n, err = exp()")
								w.wln("if err != nil {")
								{
									w := w.indent()
									w.wln("p.Goto(beg)")
									w.wln("break")
								}
								w.wln("}")
							default:
								g.errors = append(g.errors, fmt.Errorf("unknown plain child: %v", n.Types[n.Type]))
								continue
							}
						}
						w.wln("}")
						switch q.Type {
						case nd.MinZero:
							// Nothing to do here.
						case nd.MinOne:
							w.wln("if count < 1 {")
							w.indent().c("TODO")
							w.indent().wlnf("return expected(%q, p)", "")
							w.wln("}")
						case nd.MinMax:
							min := q.Children()[0].Value
							max := q.Children()[1].Value
							w.wlnf("if count < %v || %v < count {", min, max)
							w.indent().c("TODO")
							w.indent().wlnf("return expected(%q, p)", "")
							w.wln("}")
						case nd.Count:
							w.wlnf("if count != %s {", q.Value)
							w.indent().c("TODO")
							w.indent().wlnf("return expected(%q, p)", "")
							w.wln("}")
						default:
							w.c("TODO")
							fmt.Println(node.name, q)
						}
					}
					w.wln("}")
				default:
					g.errors = append(g.errors, fmt.Errorf("unknown quant child: %v", n.Types[n.Type]))
					continue
				}

			// PosLook <-- '&' Primary Quant?
			case nd.PosLook:
				g.errors = append(g.errors, errors.New("unsupported: "+n.Types[n.Type]))
				continue
			// NegLook <-- '!' Primary Quant?
			case nd.NegLook:
				g.errors = append(g.errors, errors.New("unsupported: "+n.Types[n.Type]))
				continue
			default:
				g.errors = append(g.errors, fmt.Errorf("unknown sequence child: %v", n.Types[n.Type]))
				continue
			}
			w.ln()
		}

		if seqCount > 1 {
			w.wln("return node, nil")
		}
	}
}
func (g *Generator) generateSinglePlain(w *writer, node node, n *pegn.Node) {
	switch n.Type {
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
		var count int
		for _, n := range n.Children() {
			if n.Type == nd.Sequence {
				count++
			}
		}

		// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
		// Spacing     <- ComEndLine? SP+
		// ComEndLine  <- SP* ('# ' Comment)? EndLine
		for _, n := range n.Children() {
			switch n.Type {
			case nd.Comment, nd.EndLine:
				// Ignore these.
				continue

			// Sequence <-- Rule (Spacing Rule)*
			case nd.Sequence:
				g.generateSequence(w, node, n, count)

			default:
				g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
			}
		}
	default:
		g.errors = append(g.errors, fmt.Errorf("unknown plain child: %v", n.Types[n.Type]))
	}
}
