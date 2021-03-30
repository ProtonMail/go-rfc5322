package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type fws struct {
	value string
}

func (w *walker) EnterFws(ctx *parser.FwsContext) {
	w.enter(&fws{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitFws(ctx *parser.FwsContext) {
	type withFws interface {
		withFws(*fws)
	}

	res := w.exit().(*fws)

	if parent, ok := w.parent().(withFws); ok {
		parent.withFws(res)
	}
}
