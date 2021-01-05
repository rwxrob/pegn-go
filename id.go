package pegen

import (
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
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
func (g *Generator) GetID(n *pegn.Node) (string, error) {
	id, err := g.getID(n)
	if err != nil {
		return id, err
	}
	switch n.Type {
	case nd.CheckId:
		return g.nodeName(id), nil
	case nd.ClassId:
		return g.className(id), nil
	case nd.TokenId:
		return g.tokenName(id), nil
	}
	fmt.Println(n.Type)
	return id, nil
}

func (g *Generator) getID(n *pegn.Node) (string, error) {
	// 1. Reserved class identifier.
	if len(n.Children()) != 0 {
		id := n.Children()[0].Value
		if !g.config.IgnoreReserved {
			return id, &ReservedIdentifierError{
				identifier: id,
			}
		}
		return id, nil
	}
	// Normal token identifier.
	return n.Value, nil
}
