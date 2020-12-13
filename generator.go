package pegen

import (
	"fmt"
	"gitlab.com/pegn/pegn-go"
	"gitlab.com/pegn/pegn-go/ast"
	"gitlab.com/pegn/pegn-go/nd"
	"io"
	"strconv"
	"strings"
)

type Config struct {
	IgnoreReserved bool
}

type Generator struct {
	root   *pegn.Node
	config Config
	writer writer
	errors []error

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

func New(rawGrammar interface{}, output io.Writer, config Config) (Generator, error) {
	p := new(pegn.Parser)
	if err := p.Init(rawGrammar); err != nil {
		return Generator{}, err
	}
	grammar, err := ast.Grammar(p)
	if err != nil {
		return Generator{}, err
	}
	if !p.Done() {
		return Generator{}, fmt.Errorf("parser could not read the entire file")
	}

	return Generator{
		root: grammar,
		writer: writer{
			Writer: output,
		},
		config: config,
	}, nil
}

func (g *Generator) Generate() {
	for _, node := range g.root.Children() {
		switch node.Type {
		case nd.Comment, nd.EndLine:
			// Ignore these.
		case nd.Meta:
			// Assign meta data.
			for _, node := range node.Children() {
				switch node.Type {
				case nd.Lang:
					g.meta.language = node.Value
				case nd.MajorVer:
					g.meta.version.major, _ = strconv.Atoi(node.Value)
				case nd.MinorVer:
					g.meta.version.minor, _ = strconv.Atoi(node.Value)
				case nd.PatchVer:
					g.meta.version.patch, _ = strconv.Atoi(node.Value)
				case nd.PreVer:
					g.meta.version.prerelease = node.Value
				case nd.Home:
					g.meta.url = node.Value
				}
			}
		case nd.Copyright:
			for _, node := range node.Children() {
				if node.Type == nd.Comment {
					g.copyright = node.Value
					break
				}
			}
		case nd.Licensed:
			for _, node := range node.Children() {
				if node.Type == nd.Comment {
					g.license = node.Value
					break
				}
			}
		case nd.NodeDef:
			g.parseNode(node)
		case nd.ScanDef:
			g.parseScan(node)
		case nd.ClassDef:
			g.parseClass(node)
		case nd.TokenDef:
			g.parseToken(node)
		default:
			// TODO fmt.Println(node.Types[node.Type])
		}
	}

	g.generate()
}

func (g *Generator) generate() {
	w := g.writer

	w.c("Do not edit. This file is auto-generated.")
	w.wlnf("package %s", strings.ToLower(g.meta.language))
	w.ln()
	g.generateTokens()
	w.ln()
	g.generateTypes()
}
