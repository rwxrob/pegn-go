package pegn

import (
	"fmt"
	"strings"

	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn/nd"
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

// token is a simplified representation of a TokenDef.
type token struct {
	comment   string // comment after the token.
	name      string // name of the token.
	rawValues []tokenValue
	value     string
}

type tokenValue struct {
	value    string // value of the token if it is a string.
	hexValue string // hex value of token if it is a rune.
}

func (t *tokenValue) isString() bool {
	return t.value != "" && t.hexValue == ""
}

// tokenNameGenerated returns a formatted tokenName AND adds the prefix of the
// sub package if present.
func (g *generator) tokenNameGenerated(s string) string {
	s = g.tokenName(s)
	if pkg := g.config.TokenSubPackage; pkg != "" {
		return fmt.Sprintf("%s.%s", pkg, s)
	}
	return s
}

// tokenName returns a formatted token name.
func (g *generator) tokenName(s string) string {
	if prefix := g.config.TokenPrefix; prefix != "" {
		prefix := strings.ToUpper(prefix)
		return fmt.Sprintf("%s_%s", prefix, s)
	}
	return s
}

// parseToken parses the given node as a token and adds it to the list of
// tokens within the generator.
//
// PEGN:
//	TokenDef <-- TokenId SP+ '<-' SP+ TokenVal (Spacing TokenVal)* ComEndLine
func (g *generator) parseToken(n *ast.Node) error {
	var token token
	var values []tokenValue
	for _, n := range n.Children() {
		switch n.Type {
		case nd.Comment:
			// ComEndLine
			token.comment = n.Value
		case nd.EndLine:
			// Ignore this.

		case nd.TokenId, nd.ResTokenId:
			// TokenId    <-- upper (upper / UNDER upper)+
			// ResTokenId <-- 'TAB' / 'CRLF' / 'CR' / etc...
			id, err := g.getID(n)
			if err != nil {
				return err
			}
			token.name = id

		// TokenVal (Spacing TokenVal)*
		// TokenVal <- Unicode / Binary / Hexadec / Octal / SQ String SQ
		case nd.Unicode, nd.Hexadecimal:
			hex, err := convertToHex(n.Value[1:], 16)
			if err != nil {
				return err
			}
			values = append(values, tokenValue{
				hexValue: hex,
			})
		case nd.Binary:
			hex, err := convertToHex(n.Value[1:], 2)
			if err != nil {
				return err
			}
			values = append(values, tokenValue{
				hexValue: hex,
			})
		case nd.Octal:
			hex, err := convertToHex(n.Value[1:], 8)
			if err != nil {
				return err
			}
			values = append(values, tokenValue{
				hexValue: hex,
			})
		case nd.String:
			values = append(values, tokenValue{
				value: n.Value,
			})
		default:
			return fmt.Errorf("unknown token child: %v", nd.NodeTypes[n.Type])
		}
	}

	token.rawValues = values
	if len(values) == 1 {
		tk := values[0]
		if tk.isString() {
			// Strings: "value"
			token.value = fmt.Sprintf("%q", tk.value)
		} else {
			token.value, _ = convertToRuneString(tk.hexValue, 16)
		}
	} else {
		for _, tk := range values {
			if tk.isString() {
				token.value += tk.value
			} else {
				v, _ := convertToInt(tk.hexValue, 16)
				token.value += string(rune(v))
			}
		}
		token.value = fmt.Sprintf("%q", token.value)
	}
	g.tokens = append(g.tokens, token)
	return nil
}

// generateTokens writes all the tokens to the given writer.
func (g *generator) generateTokens() error {
	w := g.writers["tk"]

	// Check whether the grammar has tokens.
	var containsTokens bool
	if len(g.tokens) != 0 {
		containsTokens = true
	}
	for _, dep := range g.dependencies {
		if len(dep.tokens) != 0 {
			containsTokens = true
		}
	}
	if !containsTokens {
		return nil
	}

	w.c("Token Definitions")
	w.wln("const (")
	{
		w := w.indent()

		if len(g.tokens) != 0 {
			w.cf("%s (%s)\n", g.languageFull(), g.meta.url)
			g.generateTokenValues(w)
		}

		for i, dep := range g.dependencies {
			if len(dep.tokens) == 0 {
				continue
			}
			switch i {
			case 0:
				if len(g.tokens) != 0 {
					w.ln()
				}
			default:
				if len(g.dependencies[i-1].tokens) != 0 {
					w.ln()
				}
			}
			w.cf("%s (%s)\n", dep.languageFull(), dep.meta.url)
			dep.generateTokenValues(w)
		}
	}
	w.wln(")")
	return nil
}

// generateTokenValues is responsible for generating token values.
func (g *generator) generateTokenValues(w *writer) {
	longestName := g.longestTokenName()
	for idx, token := range g.tokens {
		if token.comment != "" {
			longestValue := g.longestTokenValueWithComment(idx)
			token.value = fillRight(token.value, longestValue)
		}
		w.wf("%s = %s", fillRight(token.name, longestName), token.value)
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
