package pegen

import (
	"github.com/di-wu/parser"
	"testing"
)

func TestParserFromURLs(t *testing.T) {
	p, err := ParserFromURLs(Config{
		IgnoreReserved: true,
		TypeSuffix:     "Type",
		ClassAliases: map[string]string{
			"alphanum": "AlphaNum",
			"unipoint": "UniPoint",
			"bindig":   "BinDig",
			"hexdig":   "HexDig",
			"lowerhex": "LowerHex",
			"octdig":   "OctDig",
			"uphex":    "UpHex",
			"ws":       "Whitespace",
			"ascii":    "ASCII",
		},
		NodeAliases: map[string]string{
			"Hexadec": "Hexadecimal",
		},
	}, []string{
		"https://raw.githubusercontent.com/pegn/spec/master/grammar.pegn",
		"https://raw.githubusercontent.com/pegn/spec/master/classes/grammar.pegn",
		"https://raw.githubusercontent.com/pegn/spec/master/tokens/grammar.pegn",
	}...)
	if err != nil {
		t.Error(err)
		return
	}

	alpha := p.classes["Alpha"]
	ap, _ := parser.New([]byte("abc"))

	m, err := ap.Expect(alpha)
	for {
		tmp, err := ap.Expect(alpha)
		if err != nil {
			break
		}
		m = tmp
	}
	if m == nil {
		t.Error(err)
		return
	}
	if m.Rune != 'c' {
		t.Errorf(m.String())
	}
}
