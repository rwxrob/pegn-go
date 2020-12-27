package pegen

import "fmt"

func (g *Generator) typeName(s string) string {
	if prefix := g.config.TypePrefix; prefix != "" {
		s = fmt.Sprintf("%s%s", prefix, s)
	}
	if suffix := g.config.TypeSuffix; suffix != "" {
		s = fmt.Sprintf("%s%s", s, suffix)
	}
	return s
}

func (g *Generator) generateTypes() {
	w := g.writers["nd"]

	w.c("Node Types")
	w.wln("const (")
	{
		w := w.indent()
		w.wln("Unknown = iota")
		w.ln()
		for _, node := range g.nodes {
			w.wln(g.typeName(node.name))
		}
	}
	w.wln(")")
	w.ln()
	w.wln("var NodeTypes = []string{")
	{
		w := w.indent()
		w.wln("\"UNKNOWN\",")
		for _, node := range g.nodes {
			w.wlnf("%q,", node.name)
		}
	}
	w.wln("}")
}
