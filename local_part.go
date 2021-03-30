package rfc5322

import (
	"strings"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type localPart struct {
	value string
}

func (p *localPart) withDotAtom(dotAtom *dotAtom) {
	p.value = dotAtom.value
}

func (p *localPart) withQuotedString(quotedString *quotedString) {
	p.value = quotedString.value
}

func (p *localPart) withObsLocalPart(obsLocalPart *obsLocalPart) {
	p.value = strings.Join(obsLocalPart.words, ".")
}

func (w *walker) EnterLocalPart(ctx *parser.LocalPartContext) {
	w.enter(&localPart{})
}

func (w *walker) ExitLocalPart(ctx *parser.LocalPartContext) {
	type withLocalPart interface {
		withLocalPart(*localPart)
	}

	res := w.exit().(*localPart)

	if parent, ok := w.parent().(withLocalPart); ok {
		parent.withLocalPart(res)
	}
}
