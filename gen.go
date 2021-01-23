// +build ignore

// This program generates the pegn sub-package. It can be invoked by running the
// go generate command.
package main

import (
	"github.com/pegn/pegn-go"
	"log"
)

func main() {
	if err := pegn.GenerateFromURLs("pegn/", pegn.Config{
		ModulePath:      "github.com/pegn/pegn-go",
		IgnoreReserved:  true,
		ClassSubPackage: "is",
		TokenSubPackage: "tk",
		TypeSubPackage:  "nd",
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
		log.Fatal(err)
	}
	log.Println("Successfully generated the pegn sub-module.")
}
