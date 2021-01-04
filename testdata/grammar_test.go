package pegn

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
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

	if _, err := p.Expect(op.And{
		Grammar,
		parser.EOD,
	}); err != nil {
		t.Error(err)
	}
}
