package pegn

import (
	"fmt"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
	"github.com/pegn/pegn-go/pegn/nd"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

// InMemoryParser is the in-memory equivalent of generated code of generator.
type InMemoryParser struct {
	tokens  map[string]interface{}
	classes map[string]interface{}
	nodes   map[string]interface{}
}

// GetTokenDef returns a token definition.
func (i *InMemoryParser) GetTokenDef(id string) (interface{}, error) {
	if def, ok := i.tokens[id]; ok {
		return def, nil
	}
	return nil, fmt.Errorf("invalid token identifier: %s", id)
}

// GetClassDef returns a class definition.
func (i *InMemoryParser) GetClassDef(id string) (interface{}, error) {
	if def, ok := i.classes[id]; ok {
		return def, nil
	}
	return nil, fmt.Errorf("invalid class identifier: %s", id)
}

// GetNodeDef returns a node definition.
func (i *InMemoryParser) GetNodeDef(id string) (interface{}, error) {
	if def, ok := i.nodes[id]; ok {
		return def, nil
	}
	return nil, fmt.Errorf("invalid node identifier: %s", id)
}

type internalParser struct {
	sync.WaitGroup
	sync.RWMutex
	InMemoryParser
}

// ParserFromURLs creates an in-memory parser based on the given configuration
// and urls.
func ParserFromURLs(config Config, urls ...string) (InMemoryParser, error) {
	files := make([][]byte, len(urls))
	for i, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return InMemoryParser{}, err
		}
		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return InMemoryParser{}, err
		}
		files[i] = raw
	}
	return ParserFromFiles(config, files[0], files[1:]...)
}

// ParserFromFiles creates an in-memory parser based on the given configuration
// and files.
func ParserFromFiles(config Config, grammar []byte, dependencies ...[]byte) (InMemoryParser, error) {
	g, err := newFromFiles(config, grammar, dependencies...)
	if err != nil {
		return InMemoryParser{}, err
	}

	if err := g.prepare(); err != nil {
		return InMemoryParser{}, err
	}

	p := internalParser{
		InMemoryParser: InMemoryParser{
			tokens:  make(map[string]interface{}),
			classes: make(map[string]interface{}),
			nodes:   make(map[string]interface{}),
		},
	}

	if err := p.generateTokens(g); err != nil {
		return InMemoryParser{}, err
	}
	if err := p.generateClasses(g); err != nil {
		return InMemoryParser{}, err
	}
	if err := p.generateNodes(g); err != nil {
		return InMemoryParser{}, err
	}

	return p.InMemoryParser, nil
}

// generateTokens is responsible for generation the tokens.
func (p *internalParser) generateTokens(g generator) error {
	for _, token := range g.tokens {
		if _, ok := p.tokens[token.name]; ok {
			return fmt.Errorf("duplicate token: %s", token.name)
		}

		var and op.And
		for _, tk := range token.rawValues {
			if tk.isString() {
				and = append(and, tk.value)
			} else {
				i, _ := convertToInt(tk.hexValue, 16)
				and = append(and, i)
			}
		}
		p.RWMutex.Lock()
		if len(and) == 1 {
			p.tokens[token.name] = and[0]
		} else {
			p.tokens[token.name] = and
		}
		p.RWMutex.Unlock()
	}
	for _, dep := range g.dependencies {
		if err := p.generateTokens(dep); err != nil {
			return err
		}
	}
	return nil
}

