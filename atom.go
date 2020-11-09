package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type atom struct {
	value string
}

func (w *walker) EnterAtom(ctx *parser.AtomContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering atom")

	w.enter(&atom{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitAtom(ctx *parser.AtomContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting atom")

	type withAtom interface {
		withAtom(*atom)
	}

	res := w.exit().(*atom)

	if parent, ok := w.parent().(withAtom); ok {
		parent.withAtom(res)
	}
}
