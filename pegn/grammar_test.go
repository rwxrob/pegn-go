package pegn_test

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
	"github.com/pegn/pegn-go/pegn"
	"io/ioutil"
	"testing"
)

func TestGrammar(t *testing.T) {
	raw, err := ioutil.ReadFile("./testdata/grammar.pegn")
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
		pegn.Spec,
		parser.EOD,
	}); err != nil {
		t.Error(err)
	}
}
