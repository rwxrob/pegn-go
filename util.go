package pegen

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	"strings"
)

func (g *Generator) longestTokenName() int {
	var length int
	for _, token := range g.tokens {
		if l := len(token.name); length < l {
			length = l
		}
	}
	return length
}

func (g *Generator) longestTokenValueWithComment(idx int) int {
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

func (g *Generator) longestTypeName() int {
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

func fillRight(v string, size int) string {
	return fmt.Sprintf("%s%s", v, strings.Repeat(" ", size-len(v)))
}

func singleNestedValue(n *ast.Node) bool {
	if !n.IsParent() {
		return true
	}
	if len(n.Children()) != 1 {
		return false
	}
	return singleNestedValue(n.Children()[0])
}
