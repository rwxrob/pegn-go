package pegn

import "fmt"

// typePrefix adds the package prefix without formatting the string.
func (g *generator) typePrefix(s string) string {
	if pkg := g.config.TypeSubPackage; pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, s)
	}
	return s
}

// typeNameGenerated returns a formatted typeName AND adds the prefix of the
// sub package if present.
func (g *generator) typeNameGenerated(s string) string {
	s = g.typeName(s)
	if pkg := g.config.TypeSubPackage; pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, s)
	}
	return s
}

// typeName returns a formatted type name with the (optionally) predefined
// suffix.
func (g *generator) typeName(s string) string {
	return fmt.Sprintf("%s%s", s, g.config.TypeSuffix)
}

// generateTypes writes all the types to the 'nd' writer.
func (g *generator) generateTypes() error {
	w := g.writers["nd"]
	var index int

	w.c("Node Types")
	w.wln("const (")
	{
		w := w.indent()
		w.wln("Unknown = iota")
		w.ln()

		if len(g.nodes) != 0 {
			w.cf("%s (%s)\n", g.languageFull(), g.meta.url)
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
			w.cf("%s (%s)\n", dep.languageFull(), dep.meta.url)
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
			w.cf("%s (%s)\n", g.languageFull(), g.meta.url)
			for _, node := range g.nodes {
				if !node.scan {
					w.wlnf("%q,", g.nodeName(node.name))
				}
			}
		}

		for i, dep := range g.dependencies {
			if len(dep.nodes) == 0 {
				continue
			}
			switch i {
			case 0:
				if len(g.nodes) != 0 {
					w.ln()
				}
			default:
				if len(g.dependencies[i-1].nodes) != 0 {
					w.ln()
				}
			}
			w.cf("%s (%s)\n", dep.languageFull(), dep.meta.url)
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
