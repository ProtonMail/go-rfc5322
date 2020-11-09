package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type nameAddr struct {
	name, address string
}

func (a *nameAddr) withDisplayName(displayName *displayName) {
	a.name = displayName.value
}

func (a *nameAddr) withAngleAddr(angleAddr *angleAddr) {
	a.address = angleAddr.address
}

func (w *walker) EnterNameAddr(ctx *parser.NameAddrContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering nameAddr")
	w.enter(&nameAddr{})
}

func (w *walker) ExitNameAddr(ctx *parser.NameAddrContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting nameAddr")

	type withNameAddr interface {
		withNameAddr(*nameAddr)
	}

	res := w.exit().(*nameAddr)

	if parent, ok := w.parent().(withNameAddr); ok {
		parent.withNameAddr(res)
	}
}
