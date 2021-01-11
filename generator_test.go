package pegen

import (
	"testing"
)

func TestGenerateFromURLs(t *testing.T) {
	if err := GenerateFromURLs("pegn/", Config{
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
	}...); err != nil {
		t.Error(err)
		return
	}
}
