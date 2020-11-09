package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type qtext struct {
	value string
}

func (w *walker) EnterQtext(ctx *parser.QtextContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering qtext")

	w.enter(&qtext{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitQtext(ctx *parser.QtextContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting qtext")

	type withQtext interface {
		withQtext(*qtext)
	}

	res := w.exit().(*qtext)

	if parent, ok := w.parent().(withQtext); ok {
		parent.withQtext(res)
	}
}
