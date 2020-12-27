package pegen

func (g *Generator) generateTypes() {
	w := g.writers["nd"]

	w.c("Node Types")
	w.wln("const (")
	{
		w := w.indent()
		w.wln("Unknown = iota")
		w.ln()
		for _, node := range g.nodes {
			if node.scan {
				continue
			}
			w.wln(g.typeName(node.name))
		}
	}
	w.wln(")")
	w.ln()
	w.wln("var NodeTypes = [...]string{")
	{
		w := w.indent()
		w.wln("\"UNKNOWN\",")
		for _, node := range g.nodes {
			if node.scan {
				continue
			}
			w.wlnf("%q,", node.name)
		}
	}
	w.wln("}")
}
