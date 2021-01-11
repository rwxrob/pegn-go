package pegen

import (
	"fmt"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Config that offers additional configuration options to the generator.
type Config struct {
	// GrammarLocations are aliases to the grammar locations.
	// e.g. 'spec.pegn.dev': 'https://raw.githubusercontent.com/pegn/spec/master/grammar.pegn'
	GrammarLocations map[string]string
	// IgnoreReserved makes sure the generator does not return a error when it
	// tries to generate a reserved class/token. False by default, you should
	// not overwrite reserved classes or tokens.
	IgnoreReserved bool
	// TokenPrefix represents the prefix before a token name. No prefix is added
	// when the value is an empty string.
	// e.g. 'SP' with prefix 'Token' results in 'TokenSP'.
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
	root    *ast.Node          // root node of the PEGN grammar.
	config  Config             // config of the parser.
	writers map[string]*writer // a map of writers to separate generated code.

	dependencies   []Generator
	dependencyURLs []string

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

func GenerateFromURLs(outputDir string, config Config, urls ...string) error {
	files := make([][]byte, len(urls))
	for i, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		files[i] = raw
	}
	return GenerateFromFiles(outputDir, config, files[0], files[1:]...)
}

func GenerateFromFiles(outputDir string, config Config, grammar []byte, dependencies ...[]byte) error {
	parentDir := filepath.Dir(outputDir)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, os.ModePerm); err != nil {
			return err
		}
	}

	g, err := New(grammar, config)
	if err != nil {
		return err
	}

	var (
		generators []Generator
		urls       = make(map[string]bool)
	)
	for _, file := range dependencies {
		g, err := New(file, config)
		if err != nil {
			return err
		}
		if err := g.GenerateBuffers(); err != nil {
			return err
		}
		generators = append(generators, g)
		urls[g.meta.url] = true
	}

	// Check whether the grammar files has all its dependencies.
	for _, dep := range g.dependencyURLs {
		if _, ok := urls[dep]; !ok {
			return fmt.Errorf("missing dependency: %s", dep)
		}
	}
	// Check whether all the dependencies have all their dependencies.
	for _, dep := range generators {
		for _, url := range g.dependencyURLs {
			if _, ok := urls[url]; !ok {
				return fmt.Errorf("missing dependency: %s", url)
			}
		}
		g.dependencies = append(g.dependencies, dep)
	}

	// Make sure to do this after adding all the dependencies.
	if err := g.GenerateBuffers(); err != nil {
		return err
	}

	w, b := newBW()
	g.generateHeader(w)
	w.ln()
	w.wln("import (")
	{
		w := w.indent()
		w.wln("\"github.com/di-wu/parser\"")
		w.wln("\"github.com/di-wu/parser/ast\"")
		w.wln("\"github.com/di-wu/parser/op\"")
	}
	w.wln(")")
	w.ln()
	w.w(g.writers["ast"].String())
	w.w(g.writers["is"].String())
	w.ln()
	w.w(g.writers["tk"].String())
	w.ln()
	w.w(g.writers["nd"].String())

	return ioutil.WriteFile(fmt.Sprintf("%s/grammar.go", parentDir), b.Bytes(), os.ModePerm)
}

func New(rawGrammar []byte, config Config) (Generator, error) {
	// 1. Create a new PEGN parser.
	p, err := ast.New(rawGrammar)
	if err != nil {
		return Generator{}, err
	}
	// 2. Parse the given grammar.
	grammar, err := pegn.Spec(p)
	if err != nil {
		return Generator{}, err
	}
	// 3. Check whether the whole file is parsed correctly.
	if _, err := p.Expect(parser.EOD); err != nil {
		return Generator{}, fmt.Errorf("parser could not read the entire file")
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

func (g *Generator) GenerateBuffers() error {
	for _, n := range g.root.Children() {
		switch n.Type {
		case pegn.CommentType, pegn.EndLineType:
			// Ignore these.
		case pegn.MetaType:
			// Meta <-- '# ' Language ' (' Version ') ' Home EndLine
			// Language <- Lang ('-' LangExt)?
			// Version <- 'v' MajorVer '.' MinorVer '.' PatchVer ('-' PreVer)?
			// Home <-- (!ws unipoint)+
			for _, n := range n.Children() {
				switch n.Type {
				case pegn.NameType:
					g.meta.language = n.ValueString()
				case pegn.NameExtType:
					g.meta.language += fmt.Sprintf("-%s", n.ValueString())
				case pegn.MajorVerType:
					g.meta.version.major, _ = strconv.Atoi(n.ValueString())
				case pegn.MinorVerType:
					g.meta.version.minor, _ = strconv.Atoi(n.ValueString())
				case pegn.PatchVerType:
					g.meta.version.patch, _ = strconv.Atoi(n.ValueString())
				case pegn.PreVerType:
					g.meta.version.prerelease = n.ValueString()
				case pegn.HomeType:
					g.meta.url = n.ValueString()
				}
			}
		case pegn.CopyrightType:
			// Copyright <-- '# Copyright ' Comment EndLine
			for _, n := range n.Children() {
				if n.Type == pegn.CommentType {
					g.copyright = n.ValueString()
					break
				}
			}
		case pegn.LicensedType:
			// Licensed <-- '# Licensed under ' Comment EndLine
			for _, n := range n.Children() {
				if n.Type == pegn.CommentType {
					g.license = n.ValueString()
					break
				}
			}
		case pegn.UsesType:
			for _, n := range n.Children() {
				if n.Type == pegn.PathType {
					g.dependencyURLs = append(g.dependencyURLs, n.ValueString())
					break
				}
			}
		// Definition
		// Definition <- NodeDef / ScanDef / ClassDef / TokenDef
		case pegn.NodeDefType:
			if err := g.parseNode(n); err != nil {
				return err
			}
		case pegn.ScanDefType:
			if err := g.parseScan(n); err != nil {
				return err
			}
		case pegn.ClassDefType:
			if err := g.parseClass(n); err != nil {
				return err
			}
		case pegn.TokenDefType:
			if err := g.parseToken(n); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown definition child: %v", pegn.NodeTypes[n.Type])
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
	if err := g.generateClasses(g.writers["is"]); err != nil {
		return err
	}
	if err := g.generateNodes(g.writers["ast"]); err != nil {
		return err
	}
	return nil
}

func (g *Generator) generateHeader(w *writer) {
	w.c("Do not edit. This file is auto-generated.")
	w.wlnf("package %s", strings.ToLower(g.meta.language))
}
