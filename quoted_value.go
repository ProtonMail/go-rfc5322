package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
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
	logrus.WithField("text", ctx.GetText()).Trace("Entering quotedValue")
	w.enter(&quotedValue{})
}

func (w *walker) ExitQuotedValue(ctx *parser.QuotedValueContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting quotedValue")

	type withQuotedValue interface {
		withQuotedValue(*quotedValue)
	}

	res := w.exit().(*quotedValue)

	if parent, ok := w.parent().(withQuotedValue); ok {
		parent.withQuotedValue(res)
	}
}
