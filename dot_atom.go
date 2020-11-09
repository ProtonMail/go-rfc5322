package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type dotAtom struct {
	value string
}

func (w *walker) EnterDotAtom(ctx *parser.DotAtomContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering dotAtom")

	w.enter(&dotAtom{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitDotAtom(ctx *parser.DotAtomContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting dotAtom")

	type withDotAtom interface {
		withDotAtom(*dotAtom)
	}

	res := w.exit().(*dotAtom)

	if parent, ok := w.parent().(withDotAtom); ok {
		parent.withDotAtom(res)
	}
}
