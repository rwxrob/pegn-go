package pegen

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/ast"
	"gitlab.com/pegn/pegn-go/nd"
)

// Config that offers additional configuration options to the generator.
type Config struct {
	// IgnoreReserved makes sure the generator does not return a error when it
	// tries to generate a reserved class/token. False by default, you should
	// not overwrite reserved classes or tokens.
	IgnoreReserved bool
	// TokenPrefix represents the prefix before a token name. No prefix is added
	// when the value is an empty string. This can also be used if the tokens
	// are located in a submodule.
	// e.g. 'SP' with prefix 'tk.' results in 'tk.SP'.
	TokenPrefix string
	// TypePrefix / TypeSuffix represent the pre- or suffix before a class name.
	// Recommended. Type names are the same as there corresponding parse
	// functions. e.g. 'Grammar' with suffix 'Type' results in 'GrammarType'.
	TypePrefix, TypeSuffix string
	// ClassAliases is a map of original class names to an alias.
	// e.g. map[alphanum: AlphaNum]
	ClassAliases map[string]string
	// NodeAliases is a map of original node names to an alias.
	// e.g. map[Hexadec: Hexadecimal]
	NodeAliases map[string]string
}

type Generator struct {
	root    *pegn.Node         // root node of the PEGN grammar.
	config  Config             // config of the parser.
	writers map[string]*writer // a map of writers to separate generated code.

	// meta data of the grammar.
	meta struct {
		language string
		version  struct {
			major, minor, patch int
			prerelease          string
		}
		url string
	}
	copyright string
	license   string

	nodes   nodes
	classes []class
	tokens  tokens
}

func New(rawGrammar interface{}, parentDir string, config Config) (Generator, error) {
	// 1. Create a new PEGN parser.
	p := new(pegn.Parser)
	if err := p.Init(rawGrammar); err != nil {
		return Generator{}, err
	}
	// 2. Parse the given grammar.
	grammar, err := ast.Grammar(p)
	if err != nil {
		return Generator{}, err
	}
	// 3. Check whether the whole file is parsed correctly.
	if !p.Done() {
		return Generator{}, fmt.Errorf("parser could not read the entire file")
	}

	// 4. Ensure if parentDir exists.
	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		if err := os.Mkdir(parentDir, os.ModePerm); err != nil {
			return Generator{}, err
		}
	}

	return Generator{
		root: grammar,
		writers: map[string]*writer{
			"tk":  newW(), // tokens
			"nd":  newW(), // types
			"is":  newW(), // classes
			"ast": newW(), // ast nodes
		},
		config: config,
	}, nil
}

func (g *Generator) Generate() error {
	for _, n := range g.root.Children() {
		switch n.Type {
		case nd.Comment, nd.EndLine:
			// Ignore these.
		case nd.Meta:
			// Meta <-- '# ' Language ' (' Version ') ' Home EndLine
			// Language <- Lang ('-' LangExt)?
			// Version <- 'v' MajorVer '.' MinorVer '.' PatchVer ('-' PreVer)?
			// Home <-- (!ws unipoint)+
			for _, n := range n.Children() {
				switch n.Type {
				case nd.Lang:
					g.meta.language = n.Value
				case nd.MajorVer:
					g.meta.version.major, _ = strconv.Atoi(n.Value)
				case nd.MinorVer:
					g.meta.version.minor, _ = strconv.Atoi(n.Value)
				case nd.PatchVer:
					g.meta.version.patch, _ = strconv.Atoi(n.Value)
				case nd.PreVer:
					g.meta.version.prerelease = n.Value
				case nd.Home:
					g.meta.url = n.Value
				}
			}
		case nd.Copyright:
			// Copyright <-- '# Copyright ' Comment EndLine
			for _, n := range n.Children() {
				if n.Type == nd.Comment {
					g.copyright = n.Value
					break
				}
			}
		case nd.Licensed:
			// Licensed <-- '# Licensed under ' Comment EndLine
			for _, n := range n.Children() {
				if n.Type == nd.Comment {
					g.license = n.Value
					break
				}
			}
		// Definition
		// Definition <- NodeDef / ScanDef / ClassDef / TokenDef
		case nd.NodeDef:
			if err := g.parseNode(n); err != nil {
				return err
			}
		case nd.ScanDef:
			if err := g.parseScan(n); err != nil {
				return err
			}
		case nd.ClassDef:
			if err := g.parseClass(n); err != nil {
				return err
			}
		case nd.TokenDef:
			if err := g.parseToken(n); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown definition child: %v", n.Types[n.Type])
		}
	}

	// The order is important!
	// g.generateNodes() for example relies on the (pre)generated tokens.
	if err := g.generateTokens(); err != nil {
		return err
	}
	if err := g.generateTypes(); err != nil {
		return err
	}
	if err := g.generateClasses(); err != nil {
		return err
	}
	if err := g.generateNodes(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) generateHeader(w *writer) {
	w.c("Do not edit. This file is auto-generated.")
	w.wlnf("package %s", strings.ToLower(g.meta.language))
	w.ln()
}
