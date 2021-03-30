package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type dotAtom struct {
	value string
}

func (w *walker) EnterDotAtom(ctx *parser.DotAtomContext) {
	w.enter(&dotAtom{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitDotAtom(ctx *parser.DotAtomContext) {
	type withDotAtom interface {
		withDotAtom(*dotAtom)
	}

	res := w.exit().(*dotAtom)

	if parent, ok := w.parent().(withDotAtom); ok {
		parent.withDotAtom(res)
	}
}
