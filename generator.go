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
}

type Generator struct {
	root    *pegn.Node         // root node of the PEGN grammar.
	config  Config             // config of the parser.
	writers map[string]*writer // a map of writers to separate generated code.
	errors  []error            // list of errors that occurred.

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

	nodes   []node
	scans   []scan
	classes []class
	tokens  []token
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

func (g *Generator) Generate() {
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
			g.parseNode(n)
		case nd.ScanDef:
			g.parseScan(n)
		case nd.ClassDef:
			g.parseClass(n)
		case nd.TokenDef:
			g.parseToken(n)
		default:
			g.errors = append(g.errors, fmt.Errorf("unknown definition child: %v", n.Types[n.Type]))
		}
	}

	g.generate()
}

func (g *Generator) generate() {
	g.generateTokens()
	g.generateTypes()
	g.generateClasses()
}

func (g *Generator) generateHeader(w *writer) {
	w.c("Do not edit. This file is auto-generated.")
	w.wlnf("package %s", strings.ToLower(g.meta.language))
	w.ln()
}
