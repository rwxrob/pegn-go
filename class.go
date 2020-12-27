package pegen

import (
	"errors"
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
	"strconv"
	"strings"
)

type class struct {
	name       string
	expression []*pegn.Node
}

func (g *Generator) className(s string) string {
	return strings.Title(s)
}

// ClassDef <-- ClassId SP+ '<-' SP+ ClassExpr
func (g *Generator) parseClass(n *pegn.Node) {
	var class class
	for _, n := range n.Children() {
		switch n.Type {
		case nd.ClassId:
			// ClassId <-- ResClassId / lower (lower / UNDER lower)+
			// ResClassId <-- 'alphanum' / 'alpha' / 'any' / etc...
			class.name = g.GetID(n)
		case nd.ClassExpr:
			// ClassExpr <-- Simple (Spacing '/' SP+ Simple)*
			class.expression = n.Children()
		default:
			g.errors = append(g.errors, errors.New("unknown class child"))
		}
	}
	g.classes = append(g.classes, class)
}

func (g *Generator) generateClasses() {
	w := g.writers["is"]

	for idx, class := range g.classes {
		w.wlnf("var %s = %s{}", g.className(class.name), class.name)
		w.ln()

		w.wlnf("type %s struct{}", class.name)
		w.ln()

		w.wlnf("func (%s) Ident() string { ", class.name)
		w.indent().wlnf("return %q", class.name)
		w.wln("}")
		w.ln()
		w.wlnf("func (%s) PEGN() string  { ", class.name)
		w.indent().wlnf("return %q", "PEGN: unavailable") // TODO
		w.wln("}")
		w.ln()

		w.wlnf("func (%s) Desc() string  { ", class.name)
		w.indent().wlnf("return %q", "DESC: unavailable") // TODO
		w.wln("}")
		w.ln()

		w.wlnf("func (%s) Check(r rune) bool  { ", class.name)
		for _, n := range class.expression {
			// Simple <- Unicode / Binary / Hexadec / Octal /
			//           ClassId / TokenId / Range / SQ String SQ
			// Range <- AlphaRange / IntRange / UniRange /
			//          BinRange / HexRange / OctRange
			switch n.Type {
			case nd.Comment, nd.EndLine:
				// Ignore these.
			case nd.Unicode, nd.Hexadec, nd.Binary, nd.Octal:
				g.errors = append(g.errors, fmt.Errorf("%s values are not supported", n.Types[n.Type]))
			case nd.ClassId:
				id := g.className(g.GetID(n))

				w := w.indent()
				w.wlnf("if %s.Check(r) {", id)
				w.indent().wln("return true")
				w.wln("}")
			case nd.TokenId:
				id := g.tokenName(g.GetID(n))
				tk := g.tokens.get(id)
				if tk.isString() {
					g.errors = append(g.errors, errors.New("token value is a string"))
					break
				}

				w := w.indent()
				w.wlnf("if r == %s {", id)
				w.indent().wln("return true")
				w.wln("}")
			case nd.AlphaRange:
				// AlphaRange <-- '[' Letter '-' Letter ']'
				min := n.Children()[0].Value
				max := n.Children()[1].Value
				w := w.indent()
				w.wlnf("if '%s' <= r && r <= '%s' {", min, max)
				w.indent().wln("return true")
				w.wln("}")
			case nd.IntRange:
				// IntRange <-- '[' Integer '-' Integer ']'
				min, _ := strconv.Atoi(n.Children()[0].Value)
				max, _ := strconv.Atoi(n.Children()[1].Value)
				if min < 0 {
					g.errors = append(g.errors, fmt.Errorf("int range is negative: [%v-%v]", min, max))
				}
				if max <= min {
					g.errors = append(g.errors, fmt.Errorf("int range is inverted: [%v-%v]", min, max))
				}
				if 10 <= max {
					g.errors = append(g.errors, fmt.Errorf("int range too large: [%v-%v]", min, max))
				}
				w := w.indent()
				w.wlnf("if '%v' <= r && r <= '%v' {", min, max)
				w.indent().wln("return true")
				w.wln("}")
			case nd.UniRange, nd.HexRange:
				// UniRange <-- '[' Unicode '-' Unicode ']'
				// HexRange <-- '[' Hexadec '-' Hexadec ']'
				min, _ := ConvertToRuneString(n.Children()[0].Value[1:], 16)
				max, _ := ConvertToRuneString(n.Children()[1].Value[1:], 16)
				w := w.indent()
				w.wlnf("if %s <= r && r <= %s {", min, max)
				w.indent().wln("return true")
				w.wln("}")
			case nd.BinRange:
				// BinRange <-- '[' Binary '-' Binary ']'
				min, _ := ConvertToRuneString(n.Children()[0].Value[1:], 2)
				max, _ := ConvertToRuneString(n.Children()[1].Value[1:], 2)
				w := w.indent()
				w.wlnf("if %s <= r && r <= %s {", min, max)
				w.indent().wln("return true")
				w.wln("}")
			case nd.OctRange:
				// OctRange <-- '[' Octal '-' Octal ']'
				min, _ := ConvertToRuneString(n.Children()[0].Value[1:], 8)
				max, _ := ConvertToRuneString(n.Children()[1].Value[1:], 8)
				w := w.indent()
				w.wlnf("if %s <= r && r <= %s {", min, max)
				w.indent().wln("return true")
				w.wln("}")
			case nd.String:
				g.errors = append(g.errors, fmt.Errorf("string values are not supported"))
			default:
				g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
			}
		}
		w.indent().wln("return false")
		w.wln("}")
		if idx <= len(g.classes)-2 {
			// All except the last one.
			w.ln()
		}
	}
}
