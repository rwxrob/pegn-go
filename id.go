package pegen

import (
	"fmt"
	"gitlab.com/pegn/pegn-go"
)

type ReservedIdentifierError struct {
	identifier string
}

func (r *ReservedIdentifierError) Error() string {
	return fmt.Sprintf("redefining reserved identifier: %s", r.identifier)
}

// GetID extracts the token/class identifier from the given node. The node must
// be of type nd.ClassId or nd.TokenId. If the generator is configured to return an error on
// reserved classes then this will get appended to the generator errors list.
func (g *Generator) GetID(n *pegn.Node) string {
	// 1. Reserved class identifier.
	if len(n.Children()) != 0 {
		id := n.Children()[0].Value
		if !g.config.IgnoreReserved {
			g.errors = append(g.errors, &ReservedIdentifierError{
				identifier: id,
			})
		}
		return id
	}
	// Normal token identifier.
	return n.Value
}
