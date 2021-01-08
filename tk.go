package pegen

import (
	"fmt"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegen/pegn"
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
	comment string // comment after the token.
	name    string // name of the token.
	values  []tokenValue
}

type tokenValue struct {
	value    string // value of the token if it is a string.
	hexValue string // hex value of token if it is a rune.
}

func (t *tokenValue) isString() bool {
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
func (g *Generator) parseToken(n *ast.Node) error {
	var token token
	for _, n := range n.Children() {
		switch n.Type {
		case pegn.CommentType:
			// ComEndLine
			token.comment = n.ValueString()
		case pegn.EndLineType:
			// Ignore this.

		case pegn.TokenIdType, pegn.ResTokenIdType:
			// TokenId    <-- upper (upper / UNDER upper)+
			// ResTokenId <-- 'TAB' / 'CRLF' / 'CR' / etc...
			id, err := g.GetID(n)
			if err != nil {
				return err
			}
			token.name = id

		// TokenVal (Spacing TokenVal)*
		// TokenVal <- Unicode / Binary / Hexadec / Octal / SQ String SQ
		case pegn.UnicodeType, pegn.HexadecimalType:
			hex, err := ConvertToHex(n.ValueString()[1:], 16)
			if err != nil {
				return err
			}
			token.values = append(token.values, tokenValue{
				hexValue: hex,
			})
		case pegn.BinaryType:
			hex, err := ConvertToHex(n.ValueString()[1:], 2)
			if err != nil {
				return err
			}
			token.values = append(token.values, tokenValue{
				hexValue: hex,
			})
		case pegn.OctalType:
			hex, err := ConvertToHex(n.ValueString()[1:], 8)
			if err != nil {
				return err
			}
			token.values = append(token.values, tokenValue{
				hexValue: hex,
			})
		case pegn.StringType:
			token.values = append(token.values, tokenValue{
				value: n.ValueString(),
			})
		default:
			return fmt.Errorf("unknown token child: %v", pegn.NodeTypes[n.Type])
		}
	}
	g.tokens = append(g.tokens, token)
	return nil
}

func (g *Generator) generateTokens() error {
	w := g.writers["tk"]

	w.c("Token Definitions")
	w.wln("const (")
	{
		w := w.indent()
		longestName := g.longestTokenName()
		for _, token := range g.tokens {
			var value string
			if len(token.values) == 1 {
				tk := token.values[0]
				if tk.isString() {
					// Strings: "value"
					value = fmt.Sprintf("%q", tk.value)
				} else {
					value, _ = ConvertToRuneString(tk.hexValue, 16)
				}
			} else {
				for _, tk := range token.values {
					if tk.isString() {
						value += tk.value
					} else {
						v, _ := ConvertToInt(tk.hexValue, 16)
						value += string(rune(v))
					}
				}
				value = fmt.Sprintf("%q", value)
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
	return nil
}
