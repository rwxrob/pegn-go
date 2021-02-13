package pegn

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn/nd"
	"strings"
)

// class is a simplified representation of a ClassDef.
//
// PEGN:
//	ClassDef <-- ClassId SP+ '<-' SP+ ClassExpr
type class struct {
	// name represents the ClassId.
	name string
	// expression contains the children of the ClassExpr contained within the
	// definition.
	expression []*ast.Node
}

// classNameGenerated returns a formatted className AND adds the prefix of the
// sub package if present.
func (g *generator) classNameGenerated(name string) string {
	if pkg := g.config.ClassSubPackage; pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, g.className(name))
	}
	return g.className(name)
}

// className returns a formatted class name. This includes the following:
//	- Applying aliases.
//	- Removing potential underscores.
//	- Camel Case + Capitalizing.
func (g *generator) className(name string) string {
	if alias, ok := g.config.ClassAliases[name]; ok {
		name = alias
	}
	parts := strings.Split(name, "_")
	for i, s := range parts {
		parts[i] = strings.Title(s)
	}
	return strings.Join(parts, "")
}

// parseClass parses the given node as a class and adds it to the list of
// classes within the generator.
//
// PEGN:
//	ClassDef <-- ClassId SP+ '<-' SP+ ClassExpr
func (g *generator) parseClass(n *ast.Node) error {
	var class class
	for _, n := range n.Children() {
		switch n.Type {
		case nd.ClassId, nd.ResClassId:
			// ClassId    <-- lower (lower / UNDER lower)+
			// ResClassId <-- 'alphanum' / 'alpha' / 'any' / etc...
			name, err := g.getID(n)
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

// generateClasses writes all the classes to the given writer.
func (g *generator) generateClasses(w *writer) error {
	for idx, class := range g.classes {
		size := len(class.expression)
		if size == 1 {
			// I chose to not generate alias class definitions.
			// e.g. `cntrl <- control`
			if c := class.expression[0]; c.Type == nd.ClassId {
				continue
			}
		}

		w.wlnf("func %s(p *parser.Parser) (*parser.Cursor, bool) {", g.className(class.name))
		{
			w := w.indent()
			w.w("return p.Check(")

			if 1 < size {
				// Wrap the children in an or if the class has more than one
				// child. e.g. `uphex <- [0-9] / [A-F]`
				w.noIndent().w("op.Or{")
				w.ln()
				w = w.indent()
			} else {
				// Inline if the size is 1.
				// e.g. `upper <- [A-Z]`
				w = w.noIndent()
			}

			for _, n := range class.expression {
				switch n.Type {
				case nd.Comment, nd.EndLine:
					// Ignore these.
					continue
				case nd.Unicode, nd.Hexadecimal:
					writeUnicode(n.Value, w)
				case nd.Binary:
					writeBinary(n.Value, w)
				case nd.Octal:
					writeOctal(n.Value, w)
				case nd.ClassId, nd.ResClassId:
					id, err := g.getID(n)
					if err != nil {
						return err
					}
					w.w(id)
				case nd.TokenId, nd.ResTokenId:
					id, err := g.getID(n)
					if err != nil {
						return err
					}
					w.w(g.tokenNameGenerated(id))
				case nd.AlphaRange:
					writeAlphaRange(n, w)
				case nd.IntRange:
					writeIntRange(n, w)
				case nd.UniRange, nd.HexRange:
					writeUnicodeRange(n, w)
				case nd.BinRange:
					writeBinaryRange(n, w)
				case nd.OctRange:
					writeOctalRange(n, w)
				case nd.String:
					writeString(n.Value, w)
				default:
					return fmt.Errorf("unknown class child: %v", nd.NodeTypes[n.Type])
				}

				// It should only be comma separated if wrapped in an op.Or,
				// thus containing more than one child (see above).
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
