package rfc5322

import (
	"fmt"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

// When interpreting addresses, the route portion SHOULD be ignored.

type obsAngleAddr struct {
	address string
}

func (a *obsAngleAddr) withAddrSpec(addrSpec *addrSpec) {
	a.address = fmt.Sprintf("%v@%v", addrSpec.localPart, addrSpec.domain)
}

func (w *walker) EnterObsAngleAddr(ctx *parser.ObsAngleAddrContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering obsAngleAddr")
	w.enter(&obsAngleAddr{})
}

func (w *walker) ExitObsAngleAddr(ctx *parser.ObsAngleAddrContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting obsAngleAddr")

	type withObsAngleAddr interface {
		withObsAngleAddr(*obsAngleAddr)
	}

	res := w.exit().(*obsAngleAddr)

	if parent, ok := w.parent().(withObsAngleAddr); ok {
		parent.withObsAngleAddr(res)
	}
}
