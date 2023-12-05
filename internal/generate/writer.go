package generate

import (
	"bytes"
	"fmt"
)

type writer struct {
	buff  bytes.Buffer
	ident int
}

func newWriter() *writer {
	return &writer{buff: *bytes.NewBuffer(nil)}
}

func (w *writer) P(format string, a ...any) {
	for idx := 0; idx < w.ident; idx++ {
		w.buff.WriteString("\t")
	}
	if _, err := fmt.Fprintf(&w.buff, format, a...); err != nil {
		panic(err)
	}
	w.buff.WriteString("\n")
}

func (w *writer) Ident() {
	w.ident++
}

func (w *writer) EndIdent() {
	w.ident--
	if w.ident < 0 {
		w.ident = 0
	}
}

func (w *writer) String() string {
	return w.buff.String()
}
