package pegen

import (
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
			node.name = n.Value
		}
	}
	g.nodes = append(g.nodes, node)
}