// generateClasses is responsible for generation the classes.
func (p *internalParser) generateClasses(g generator) error {
	var wg sync.WaitGroup
	errChannel := make(chan error, 1)
	generate := func(g generator) {
		defer wg.Done()
		for _, class := range g.classes {
			var or op.Or
			for _, n := range class.expression {
				p.WaitGroup.Add(1)
				go func(n *ast.Node) {
					defer p.WaitGroup.Done()

					switch n.Type {
					case nd.Comment, nd.EndLine:
						// Ignore these.
						return
					case nd.Unicode, nd.Hexadecimal:
						i, _ := convertToInt(n.Value[1:], 16)
						or = append(or, i)
					case nd.Binary:
						i, _ := convertToInt(n.Value[1:], 2)
						or = append(or, i)
					case nd.Octal:
						i, _ := convertToInt(n.Value[1:], 8)
						or = append(or, i)
					case nd.ClassId, nd.ResClassId:
						name := g.className(n.Value)
						or = append(or, ast.LoopUp{
							Key:   name,
							Table: &p.nodes,
						})
					case nd.TokenId, nd.ResTokenId:
						name := g.tokenName(n.Value)

						p.RWMutex.RLock()
						if token, ok := p.tokens[name]; ok {
							or = append(or, token)
						}
						p.RWMutex.RUnlock()

						var wg sync.WaitGroup
						wg.Add(1)
						go func(name string) {
							defer wg.Done()
							for {
								p.RWMutex.RLock()
								v := p.tokens[name]
								p.RWMutex.RUnlock()
								if v != nil {
									break
								}
							}
						}(name)
						wg.Wait()

						p.RWMutex.RLock()
						or = append(or, p.tokens[name])
						p.RWMutex.RUnlock()
					case nd.AlphaRange, nd.IntRange:
						min := rune(n.Children()[0].Value[0])
						max := rune(n.Children()[1].Value[0])
						or = append(or, parser.CheckRuneRange(min, max))
					case nd.UniRange, nd.HexRange:
						min, _ := convertToInt(n.Children()[0].Value[1:], 16)
						max, _ := convertToInt(n.Children()[1].Value[1:], 16)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case nd.BinRange:
						min, _ := convertToInt(n.Children()[0].Value[1:], 2)
						max, _ := convertToInt(n.Children()[1].Value[1:], 2)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case nd.OctRange:
						min, _ := convertToInt(n.Children()[0].Value[1:], 8)
						max, _ := convertToInt(n.Children()[1].Value[1:], 8)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case nd.String:
						if v := n.Value; len(v) == 1 {
							or = append(or, parser.CheckRune([]rune(v)[0]))
						} else {
							or = append(or, parser.CheckString(v))
						}
					default:
						errChannel <- fmt.Errorf("unknown class child: %v", nd.NodeTypes[n.Type])
					}
				}(n)
			}
			p.Wait()

			p.RWMutex.Lock()
			if len(or) == 1 {
				p.classes[class.name] = or[0]
			} else {
				p.classes[class.name] = or
			}
			p.RWMutex.Unlock()
		}
	}

	wg.Add(1)
	go generate(g)
	for _, dep := range g.dependencies {
		wg.Add(1)
		go generate(dep)
	}

	finished := make(chan bool, 1)
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChannel:
		if err != nil {
			return err
		}
	}
	return nil
}

// generateNodes is responsible for generation the nodes.
func (p *internalParser) generateNodes(g generator) error {
	var wg sync.WaitGroup
	errChannel := make(chan error, 1)
	generate := func(g generator, node node) {
		defer wg.Done()

		i, err := p.generateExpression(g, node.expression)
		if err != nil {
			errChannel <- err
		}
		p.RWMutex.Lock()
		p.nodes[g.nodeName(node.name)] = i
		p.RWMutex.Unlock()
	}

	for _, node := range g.nodes {
		wg.Add(1)
		go generate(g, node)
	}
	for _, dep := range g.dependencies {
		for _, node := range dep.nodes {
			wg.Add(1)
			go generate(dep, node)
		}
	}

	finished := make(chan bool, 1)
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChannel:
		if err != nil {
			return err
		}
	}
	return nil
}

// generateExpression is responsible for generation the expressions.
func (p *internalParser) generateExpression(g generator, expression []*ast.Node) (interface{}, error) {
	var or op.Or
	for _, n := range expression {
		switch n.Type {
		case nd.Comment, nd.EndLine:
			// Ignore these.
			continue

		// Sequence <-- Rule (Spacing Rule)*
		case nd.Sequence:
			i, err := p.generateSequence(g, n.Children())
			if err != nil {
				return nil, err
			}
			or = append(or, i)
		default:
			return nil, fmt.Errorf("unknown expression child: %v", nd.NodeTypes[n.Type])
		}
	}
	if len(or) == 1 {
		return or[0], nil
	}
	return or, nil
}

