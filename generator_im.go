package pegen

import (
	"fmt"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
	"github.com/pegn/pegn-go/pegn"
	"io/ioutil"
	"net/http"
	"sync"
)

type Parser struct {
	tokens  map[string]interface{}
	classes map[string]interface{}
}

type internalParser struct {
	sync.WaitGroup
	sync.RWMutex
	Parser
}

func ParserFromURLs(config Config, urls ...string) (Parser, error) {
	files := make([][]byte, len(urls))
	for i, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return Parser{}, err
		}
		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Parser{}, err
		}
		files[i] = raw
	}
	return ParserFromFiles(config, files[0], files[1:]...)
}

func ParserFromFiles(config Config, grammar []byte, dependencies ...[]byte) (Parser, error) {
	g, err := NewFromFiles(config, grammar, dependencies...)
	if err != nil {
		return Parser{}, err
	}

	if err := g.prepare(); err != nil {
		return Parser{}, err
	}

	p := internalParser{
		Parser: Parser{
			tokens:  make(map[string]interface{}),
			classes: make(map[string]interface{}),
		},
	}
	p.tokens["TODO"] = '\u0000' // TODO: remove me

	if err := p.generateTokens(g); err != nil {
		return Parser{}, err
	}
	if err := p.generateClasses(g); err != nil {
		return Parser{}, err
	}

	return p.Parser, nil
}

func (p *internalParser) generateTokens(g Generator) error {
	for _, token := range g.tokens {
		if _, ok := p.tokens[token.name]; ok {
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

func (p *internalParser) generateClasses(g Generator) error {
	var wg sync.WaitGroup

	generate := func(g Generator) {
		defer wg.Done()

		for _, class := range g.classes {
			var or op.Or
			for _, n := range class.expression {
				p.WaitGroup.Add(1)
				go func(n *ast.Node) {
					defer p.WaitGroup.Done()

					switch n.Type {
					case pegn.CommentType, pegn.EndLineType:
						// Ignore these.
						return
					case pegn.UnicodeType, pegn.HexadecimalType:
						i, _ := ConvertToInt(n.ValueString()[1:], 16)
						or = append(or, i)
					case pegn.BinaryType:
						i, _ := ConvertToInt(n.ValueString()[1:], 2)
						or = append(or, i)
					case pegn.OctalType:
						i, _ := ConvertToInt(n.ValueString()[1:], 8)
						or = append(or, i)
					case pegn.ClassIdType, pegn.ResClassIdType:
						name := g.className(n.ValueString())

						p.RWMutex.RLock()
						if class, ok := p.classes[name]; ok {
							or = append(or, class)
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
						or = append(or, p.classes[name])
					case pegn.TokenIdType, pegn.ResTokenIdType:
						name := g.tokenName(n.ValueString())

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
						or = append(or, p.tokens[name])
					case pegn.AlphaRangeType, pegn.IntRangeType:
						min := rune(n.Children()[0].ValueString()[0])
						max := rune(n.Children()[1].ValueString()[0])
						or = append(or, parser.CheckRuneRange(min, max))
					case pegn.UniRangeType, pegn.HexRangeType:
						min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 16)
						max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 16)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case pegn.BinRangeType:
						min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 2)
						max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 2)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case pegn.OctRangeType:
						min, _ := ConvertToInt(n.Children()[0].ValueString()[1:], 8)
						max, _ := ConvertToInt(n.Children()[1].ValueString()[1:], 8)
						or = append(or, parser.CheckRuneRange(rune(min), rune(max)))
					case pegn.StringType:
						or = append(or, parser.CheckString(n.ValueString()))
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
	wg.Wait()

	return nil
}
