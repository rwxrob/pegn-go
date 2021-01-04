package pegn

import (
	"github.com/di-wu/parser/ast"
	"io/ioutil"
	"testing"
)

func TestGrammar(t *testing.T) {
	raw, err := ioutil.ReadFile("./grammar.pegn")
	if err != nil {
		t.Error()
		return
	}

	p, err := ast.New(raw)
	if err != nil {
		t.Error(err)
		return
	}

	if 	_, err := p.Expect(Meta); err != nil {
		t.Error(err)
	}
	if 	_, err := p.Expect(Copyright); err != nil {
		t.Error(err)
	}
	if 	_, err := p.Expect(Licensed); err != nil {
		t.Error(err)
	}
}
