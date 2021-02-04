package pegn

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn/nd"
	"strconv"
	"strings"
)

type class struct {
	name       string
	expression []*ast.Node
}

func (g *Generator) classNameGenerated(s string) string {
	s = g.className(s)
	if pkg := g.config.ClassSubPackage; pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, s)
	}
	return s
}

func (g *Generator) className(s string) string {
	// Check whether the class has an alias.
	if c, ok := g.config.ClassAliases[s]; ok {
		s = c
	}
	parts := strings.Split(s, "_")
	for i, s := range parts {
		parts[i] = strings.Title(s)
	}
	return strings.Join(parts, "")
}

// ClassDef <-- ClassId SP+ '<-' SP+ ClassExpr
func (g *Generator) parseClass(n *ast.Node) error {
	var class class
	for _, n := range n.Children() {
		switch n.Type {
		case nd.ClassId, nd.ResClassId:
			// ClassId    <-- lower (lower / UNDER lower)+
			// ResClassId <-- 'alphanum' / 'alpha' / 'any' / etc...
			name, err := g.GetID(n)
			if err != nil {
				return err
			}
			class.name = name
		case nd.ClassExpr:
			// ClassExpr <-- Simple (Spacing '/' SP+ Simple)*
			class.expression = n.Children()
		default:
			return fmt.Errorf("unknown class child: %v", nd.NodeTypes[n.Type])
		}
	}
	g.classes = append(g.classes, class)
	return nil
}

func (g *Generator) generateClasses(w *writer) error {
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
				case nd.Unicode, nd.Hexadecimal:
					v, _ := ConvertToRuneString(n.Value[1:], 16)
					w.w(v)
				case nd.Binary:
					v, _ := ConvertToRuneString(n.Value[1:], 2)
					w.w(v)
				case nd.Octal:
					v, _ := ConvertToRuneString(n.Value[1:], 8)
					w.w(v)
				case nd.ClassId, nd.ResClassId:
					id, err := g.GetID(n)
					if err != nil {
						return err
					}
					w.w(id)
				case nd.TokenId, nd.ResTokenId:
					id, err := g.GetID(n)
					if err != nil {
						return err
					}
					w.w(g.tokenNameGenerated(id))
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
						return fmt.Errorf("int range is negative: [%v-%v]", min, max)
					}
					if max <= min {
						return fmt.Errorf("int range is inverted: [%v-%v]", min, max)
					}
					if 10 <= max {
						return fmt.Errorf("int range too large: [%v-%v]", min, max)
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
					if v := n.Value; len(v) == 1 {
						w.wf("'%s'", v)
					} else {
						w.wf("%q", v)
					}
				default:
					return fmt.Errorf("unknown class child: %v", nd.NodeTypes[n.Type])
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

	for _, dep := range g.dependencies {
		if err := dep.generateClasses(w); err != nil {
			return err
		}
	}

	return nil
}
