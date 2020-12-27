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
	w.wlnf("import %q", "gitlab.com/pegn/pegn-go")
	w.ln()
	w.w(g.writers["ast"].String())
	w.ln()
	w.w(g.writers["is"].String())
	w.ln()
	w.w(g.writers["tk"].String())
	w.ln()
	w.w(g.writers["nd"].String())
	_ = ioutil.WriteFile("./testdata/test.go", b.Bytes(), os.ModePerm)
}
