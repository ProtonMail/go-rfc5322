package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type addrSpec struct {
	localPart, domain string
}

func (a *addrSpec) withLocalPart(localPart *localPart) {
	a.localPart = localPart.value
}

func (a *addrSpec) withDomain(domain *domain) {
	a.domain = domain.value
}

func (a *addrSpec) withPort(port *port) {
	a.domain += ":" + port.value
}

func (w *walker) EnterAddrSpec(ctx *parser.AddrSpecContext) {
	w.enter(&addrSpec{})
}

func (w *walker) ExitAddrSpec(ctx *parser.AddrSpecContext) {
	type withAddrSpec interface {
		withAddrSpec(*addrSpec)
	}

	res := w.exit().(*addrSpec)

	if parent, ok := w.parent().(withAddrSpec); ok {
		parent.withAddrSpec(res)
	}
}
