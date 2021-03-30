package rfc5322

import (
	"strings"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type domain struct {
	value string
}

func (d *domain) withDotAtom(dotAtom *dotAtom) {
	d.value = dotAtom.value
}

func (d *domain) withDomainLiteral(domainLiteral *domainLiteral) {
	d.value = domainLiteral.value
}

func (d *domain) withObsDomain(obsDomain *obsDomain) {
	d.value = strings.Join(obsDomain.atoms, ".")
}

func (w *walker) EnterDomain(ctx *parser.DomainContext) {
	w.enter(&domain{})
}

func (w *walker) ExitDomain(ctx *parser.DomainContext) {
	type withDomain interface {
		withDomain(*domain)
	}

	res := w.exit().(*domain)

	if parent, ok := w.parent().(withDomain); ok {
		parent.withDomain(res)
	}
}
