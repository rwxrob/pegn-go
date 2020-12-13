package pegen

func (g *Generator) generateTypes() {
	w := g.writer

	w.c("Node Types")
	w.wln("const (")
	{
		w := w.indent()
		w.wln("Unknown = iota")
		w.ln()
		for _, node := range g.nodes {
			w.wln(node.name)
		}
	}
	w.wln(")")
	w.ln()
	w.wln("var NodeTypes = [...]string{")
	{
		w := w.indent()
		w.wln("\"UNKNOWN\",")
		for _, node := range g.nodes {
			w.wlnf("%q,", node.name)
		}
	}
	w.wln("}")
}
