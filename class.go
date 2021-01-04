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
	// Check whether the class has an alias.
	if c, ok := g.config.ClassAliases[s]; ok {
		s = c
	}
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
		size := len(class.expression)
		if size == 1 {
			// Duplicate (alias) class definition.
			if c := class.expression[0]; c.Type == nd.ClassId {
				continue
			}
		}

		w.wlnf("func %s(p *parser.Parser) (*parser.Cursor, bool) {", g.className(class.name))
		{
			w := w.indent()
			w.w("return p.Check(")
			if 1 < size {
				w.noIndent().w("op.Or{")
				w.ln()
				w = w.indent()
			} else {
				w = w.noIndent()
			}
			for _, n := range class.expression {
				switch n.Type {
				case nd.Comment, nd.EndLine:
					// Ignore these.
					continue
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
						continue
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
						continue
					}
					if max <= min {
						g.errors = append(g.errors, fmt.Errorf("int range is inverted: [%v-%v]", min, max))
						continue
					}
					if 10 <= max {
						g.errors = append(g.errors, fmt.Errorf("int range too large: [%v-%v]", min, max))
						continue
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
				default:
					g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
					continue
				}
				if 1 < size {
					w.noIndent().wln(",")
				}
			}
		}
		if 1 < size {
			w.indent().wln("})")
		} else {
			w.wln(")")
		}
		w.wln("}")
		if idx <= len(g.classes)-2 {
			// All except the last one.
			w.ln()
		}
	}
}
