package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type atom struct {
	value string
}

func (w *walker) EnterAtom(ctx *parser.AtomContext) {
	w.enter(&atom{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitAtom(ctx *parser.AtomContext) {
	type withAtom interface {
		withAtom(*atom)
	}

	res := w.exit().(*atom)

	if parent, ok := w.parent().(withAtom); ok {
		parent.withAtom(res)
	}
}
