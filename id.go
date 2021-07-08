package pegn

import (
	"fmt"

	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn/nd"
)

// ReservedIdentifierError indicates that the grammar overwrites one of the
// predefined/reserved (default) PEGN classes or tokens.
type ReservedIdentifierError struct {
	identifier string
}

func (r *ReservedIdentifierError) Error() string {
	return fmt.Sprintf("redefining reserved identifier: %s", r.identifier)
}

// getID extracts the token/class identifier from the given node. The node must
// be of type nd.ClassId or nd.TokenId. If the generator is configured to return an error on
// reserved classes then this will get appended to the generator errors list.
func (g *generator) getID(n *ast.Node) (string, error) {
	id := n.Value
	switch n.Type {
	case nd.CheckId:
		return g.nodeName(id), nil
	case nd.ClassId:
		return g.className(id), nil
	case nd.ResClassId:
		if !g.config.IgnoreReserved {
			return id, &ReservedIdentifierError{
				identifier: id,
			}
		}
		return g.className(id), nil
	case nd.TokenId:
		return g.tokenName(n.Value), nil
	case nd.ResTokenId:
		if !g.config.IgnoreReserved {
			return id, &ReservedIdentifierError{
				identifier: id,
			}
		}
		return g.tokenName(id), nil
	}
	return id, nil
}
