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
		ClassAliases: map[string]string{
			"alphanum": "AlphaNum",
			"unipoint": "UniPoint",
			"bindig":   "BinDig",
			"hexdig":   "HexDig",
			"lowerhex": "LowerHex",
			"octdig":   "OctDig",
			"uphex":    "UpHex",
			"ws":       "Whitespace",
			"ascii":    "ASCII",
		},
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
		w.wln("\"github.com/di-wu/parser\"")
		w.wln("\"github.com/di-wu/parser/op\"")
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
