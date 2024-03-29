package pegn

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/rwxrob/pegn-go/pegn"
	"github.com/rwxrob/pegn-go/pegn/nd"
)

//go:generate go run gen.go

// Config that offers additional configuration options to the generator.
type Config struct {
	// ModulePath is the path to the module. REQUIRED!
	ModulePath string
	// MuduleName is the name of the module. Uses the meta.language field by default.
	ModuleName string

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
	// TokenSubPackage is the name of the sub-package for tokens. It left empty
	// the tokens will get added to the main file.
	TokenSubPackage string

	// TypeSuffix represent the pre- or suffix before a class name.
	// Recommended. Type names are the same as there corresponding parse
	// functions. e.g. 'Grammar' with suffix 'Type' results in 'GrammarType'.
	TypeSuffix string
	// TypeSubPackage is the name of the sub-package for types. It left empty
	// the types will get added to the main file.
	TypeSubPackage string

	// ClassAliases is a map of original class names to an alias.
	// e.g. map[alphanum: AlphaNum]
	ClassAliases map[string]string
	// ClassSubPackage is the name of the sub-package for classes. It left empty
	// the classes will get added to the main file.
	ClassSubPackage string

	// NodeAliases is a map of original node names to an alias.
	// e.g. map[Hexadec: Hexadecimal]
	NodeAliases map[string]string
}

// GenerateFromURLs generates Go code to the given outputDir based on the given
// configuration an urls.
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

// GenerateFromFiles generates Go code to the given outputDir based on the given
// configuration an files.
func GenerateFromFiles(outputDir string, config Config, grammar []byte, dependencies ...[]byte) error {
	parentDir := filepath.Dir(outputDir)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, os.ModePerm); err != nil {
			return err
		}
	}

	g, err := newFromFiles(config, grammar, dependencies...)
	if err != nil {
		return err
	}

	// Make sure to do this after adding all the dependencies.
	if err := g.generateBuffers(); err != nil {
		return err
	}

	if config.TypeSubPackage == "" && config.TypeSuffix == "" {
		return fmt.Errorf("a type suffix should be defined when there is no type sub-package")
	}

	w, b := newBW()
	g.generateHeader(w)
	w.ln()
	w.wln("import (")
	{
		w := w.indent()

		imports := []string{
			"github.com/di-wu/parser/ast",
			"github.com/di-wu/parser/op",
		}
		if g.config.TypeSubPackage != "" {
			imports = append(imports, fmt.Sprintf(
				"%s/%s/%s",
				g.config.ModulePath, parentDir, g.config.TypeSubPackage,
			))
		}
		if g.config.TokenSubPackage != "" {
			imports = append(imports, fmt.Sprintf(
				"%s/%s/%s",
				g.config.ModulePath, parentDir, g.config.TokenSubPackage,
			))
		}
		if g.config.ClassSubPackage != "" {
			imports = append(imports, fmt.Sprintf(
				"%s/%s/%s",
				g.config.ModulePath, parentDir, g.config.ClassSubPackage,
			))
		} else {
			imports = append(imports, "github.com/di-wu/parser")
		}

		sort.Strings(imports)
		for _, i := range imports {
			w.wlnf("%q", i)
		}
	}
	w.wln(")")
	w.ln()
	w.w(g.writers["ast"].String())
	if classes := g.writers["is"].String(); len(classes) != 0 {
		if g.config.ClassSubPackage == "" {
			w.ln()
			w.w(g.writers["is"].String())
		} else {
			if err := g.generateClassFile(parentDir); err != nil {
				return err
			}
		}
	}
	if tokens := g.writers["tk"].String(); len(tokens) != 0 {
		if g.config.TokenSubPackage == "" {
			w.ln()
			w.w(tokens)
		} else {
			if err := g.generateTokenFile(parentDir); err != nil {
				return err
			}
		}
	}
	if g.config.TypeSubPackage == "" {
		w.ln()
		w.w(g.writers["nd"].String())
	} else {
		if err := g.generateTypeFile(parentDir); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(
		fmt.Sprintf("%s/grammar.go", parentDir),
		b.Bytes(), os.ModePerm,
	)
}

// newFromFiles creates generator(s) based on the given configuration and files.
func newFromFiles(config Config, grammar []byte, dependencies ...[]byte) (generator, error) {
	g, err := newGenerator(grammar, config)
	if err != nil {
		return generator{}, err
	}

	var (
		generators []generator
		urls       = make(map[string]bool)
	)
	for _, file := range dependencies {
		g, err := newGenerator(file, config)
		if err != nil {
			return generator{}, err
		}
		if err := g.generateBuffers(); err != nil {
			return generator{}, err
		}
		generators = append(generators, g)
		urls[g.meta.url] = true
	}

	// Check whether the grammar files has all its dependencies.
	for _, dep := range g.dependencyURLs {
		if _, ok := urls[dep]; !ok {
			return generator{}, fmt.Errorf("missing dependency: %s", dep)
		}
	}
	// Check whether all the dependencies have all their dependencies.
	for _, dep := range generators {
		for _, url := range g.dependencyURLs {
			if _, ok := urls[url]; !ok {
				return generator{}, fmt.Errorf("missing dependency: %s", url)
			}
		}
		g.dependencies = append(g.dependencies, dep)
	}

	return g, nil
}

