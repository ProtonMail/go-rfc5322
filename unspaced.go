package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type unspaced struct {
	value string
}

func (w *walker) EnterUnspaced(ctx *parser.UnspacedContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering unspaced")

	w.enter(&unspaced{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitUnspaced(ctx *parser.UnspacedContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting unspaced")

	type withUnspaced interface {
		withUnspaced(*unspaced)
	}

	res := w.exit().(*unspaced)

	if parent, ok := w.parent().(withUnspaced); ok {
		parent.withUnspaced(res)
	}
}
