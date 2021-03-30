package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type quotedChar struct {
	value string
}

func (w *walker) EnterQuotedChar(ctx *parser.QuotedCharContext) {
	w.enter(&quotedChar{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitQuotedChar(ctx *parser.QuotedCharContext) {
	type withQuotedChar interface {
		withQuotedChar(*quotedChar)
	}

	res := w.exit().(*quotedChar)

	if parent, ok := w.parent().(withQuotedChar); ok {
		parent.withQuotedChar(res)
	}
}
