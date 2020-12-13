package pegen

func (g *Generator) longestTokenName() int {
	var length int
	for _, token := range g.tokens {
		if l := len(token.name); length < l {
			length = l
		}
	}
	return length
}
