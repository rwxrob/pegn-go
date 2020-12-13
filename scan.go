package pegen

import (
	"errors"
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
			// Reserved class identifier.
			if len(n.Children()) != 0 {
				if !g.config.IgnoreReserved {
					g.errors = append(g.errors, errors.New("redefining reserved token identifier"))
				}
				scan.name = n.Children()[0].Value
				break
			}
			scan.name = n.Value
		}
	}
	g.scans = append(g.scans, scan)
}
