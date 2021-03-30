package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type domainLiteral struct {
	value string
}

func (w *walker) EnterDomainLiteral(ctx *parser.DomainLiteralContext) {
	w.enter(&domainLiteral{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitDomainLiteral(ctx *parser.DomainLiteralContext) {
	type withDomainLiteral interface {
		withDomainLiteral(*domainLiteral)
	}

	res := w.exit().(*domainLiteral)

	if parent, ok := w.parent().(withDomainLiteral); ok {
		parent.withDomainLiteral(res)
	}
}
