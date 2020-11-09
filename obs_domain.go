package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type obsDomain struct {
	atoms []string
}

func (p *obsDomain) withAtom(atom *atom) {
	p.atoms = append(p.atoms, atom.value)
}

func (w *walker) EnterObsDomain(ctx *parser.ObsDomainContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering obsDomain")
	w.enter(&obsDomain{})
}

func (w *walker) ExitObsDomain(ctx *parser.ObsDomainContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting obsDomain")

	type withObsDomain interface {
		withObsDomain(*obsDomain)
	}

	res := w.exit().(*obsDomain)

	if parent, ok := w.parent().(withObsDomain); ok {
		parent.withObsDomain(res)
	}
}