// generateSequence is responsible for generation the sequences.
func (p *internalParser) generateSequence(g generator, sequence []*ast.Node) (interface{}, error) {
	var and op.And
	for _, n := range sequence {
		switch n.Type {
		case nd.Comment, nd.EndLine:
			// Ignore these.
			continue
		// Plain <-- Primary Quant?
		case nd.Plain:
			// Plain <-- Primary Quant?
			var quant *ast.Node
			switch last := n.Children()[len(n.Children())-1]; last.Type {
			case nd.Optional, nd.MinZero, nd.MinOne, nd.MinMax, nd.Count:
				quant = last
			}
			if quant == nil {
				i, err := p.generatePrimary(g, n.Children()[0])
				if err != nil {
					return nil, err
				}
				if i != nil {
					and = append(and, i)
				}
			} else {
				i, err := p.generatePrimary(g, n.Children()[0])
				if err != nil {
					return nil, err
				}
				if i == nil {
					// This can occur if the node is a comment or eol.
					continue
				}
				switch q := n.Children()[1]; q.Type {
				case nd.Optional:
					and = append(and, op.Optional(i))
				case nd.MinZero:
					and = append(and, op.MinZero(i))
				case nd.MinOne:
					and = append(and, op.MinOne(i))
				case nd.MinMax:
					min, _ := strconv.Atoi(q.Children()[0].Value)
					max := -1
					if len(q.Children()) == 2 {
						max, _ = strconv.Atoi(q.Children()[1].Value)
					}
					and = append(and, op.MinMax(min, max, i))
				case nd.Count:
					min, _ := strconv.Atoi(q.Value)
					and = append(and, op.Repeat(min, i))
				default:
					return nil, fmt.Errorf("unknown quant child: %v", nd.NodeTypes[n.Type])
				}
			}
		// PosLook <-- '&' Primary Quant?
		case nd.PosLook:
			i, err := p.generatePrimary(g, n.Children()[0])
			if err != nil {
				return nil, err
			}
			if i != nil {
				and = append(and, op.Ensure{Value: i})
			}
		// NegLook <-- '!' Primary Quant?
		case nd.NegLook:
			i, err := p.generatePrimary(g, n.Children()[0])
			if err != nil {
				return nil, err
			}
			if i != nil {
				and = append(and, op.Not{Value: i})
			}
		default:
			return nil, fmt.Errorf("unknown sequence child: %v", nd.NodeTypes[n.Type])
		}
	}
	if len(and) == 1 {
		return and[0], nil
	}
	return and, nil
}

// generatePrimary is responsible for generation the primary values.
func (p *internalParser) generatePrimary(g generator, n *ast.Node) (interface{}, error) {
	switch n.Type {
	case nd.Comment, nd.EndLine:
		// Ignore these.
		return nil, nil
	case nd.Unicode, nd.Hexadecimal:
		i, _ := convertToInt(n.Value[1:], 16)
		return i, nil
	case nd.Binary:
		i, _ := convertToInt(n.Value[1:], 2)
		return i, nil
	case nd.Octal:
		i, _ := convertToInt(n.Value[1:], 8)
		return i, nil
	case nd.ClassId, nd.ResClassId:
		name := g.className(n.Value)

		p.RWMutex.RLock()
		if class, ok := p.classes[name]; ok {
			p.RWMutex.RUnlock()
			return class, nil
		}
		p.RWMutex.RUnlock()

		var wg sync.WaitGroup
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			for {
				p.RWMutex.RLock()
				v := p.classes[name]
				p.RWMutex.RUnlock()
				if v != nil {
					break
				}
			}
		}(name)
		wg.Wait()

		p.RWMutex.RLock()
		i := p.classes[name]
		p.RWMutex.RUnlock()
		return i, nil
	case nd.TokenId, nd.ResTokenId:
		name := g.tokenName(n.Value)

		p.RWMutex.RLock()
		if class, ok := p.tokens[name]; ok {
			p.RWMutex.RUnlock()
			return class, nil
		}
		p.RWMutex.RUnlock()

		var wg sync.WaitGroup
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			for {
				p.RWMutex.RLock()
				v := p.tokens[name]
				p.RWMutex.RUnlock()
				if v != nil {
					break
				}
			}
		}(name)
		wg.Wait()

		p.RWMutex.RLock()
		i := p.tokens[name]
		p.RWMutex.RUnlock()
		return i, nil
	case nd.CheckId:
		name := g.nodeName(n.Value)
		return ast.LoopUp{
			Key:   name,
			Table: &p.nodes,
		}, nil
	case nd.AlphaRange, nd.IntRange:
		min := rune(n.Children()[0].Value[0])
		max := rune(n.Children()[1].Value[0])
		return parser.CheckRuneRange(min, max), nil
	case nd.UniRange, nd.HexRange:
		min, _ := convertToInt(removeUnicodePrefix(n.Children()[0].Value), 16)
		max, _ := convertToInt(removeUnicodePrefix(n.Children()[1].Value), 16)
		return parser.CheckRuneRange(rune(min), rune(max)), nil
	case nd.BinRange:
		min, _ := convertToInt(removeBinaryPrefix(n.Children()[0].Value), 2)
		max, _ := convertToInt(removeBinaryPrefix(n.Children()[1].Value), 2)
		return parser.CheckRuneRange(rune(min), rune(max)), nil
	case nd.OctRange:
		min, _ := convertToInt(removeOctalPrefix(n.Children()[0].Value), 8)
		max, _ := convertToInt(removeOctalPrefix(n.Children()[1].Value), 8)
		return parser.CheckRuneRange(rune(min), rune(max)), nil
	case nd.String:
		if v := n.Value; len(v) == 1 {
			return parser.CheckRune([]rune(v)[0]), nil
		}
		return parser.CheckString(n.Value), nil
	case nd.Expression:
		return p.generateExpression(g, n.Children())
	default:
		return nil, fmt.Errorf("unknown plain child: %v", nd.NodeTypes[n.Type])
	}
}
