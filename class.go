package pegen

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn"
	"strconv"
	"strings"
)

type class struct {
	name       string
	expression []*ast.Node
}

func (g *Generator) className(s string) string {
	// Check whether the class has an alias.
	if c, ok := g.config.ClassAliases[s]; ok {
		s = c
	}
	return strings.Title(s)
}

// ClassDef <-- ClassId SP+ '<-' SP+ ClassExpr
func (g *Generator) parseClass(n *ast.Node) error {
	var class class
	for _, n := range n.Children() {
		switch n.Type {
		case pegn.ClassIdType, pegn.ResClassIdType:
			// ClassId    <-- lower (lower / UNDER lower)+
			// ResClassId <-- 'alphanum' / 'alpha' / 'any' / etc...
			name, err := g.GetID(n)
			if err != nil {
				return err
			}
			class.name = name
		case pegn.ClassExprType:
			// ClassExpr <-- Simple (Spacing '/' SP+ Simple)*
			class.expression = n.Children()
		default:
			return fmt.Errorf("unknown class child: %v", pegn.NodeTypes[n.Type])
		}
	}
	g.classes = append(g.classes, class)
	return nil
}

func (g *Generator) generateClasses() error {
	w := g.writers["is"]

	for idx, class := range g.classes {
		size := len(class.expression)
		if size == 1 {
			// Duplicate (alias) class definition.
			if c := class.expression[0]; c.Type == pegn.ClassIdType {
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
				case pegn.CommentType, pegn.EndLineType:
					// Ignore these.
					continue
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
					pegn.TokenIdType, pegn.ResTokenIdType:
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
					w.wf("%q", n.ValueString())
				default:
					return fmt.Errorf("unknown class child: %v", pegn.NodeTypes[n.Type])
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
	return nil
}
