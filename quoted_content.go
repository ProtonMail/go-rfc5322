package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type quotedContent struct {
	value string
}

func (c *quotedContent) withQtext(qtext *qtext) {
	c.value = qtext.value
}

func (c *quotedContent) withQuotedPair(quotedPair *quotedPair) {
	c.value = quotedPair.value
}

func (w *walker) EnterQuotedContent(ctx *parser.QuotedContentContext) {
	w.enter(&quotedContent{})
}

func (w *walker) ExitQuotedContent(ctx *parser.QuotedContentContext) {
	type withQuotedContent interface {
		withQuotedContent(*quotedContent)
	}

	res := w.exit().(*quotedContent)

	if parent, ok := w.parent().(withQuotedContent); ok {
		parent.withQuotedContent(res)
	}
}
