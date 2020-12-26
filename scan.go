package pegen

import (
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
)

type scan struct {
	name string
}

func (g *Generator) parseScan(n *pegn.Node) {
	var scan scan
	for _, n := range n.Children() {
		switch n.Type {
		case nd.CheckId:
			scan.name = n.Value
		}
	}
	g.scans = append(g.scans, scan)
}
