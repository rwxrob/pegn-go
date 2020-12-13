package pegen

import (
	"errors"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
)

type node struct {
	name string
}

func (g *Generator) parseNode(n *pegn.Node) {
	var node node
	for _, n := range n.Children() {
		switch n.Type {
		case nd.CheckId:
			// Reserved class identifier.
			if len(n.Children()) != 0 {
				if !g.config.IgnoreReserved {
					g.errors = append(g.errors, errors.New("redefining reserved token identifier"))
				}
				node.name = n.Children()[0].Value
				break
			}
			node.name = n.Value
		}
	}
	g.nodes = append(g.nodes, node)
}
