package pegen

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegen/pegn"
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
func (g *Generator) GetID(n *ast.Node) (string, error) {
	id := n.ValueString()
	switch n.Type {
	case pegn.CheckIdType:
		return g.nodeName(id), nil
	case pegn.ClassIdType:
		return g.className(id), nil
	case pegn.ResClassIdType:
		if !g.config.IgnoreReserved {
			return id, &ReservedIdentifierError{
				identifier: id,
			}
		}
		return g.className(id), nil
	case pegn.TokenIdType:
		return g.tokenName(n.ValueString()), nil
	case pegn.ResTokenIdType:
		if !g.config.IgnoreReserved {
			return id, &ReservedIdentifierError{
				identifier: id,
			}
		}
		return g.tokenName(id), nil
	}
	return id, nil
}

func (g *Generator) getID(n *ast.Node) (string, error) {
	// 1. Reserved class identifier.
	if len(n.Children()) != 0 {
		id := n.Children()[0].ValueString()
		if !g.config.IgnoreReserved {
			return id, &ReservedIdentifierError{
				identifier: id,
			}
		}
		return id, nil
	}
	// Normal token identifier.
	return n.ValueString(), nil
}
