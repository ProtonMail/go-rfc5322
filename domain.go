package rfc5322

import (
	"strings"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
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
	logrus.WithField("text", ctx.GetText()).Trace("Entering domain")
	w.enter(&domain{})
}

func (w *walker) ExitDomain(ctx *parser.DomainContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting domain")

	type withDomain interface {
		withDomain(*domain)
	}

	res := w.exit().(*domain)

	if parent, ok := w.parent().(withDomain); ok {
		parent.withDomain(res)
	}
}
