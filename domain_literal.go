package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type domainLiteral struct {
	value string
}

func (w *walker) EnterDomainLiteral(ctx *parser.DomainLiteralContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering domainLiteral")

	w.enter(&domainLiteral{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitDomainLiteral(ctx *parser.DomainLiteralContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting domainLiteral")

	type withDomainLiteral interface {
		withDomainLiteral(*domainLiteral)
	}

	res := w.exit().(*domainLiteral)

	if parent, ok := w.parent().(withDomainLiteral); ok {
		parent.withDomainLiteral(res)
	}
}