// generateBuffers writes the pre-parsed data to different buffers.
func (g *generator) generateBuffers() error {
	if err := g.prepare(); err != nil {
		return err
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

// prepare does all the pre-parsing.
func (g *generator) prepare() error {
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
				case nd.Name:
					g.meta.language = n.Value
				case nd.NameExt:
					g.meta.languageSuffix = n.Value
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
		case nd.Uses:
			for _, n := range n.Children() {
				if n.Type == nd.Path {
					g.dependencyURLs = append(g.dependencyURLs, n.Value)
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
			return fmt.Errorf("unknown definition child: %v", nd.NodeTypes[n.Type])
		}
	}
	return nil
}

// generator does all the file generation.
type generator struct {
	root    *ast.Node          // root node of the PEGN grammar.
	config  Config             // config of the parser.
	writers map[string]*writer // a map of writers to separate generated code.

	dependencies   []generator
	dependencyURLs []string

	// meta data of the grammar.
	meta struct {
		language       string
		languageSuffix string
		version        struct {
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

func (g generator) languageFull() string {
	if g.meta.languageSuffix == "" {
		return g.meta.language
	}
	return fmt.Sprintf("%s-%s", g.meta.language, g.meta.languageSuffix)
}

// newGenerator returns a new configured generator.
func newGenerator(rawGrammar []byte, config Config) (generator, error) {
	// 1. Create a new PEGN parser.
	p, err := ast.New(rawGrammar)
	if err != nil {
		return generator{}, err
	}
	// 2. Parse the given grammar.
	grammar, err := pegn.Spec(p)
	if err != nil {
		return generator{}, err
	}
	// 3. Check whether the whole file is parsed correctly.
	if _, err := p.Expect(parser.EOD); err != nil {
		return generator{}, fmt.Errorf("parser could not read the entire file: %s", err.Error())
	}

	return generator{
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

// generateHeader writes the header of the main file.
func (g *generator) generateHeader(w *writer) {
	w.c("Do not edit. This file is auto-generated.")
	version := fmt.Sprintf("%d.%d.%d", g.meta.version.major, g.meta.version.minor, g.meta.version.patch)
	if pre := g.meta.version.prerelease; pre != "" {
		version = fmt.Sprintf("%s-%s", version, pre)
	}
	w.cf(
		"Grammar: %s (v%s) %s",
		g.languageFull(), version, g.meta.url,
	)
	w.ln()
	w.wlnf("package %s", strings.ToLower(g.moduleName()))
}

// generateClassFile is responsible for generating a standalone class file.
func (g *generator) generateClassFile(parentDir string) error {
	pkg := strings.ToLower(g.config.ClassSubPackage)
	dir := fmt.Sprintf("%s/%s", parentDir, pkg)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	w, b := newBW()
	w.c("Do not edit. This file is auto-generated.")
	w.ln()
	w.wlnf("package %s", strings.ToLower(pkg))
	w.ln()
	w.wln("import (")
	{
		w := w.indent()

		imports := []string{
			"github.com/di-wu/parser",
			"github.com/di-wu/parser/op",
		}

		if g.config.TokenSubPackage != "" {
			imports = append(imports, fmt.Sprintf(
				"%s/%s/%s",
				g.config.ModulePath, parentDir, g.config.TokenSubPackage,
			))
		}

		sort.Strings(imports)
		for _, i := range imports {
			w.wlnf("%q", i)
		}
	}
	w.wln(")")
	w.ln()
	w.w(g.writers["is"].String())

	if err := ioutil.WriteFile(
		fmt.Sprintf("%s/classes.go", dir),
		b.Bytes(), os.ModePerm,
	); err != nil {
		return err
	}
	return nil
}

// generateTokenFile is responsible for generating a standalone token file.
func (g *generator) generateTokenFile(parentDir string) error {
	pkg := strings.ToLower(g.config.TokenSubPackage)
	dir := fmt.Sprintf("%s/%s", parentDir, pkg)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	w, b := newBW()
	w.c("Do not edit. This file is auto-generated.")
	w.ln()
	w.wlnf("package %s", strings.ToLower(pkg))
	w.ln()
	w.w(g.writers["tk"].String())

	if err := ioutil.WriteFile(
		fmt.Sprintf("%s/tokens.go", dir),
		b.Bytes(), os.ModePerm,
	); err != nil {
		return err
	}
	return nil
}

// generateTypeFile is responsible for generating a standalone types file.
func (g *generator) generateTypeFile(parentDir string) error {
	pkg := strings.ToLower(g.config.TypeSubPackage)
	dir := fmt.Sprintf("%s/%s", parentDir, pkg)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	w, b := newBW()
	w.c("Do not edit. This file is auto-generated.")
	w.ln()
	w.wlnf("package %s", strings.ToLower(pkg))
	w.ln()
	w.w(g.writers["nd"].String())

	if err := ioutil.WriteFile(
		fmt.Sprintf("%s/types.go", dir),
		b.Bytes(), os.ModePerm,
	); err != nil {
		return err
	}
	return nil
}
