package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type quotedValue struct {
	value string
}

func (v *quotedValue) withFws(fws *fws) {
	v.value += fws.value
}

func (v *quotedValue) withQuotedContent(quotedContent *quotedContent) {
	v.value += quotedContent.value
}

func (w *walker) EnterQuotedValue(ctx *parser.QuotedValueContext) {
	w.enter(&quotedValue{})
}

func (w *walker) ExitQuotedValue(ctx *parser.QuotedValueContext) {
	type withQuotedValue interface {
		withQuotedValue(*quotedValue)
	}

	res := w.exit().(*quotedValue)

	if parent, ok := w.parent().(withQuotedValue); ok {
		parent.withQuotedValue(res)
	}
}
