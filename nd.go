package pegen

import "fmt"

func (g *Generator) typeNameGenerated(s string) string {
	s = g.typeName(s)
	if pkg := g.config.TypeSubPackage; pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, s)
	}
	return s
}

func (g *Generator) typeName(s string) string {
	return fmt.Sprintf("%s%s", s, g.config.TypeSuffix)
}

func (g *Generator) generateTypes() error {
	w := g.writers["nd"]
	var index int

	w.c("Node Types")
	w.wln("const (")
	{
		w := w.indent()
		w.wln("Unknown = iota")
		w.ln()

		if len(g.nodes) != 0 {
			w.cf("%s (%s)", g.meta.language, g.meta.url)
			indent := g.longestTypeName() + 1
			for _, node := range g.nodes {
				if !node.scan {
					index++
					w.w(fillRight(g.typeName(g.nodeName(node.name)), indent))
					w.noIndent().cf("%03d", index)
				}
			}
		}

		for _, dep := range g.dependencies {
			if len(dep.nodes) == 0 {
				continue
			}
			w.cf("%s (%s)", dep.meta.language, dep.meta.url)
			indent := g.longestTypeName() + 1
			for _, node := range dep.nodes {
				if !node.scan {
					index++
					w.w(fillRight(g.typeName(g.nodeName(node.name)), indent))
					w.noIndent().cf("%03d", index)
				}
			}
		}
	}
	w.wln(")")
	w.ln()
	w.wln("var NodeTypes = []string{")
	{
		w := w.indent()
		w.wln("\"UNKNOWN\",")
		w.ln()

		if len(g.nodes) != 0 {
			w.cf("%s (%s)", g.meta.language, g.meta.url)
			for _, node := range g.nodes {
				if !node.scan {
					w.wlnf("%q,", g.nodeName(node.name))
				}
			}
		}

		for _, dep := range g.dependencies {
			if len(dep.nodes) == 0 {
				continue
			}
			w.cf("%s (%s)", dep.meta.language, dep.meta.url)
			for _, node := range dep.nodes {
				if !node.scan {
					w.wlnf("%q,", dep.nodeName(node.name))
				}
			}
		}
	}
	w.wln("}")
	return nil
}
