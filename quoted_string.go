package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type quotedString struct {
	value string
}

func (s *quotedString) withQuotedValue(quotedValue *quotedValue) {
	s.value = quotedValue.value
}

func (w *walker) EnterQuotedString(ctx *parser.QuotedStringContext) {
	w.enter(&quotedString{})
}

func (w *walker) ExitQuotedString(ctx *parser.QuotedStringContext) {
	type withQuotedString interface {
		withQuotedString(*quotedString)
	}

	res := w.exit().(*quotedString)

	if parent, ok := w.parent().(withQuotedString); ok {
		parent.withQuotedString(res)
	}
}
