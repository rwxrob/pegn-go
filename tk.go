package pegen

import (
	"encoding/hex"
	"errors"
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/nd"
	"strings"
)

type token struct {
	comment string
	name    string
	value   string
	values  []byte
}

func (g *Generator) parseToken(n *pegn.Node) {
	var token token
	for _, n := range n.Children() {
		switch n.Type {
		case nd.Comment:
			token.comment = n.Value
		case nd.EndLine:
			// Ignore this.
		case nd.TokenId:
			// Reserved token identifier.
			if len(n.Children()) != 0 {
				if !g.config.IgnoreReserved {
					g.errors = append(g.errors, errors.New("redefining reserved token identifier"))
				}
				token.name = n.Children()[0].Value
				break
			}
			token.name = n.Value

		case nd.Unicode, nd.Hexadec:
			hexValue := n.Value[1:]
			if len(hexValue)%2 != 0 {
				hexValue = fmt.Sprintf("0%s", hexValue)
			}
			v, err := hex.DecodeString(hexValue)
			if err != nil {
				panic(err)
			}

			// Go does not support unicode escapes bigger than two bytes.
			if 2 < len(v) {
				return
			}
			token.values = append(token.values, v...)
		case nd.Binary:
			g.errors = append(g.errors, errors.New("binary token value not supported"))
		case nd.Octal:
			g.errors = append(g.errors, errors.New("octal token value not supported"))
		case nd.String:
			token.value = n.Value
		default:
			g.errors = append(g.errors, errors.New("unknown token type"))
		}
	}
	g.tokens = append(g.tokens, token)
}

func (g *Generator) generateTokens() {
	w := g.writer

	w.c("Token Definitions")
	w.wln("const (")
	{
		w := w.indent()
		longestName := g.longestTokenName()
		for _, token := range g.tokens {
			var value string
			if token.value != "" {
				// Strings: "value"
				value = fmt.Sprintf("%q", token.value)
			} else {
				// Hex: (00)07
				value = strings.ToUpper(hex.EncodeToString(token.values))
				if len(value) == 2 {
					// Add zeros if not present.
					value = fmt.Sprintf("00%s", value)
				}
				// Runes: '\u0000'
				value = fmt.Sprintf("'\\u%s'", value)
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
