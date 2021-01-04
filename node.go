package pegen

import (
	"errors"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
)

type nodes []node

func (ns nodes) get(id string) node {
	for _, node := range ns {
		if node.name == id {
			return node
		}
	}
	return node{}
}

type node struct {
	name       string
	scan       bool
	expression []*pegn.Node
}

func (g *Generator) parseNode(n *pegn.Node) {
	var node node
	for _, n := range n.Children() {
		switch n.Type {
		case nd.CheckId:
			node.name = n.Value

		// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
		case nd.Expression:
			node.expression = n.Children()
		default:
			g.errors = append(g.errors, errors.New("unknown node child"))
		}
	}
	g.nodes = append(g.nodes, node)
}

func (g *Generator) parseScan(n *pegn.Node) {
	scan := node{
		scan: true,
	}
	for _, n := range n.Children() {
		switch n.Type {
		case nd.CheckId:
			scan.name = n.Value
		case nd.Expression:
			// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
			scan.expression = n.Children()
		default:
			g.errors = append(g.errors, errors.New("unknown scan child"))
		}
	}
	g.nodes = append(g.nodes, scan)
}

func (g *Generator) generateNodes() {
	//  Write to the ast buffer.
	w := g.writers["ast"]
	w.c("TODO: nodes")
	for _, node := range g.nodes {
		_ = node
	}
}
