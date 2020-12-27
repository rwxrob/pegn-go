package pegen

import (
	"errors"
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
)

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
		case nd.Expression:
			// Expression <-- Sequence (Spacing '/' SP+ Sequence)*
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
	w := g.writers["ast"]

	for idx, node := range g.nodes {
		w.wlnf("func %s(p *pegn.Parser) (*pegn.Node, error) {", node.name)
		for _, n := range node.expression {
			// Sequence <-- Rule (Spacing Rule)*
			switch n.Type {
			case nd.Comment, nd.EndLine:
				// Ignore these.
			case nd.Sequence:
			default:
				g.errors = append(g.errors, fmt.Errorf("unknown class child: %v", n.Types[n.Type]))
			}
		}
		w.indent().wln("return nil, nil")
		w.wln("}")
		if idx <= len(g.nodes)-2 {
			// All except the last one.
			w.ln()
		}
	}
}
