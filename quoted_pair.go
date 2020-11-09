package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type quotedPair struct {
	value string
}

func (p *quotedPair) withQuotedChar(quotedChar *quotedChar) {
	p.value = quotedChar.value
}

func (w *walker) EnterQuotedPair(ctx *parser.QuotedPairContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering quotedPair")
	w.enter(&quotedPair{})
}

func (w *walker) ExitQuotedPair(ctx *parser.QuotedPairContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting quotedPair")

	type withQuotedPair interface {
		withQuotedPair(*quotedPair)
	}

	res := w.exit().(*quotedPair)

	if parent, ok := w.parent().(withQuotedPair); ok {
		parent.withQuotedPair(res)
	}
}
