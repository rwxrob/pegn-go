package pegen

import (
	"errors"
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
	"strings"
)

type class struct {
	name       string
	expression []*pegn.Node
}

// ClassDef <-- ClassId SP+ '<-' SP+ ClassExpr
func (g *Generator) parseClass(n *pegn.Node) {
	var class class
	for _, n := range n.Children() {
		switch n.Type {
		case nd.ClassId:
			// ClassId <-- ResClassId / lower (lower / UNDER lower)+
			// ResClassId <-- 'alphanum' / 'alpha' / 'any' / etc...

			// 1. Reserved class identifier.
			if len(n.Children()) != 0 {
				if !g.config.IgnoreReserved {
					g.errors = append(g.errors, errors.New("redefining reserved class identifier"))
				}
				class.name = n.Children()[0].Value
				break
			}
			// 2. Normal token identifier.
			class.name = n.Value
		case nd.ClassExpr:
			// ClassExpr <-- Simple (Spacing '/' SP+ Simple)*
			// Simple <- Unicode / Binary / Hexadec / Octal /
			//           ClassId / TokenId / Range / SQ String SQ
			// Range <- AlphaRange / IntRange / UniRange /
			//          BinRange / HexRange / OctRange
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
		w.wlnf("var %s = %s{}", strings.Title(class.name), class.name)
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
			switch n.Type {
			case nd.Comment, nd.EndLine:
				// Ignore these.
			case nd.Unicode, nd.Hexadec:
			case nd.Binary:
				g.errors = append(g.errors, errors.New("binary values are not supported"))
			case nd.Octal:
				g.errors = append(g.errors, errors.New("octal values are not supported"))
			case nd.ClassId:
			case nd.TokenId:
			case nd.AlphaRange:
			case nd.IntRange:
			case nd.UniRange:
			case nd.BinRange:
			case nd.HexRange:
			case nd.OctRange:
			case nd.String:
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
