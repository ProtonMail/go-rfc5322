package rfc5322

import (
	"fmt"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type angleAddr struct {
	address string
}

func (a *angleAddr) withAddrSpec(addrSpec *addrSpec) {
	a.address = fmt.Sprintf("%v@%v", addrSpec.localPart, addrSpec.domain)
}

func (a *angleAddr) withObsAngleAddr(obsAngleAddr *obsAngleAddr) {
	a.address = obsAngleAddr.address
}

func (w *walker) EnterAngleAddr(ctx *parser.AngleAddrContext) {
	w.enter(&angleAddr{})
}

func (w *walker) ExitAngleAddr(ctx *parser.AngleAddrContext) {
	type withAngleAddr interface {
		withAngleAddr(*angleAddr)
	}

	res := w.exit().(*angleAddr)

	if parent, ok := w.parent().(withAngleAddr); ok {
		parent.withAngleAddr(res)
	}
}
