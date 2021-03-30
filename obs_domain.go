package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type obsDomain struct {
	atoms []string
}

func (p *obsDomain) withAtom(atom *atom) {
	p.atoms = append(p.atoms, atom.value)
}

func (w *walker) EnterObsDomain(ctx *parser.ObsDomainContext) {
	w.enter(&obsDomain{})
}

func (w *walker) ExitObsDomain(ctx *parser.ObsDomainContext) {
	type withObsDomain interface {
		withObsDomain(*obsDomain)
	}

	res := w.exit().(*obsDomain)

	if parent, ok := w.parent().(withObsDomain); ok {
		parent.withObsDomain(res)
	}
}
