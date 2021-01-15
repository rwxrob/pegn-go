package pegen

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

type InMemoryParser struct {
	Tokens  map[string]interface{}
	Classes map[string]interface{}
	Nodes   map[string]interface{}
}

type internalParser struct {
	sync.WaitGroup
	sync.RWMutex
	InMemoryParser
}

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

func ParserFromFiles(config Config, grammar []byte, dependencies ...[]byte) (InMemoryParser, error) {
	g, err := NewFromFiles(config, grammar, dependencies...)
	if err != nil {
		return InMemoryParser{}, err
	}

	if err := g.prepare(); err != nil {
		return InMemoryParser{}, err
	}

	p := internalParser{
		InMemoryParser: InMemoryParser{
			Tokens:  make(map[string]interface{}),
			Classes: make(map[string]interface{}),
			Nodes:   make(map[string]interface{}),
		},
	}
	p.Tokens["TODO"] = '\u0000' // TODO: remove me

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

func (p *internalParser) generateTokens(g Generator) error {
	for _, token := range g.tokens {
		if _, ok := p.Tokens[token.name]; ok {
			return fmt.Errorf("duplicate token: %s", token.name)
		}

		var and op.And
		for _, tk := range token.rawValues {
			if tk.isString() {
				and = append(and, tk.value)
			} else {
				i, _ := ConvertToInt(tk.hexValue, 16)
				and = append(and, i)
			}
		}

		p.RWMutex.Lock()
		if len(and) == 1 {
			p.Tokens[token.name] = and[0]
		} else {
			p.Tokens[token.name] = and
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

func (p *internalParser) generateClasses(g Generator) error {
	var wg sync.WaitGroup
	errChannel := make(chan error, 1)
	generate := func(g Generator) {
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
						i, _ := ConvertToInt(n.ValueString()[1:], 16)
						or = append(or, i)
					case nd.Binary:
						i, _ := ConvertToInt(n.ValueString()[1:], 2)
						or = append(or, i)
					case nd.Octal:
						i, _ := ConvertToInt(n.ValueString()[1:], 8)
						or = append(or, i)
					case nd.ClassId, nd.ResClassId:
						name := g.className(n.ValueString())

						p.RWMutex.RLock()
						if class, ok := p.Classes[name]; ok {
							or = append(or, class)
						}
						p.RWMutex.RUnlock()

						var wg sync.WaitGroup
						wg.Add(1)
						go func(name string) {
							defer wg.Done()
							for {
								p.RWMutex.RLock()
								v := p.Classes[name]
								p.RWMutex.RUnlock()
								if v != nil {
									break
								}
							}
						}(name)
						wg.Wait()

						p.RWMutex.RLock()
						or = append(or, p.Classes[name])
						p.RWMutex.RUnlock()
					case nd.TokenId, nd.ResTokenId:
						name := g.tokenName(n.ValueString())

						p.RWMutex.RLock()
						if token, ok := p.Tokens[name]; ok {
							or = append(or, token)
						}
						p.RWMutex.RUnlock()

						var wg sync.WaitGroup
						wg.Add(1)
						go func(name string) {
							defer wg.Done()
							for {
								p.RWMutex.RLock()
								v := p.Tokens[name]
								p.RWMutex.RUnlock()
								if v != nil {
									break
								}
							}
						}(name)
						wg.Wait()

						p.RWMutex.RLock()
						or = append(or, p.Tokens[name])
						p.RWMutex.RUnlock()
					case nd.AlphaRange, nd.IntRange:
						min := rune(n.Children()[0].ValueString()[0])
						max := rune(n.Children()[1].ValueString()[0])
						or = append(or, parser.CheckRuneRange(min, max))
					case nd.UniRange, nd.HexRange:
						min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 16)
						max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 16)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case nd.BinRange:
						min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 2)
						max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 2)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case nd.OctRange:
						min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 8)
						max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 8)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case nd.String:
						or = append(or, parser.CheckString(n.ValueString()))
					default:
						errChannel <- fmt.Errorf("unknown class child: %v", nd.NodeTypes[n.Type])
					}
				}(n)
			}
			p.Wait()

			p.RWMutex.Lock()
			if len(or) == 1 {
				p.Classes[class.name] = or[0]
			} else {
				p.Classes[class.name] = or
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

func (p *internalParser) generateNodes(g Generator) error {
	var wg sync.WaitGroup
	errChannel := make(chan error, 1)
	generate := func(g Generator, node node) {
		defer wg.Done()

		i, err := p.generateExpression(g, node.expression)
		if err != nil {
			errChannel <- err
		}
		p.RWMutex.Lock()
		p.Nodes[g.nodeName(node.name)] = i
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

func (p *internalParser) generateExpression(g Generator, expression []*ast.Node) (interface{}, error) {
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

func (p *internalParser) generateSequence(g Generator, sequence []*ast.Node) (interface{}, error) {
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
					min, _ := strconv.Atoi(q.Children()[0].ValueString())
					max, _ := strconv.Atoi(q.Children()[1].ValueString())
					and = append(and, op.MinMax(min, max, i))
				case nd.Count:
					min, _ := strconv.Atoi(q.ValueString())
					and = append(and, op.Repeat(min, i))
				default:
					return nil, fmt.Errorf("unknown quant child: %v", nd.NodeTypes[n.Type])
				}
			}
		// PosLook <-- '&' Primary Quant?
		case nd.PosLook:
			return nil, fmt.Errorf("unsupported: %s", nd.NodeTypes[n.Type])
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

func (p *internalParser) generatePrimary(g Generator, n *ast.Node) (interface{}, error) {
	switch n.Type {
	case nd.Comment, nd.EndLine:
		// Ignore these.
		return nil, nil
	case nd.Unicode, nd.Hexadecimal:
		i, _ := ConvertToInt(n.ValueString()[1:], 16)
		return i, nil
	case nd.Binary:
		i, _ := ConvertToInt(n.ValueString()[1:], 2)
		return i, nil
	case nd.Octal:
		i, _ := ConvertToInt(n.ValueString()[1:], 8)
		return i, nil
	case nd.ClassId, nd.ResClassId:
		name := g.className(n.ValueString())

		p.RWMutex.RLock()
		if class, ok := p.Classes[name]; ok {
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
				v := p.Classes[name]
				p.RWMutex.RUnlock()
				if v != nil {
					break
				}
			}
		}(name)
		wg.Wait()

		p.RWMutex.RLock()
		i := p.Classes[name]
		p.RWMutex.RUnlock()
		return i, nil
	case nd.TokenId, nd.ResTokenId:
		name := g.tokenName(n.ValueString())

		p.RWMutex.RLock()
		if class, ok := p.Tokens[name]; ok {
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
				v := p.Tokens[name]
				p.RWMutex.RUnlock()
				if v != nil {
					break
				}
			}
		}(name)
		wg.Wait()

		p.RWMutex.RLock()
		i := p.Tokens[name]
		p.RWMutex.RUnlock()
		return i, nil
	case nd.CheckId:
		name := g.nodeName(n.ValueString())
		return ast.LoopUp{
			Key:   name,
			Table: &p.Nodes,
		}, nil
	case nd.AlphaRange, nd.IntRange:
		min := rune(n.Children()[0].ValueString()[0])
		max := rune(n.Children()[1].ValueString()[0])
		return parser.CheckRuneRange(min, max), nil
	case nd.UniRange, nd.HexRange:
		min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 16)
		max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 16)
		return parser.CheckRuneRange(rune(min), rune(max)), nil
	case nd.BinRange:
		min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 2)
		max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 2)
		return parser.CheckRuneRange(rune(min), rune(max)), nil
	case nd.OctRange:
		min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 8)
		max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 8)
		return parser.CheckRuneRange(rune(min), rune(max)), nil
	case nd.String:
		return parser.CheckString(n.ValueString()), nil
	case nd.Expression:
		return p.generateExpression(g, n.Children())
	default:
		return nil, fmt.Errorf("unknown plain child: %v", nd.NodeTypes[n.Type])
	}
}
