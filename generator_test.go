package pegen

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test(t *testing.T) {
	grammar, err := ioutil.ReadFile("./testdata/grammar.pegn")
	if err != nil {
		t.Error(err)
		return
	}

	g, err := New(grammar, "testdata", Config{
		IgnoreReserved: true,
		TypeSuffix:     "Type",
	})
	if err != nil {
		t.Error(err)
		return
	}
	g.Generate()
	for _, err := range g.errors {
		t.Error(err)
	}

	w, b := newBW()
	g.generateHeader(w)
	w.wln("import (")
	{
		w := w.indent()
		w.wln("\"fmt\"")
		w.wln("\"gitlab.com/pegn/pegn-go\"")
	}
	w.wln(")")
	w.ln()
	w.w(g.writers["ast"].String())
	w.ln()
	w.w(g.writers["is"].String())
	w.ln()
	w.w(g.writers["tk"].String())
	w.ln()
	w.w(g.writers["nd"].String())
	_ = ioutil.WriteFile("./testdata/grammar.go", b.Bytes(), os.ModePerm)
}
