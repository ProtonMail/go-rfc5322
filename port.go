package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type port struct {
	value string
}

func (w *walker) EnterPort(ctx *parser.PortContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering port")

	w.enter(&port{
		value: ctx.GetText(),
	})
}

func (w *walker) ExitPort(ctx *parser.PortContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting port")

	type withPort interface {
		withPort(*port)
	}

	res := w.exit().(*port)

	if parent, ok := w.parent().(withPort); ok {
		parent.withPort(res)
	}
}
