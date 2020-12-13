package pegen

import (
	"errors"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
)

type class struct {
	name string
}

func (g *Generator) parseClass(n *pegn.Node)  {
	var class class
	for _, n := range n.Children() {
		switch n.Type {
		case nd.ClassId:
			// Reserved class identifier.
			if len(n.Children()) != 0 {
				if !g.config.IgnoreReserved {
					g.errors = append(g.errors, errors.New("redefining reserved token identifier"))
				}
				class.name = n.Children()[0].Value
				break
			}
			class.name = n.Value
		}
	}
	g.classes = append(g.classes, class)
}
