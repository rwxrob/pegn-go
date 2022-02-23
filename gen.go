//go:build ignore
// +build ignore

// This program generates the pegn sub-package. It can be invoked by
// running the `go generate` command.
package main

import (
	"log"

	"github.com/rwxrob/pegn-go"
)

func main() {
	if err := pegn.GenerateFromURLs("pegn/", pegn.Config{
		ModulePath:      "github.com/rwxrob/pegn-go",
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
		"https://pegn.dev/spec/types.pegn",
		"https://pegn.dev/spec/classes.pegn",
		"https://pegn.dev/spec/tokens.pegn",
	}...); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully generated the pegn sub-module.")
}
