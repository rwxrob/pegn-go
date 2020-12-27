package pegen

import (
	"errors"
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
	"strings"
)

type tokens []token

func (ts tokens) get(id string) token {
	for _, tk := range ts {
		if tk.name == id {
			return tk
		}
	}
	return token{}
}

type token struct {
	comment  string // comment after the token.
	name     string // name of the token.
	value    string // value of the token if it is a string.
	hexValue string // hex value of token if it is a rune.
}

func (t *token) isString() bool {
	return t.value != "" && t.hexValue == ""
}

func (g *Generator) tokenName(s string) string {
	if prefix := g.config.TokenPrefix; prefix != "" {
		prefix := strings.ToUpper(prefix)
		return fmt.Sprintf("%s_%s", prefix, s)
	}
	return s
}

// TokenDef <-- TokenId SP+ '<-' SP+
//              TokenVal (Spacing TokenVal)*
//              ComEndLine
func (g *Generator) parseToken(n *pegn.Node) {
	var token token
	for _, n := range n.Children() {
		switch n.Type {
		case nd.Comment:
			// ComEndLine
			token.comment = n.Value
		case nd.EndLine:
			// Ignore this.

		case nd.TokenId:
			// TokenId <-- ResTokenId / upper (upper / UNDER upper)+
			// ResTokenId <-- 'TAB' / 'CRLF' / 'CR' / etc...
			token.name = g.tokenName(g.GetID(n))

		// TokenVal (Spacing TokenVal)*
		// TokenVal <- Unicode / Binary / Hexadec / Octal / SQ String SQ
		case nd.Unicode, nd.Hexadec:
			hex, err := ConvertToHex(n.Value[1:], 16)
			if err != nil {
				g.errors = append(g.errors, err)
				return
			}
			token.hexValue = hex
		case nd.Binary:
			hex, err := ConvertToHex(n.Value[1:], 2)
			if err != nil {
				g.errors = append(g.errors, err)
				return
			}
			token.hexValue = hex
		case nd.Octal:
			hex, err := ConvertToHex(n.Value[1:], 8)
			if err != nil {
				g.errors = append(g.errors, err)
				return
			}
			token.hexValue = hex
		case nd.String:
			token.value = n.Value
		default:
			g.errors = append(g.errors, errors.New("unknown token child"))
		}
	}
	g.tokens = append(g.tokens, token)
}

func (g *Generator) generateTokens() {
	w := g.writers["tk"]

	w.c("Token Definitions")
	w.wln("const (")
	{
		w := w.indent()
		longestName := g.longestTokenName()
		for _, token := range g.tokens {
			var value string
			if token.isString() {
				// Strings: "value"
				value = fmt.Sprintf("%q", token.value)
			} else {
				value, _ = ConvertToRuneString(token.hexValue, 16)
			}
			w.wf("%s = %s", fillRight(token.name, longestName), value)
			{
				w := w.noIndent()
				if token.comment != "" {
					w.w(" ")
					w.c(token.comment)
				} else {
					w.ln()
				}
			}
		}
	}
	w.wln(")")
}
