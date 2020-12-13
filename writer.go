package pegen

import (
	"fmt"
	"io"
	"strings"
)

type writer struct {
	io.Writer

	prefix string
}

func (w *writer) write(p []byte) (int, error) {
	return w.Write(append([]byte(w.prefix), p...))
}

func (w *writer) w(p string) {
	_, _ = w.write([]byte(p))
}
func (w *writer) wf(format string, args ...interface{}) {
	w.w(fmt.Sprintf(format, args...))
}

func (w *writer) ln() {
	w.Write([]byte("\n"))
}

func (w *writer) wln(p string) {
	w.w(p + "\n")
}

func (w *writer) wlnf(format string, args ...interface{}) {
	w.wln(fmt.Sprintf(format, args...))
}

func (w *writer) c(comment string) {
	w.wlnf("// %s", comment)
}

func (w *writer) indent() *writer {
	return &writer{
		Writer: w.Writer,
		prefix: w.prefix + "\t",
	}
}

func (w *writer) noIndent() *writer {
	return &writer{
		Writer: w.Writer,
	}
}

func fillRight(v string, size int) string {
	return fmt.Sprintf("%s%s", v, strings.Repeat(" ", size-len(v)))
}
