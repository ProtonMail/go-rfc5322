package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
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
	w.enter(&nameAddr{})
}

func (w *walker) ExitNameAddr(ctx *parser.NameAddrContext) {
	type withNameAddr interface {
		withNameAddr(*nameAddr)
	}

	res := w.exit().(*nameAddr)

	if parent, ok := w.parent().(withNameAddr); ok {
		parent.withNameAddr(res)
	}
}
