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

func (g *Generator) generateTypes() error {
	w := g.writers["nd"]

	w.c("Node Types")
	w.wln("const (")
	{
		w := w.indent()
		w.wln("Unknown = iota")
		w.ln()
		for _, node := range g.nodes {
			if !node.scan {
				w.wln(g.typeName(g.nodeName(node.name)))
			}
		}
	}
	w.wln(")")
	w.ln()
	w.wln("var NodeTypes = []string{")
	{
		w := w.indent()
		w.wln("\"UNKNOWN\",")
		for _, node := range g.nodes {
			if !node.scan {
				w.wlnf("%q,", g.nodeName(node.name))
			}
		}
	}
	w.wln("}")
	return nil
}
