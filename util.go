package pegn

import (
	"fmt"
	"strings"

	"github.com/di-wu/parser/ast"
)

func (g generator) moduleName() string {
	if g.config.ModuleName != "" {
		return g.config.ModuleName
	}
	return g.meta.language
}

// longestTokenName returns the length of the longest token name.
func (g *generator) longestTokenName() int {
	var length int
	for _, token := range g.tokens {
		if l := len(token.name); length < l {
			length = l
		}
	}
	return length
}

// longestTokenValueWithComment returns the length of the longest token name
// including the comment.
func (g *generator) longestTokenValueWithComment(idx int) int {
	var length int
	for _, token := range g.tokens[idx:] {
		if token.comment == "" {
			if length != 0 {
				break
			}
			continue
		}
		if l := len(token.value); length < l {
			length = l
		}
	}
	return length
}

// longestTypeName returns the length of the longest type name.
func (g *generator) longestTypeName() int {
	var length int
	for _, node := range g.nodes {
		if !node.scan {
			if l := len(g.typeName(g.nodeName(node.name))); length < l {
				length = l
			}
		}
	}
	return length
}

// fillRight fills the given string on the right with spaces until it reaches
// the requested size(/length).
func fillRight(v string, size int) string {
	return fmt.Sprintf("%s%s", v, strings.Repeat(" ", size-len(v)))
}

// singleNestedValue returns whether the given node is a single nested value.
func singleNestedValue(n *ast.Node) bool {
	if !n.IsParent() {
		return true
	}
	if len(n.Children()) != 1 {
		return false
	}
	return singleNestedValue(n.Children()[0])
}
