package pegen

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	grammar, err := ioutil.ReadFile("./testdata/grammar.pegn")
	if err != nil {
		t.Error(err)
	}

	b := new(bytes.Buffer)
	g, err := New(grammar, b, Config{
		IgnoreReserved: true,
	})
	if err != nil {
		t.Error(err)
	}
	g.Generate()
	_ = ioutil.WriteFile("./testdata/test.go", b.Bytes(), 0777)
}
