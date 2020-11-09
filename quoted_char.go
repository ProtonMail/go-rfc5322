package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type quotedChar struct {
	value string
}

func (w *walker) EnterQuotedChar(ctx *parser.QuotedCharContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering quotedChar")

	w.enter(&quotedChar{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitQuotedChar(ctx *parser.QuotedCharContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting quotedChar")

	type withQuotedChar interface {
		withQuotedChar(*quotedChar)
	}

	res := w.exit().(*quotedChar)

	if parent, ok := w.parent().(withQuotedChar); ok {
		parent.withQuotedChar(res)
	}
}
