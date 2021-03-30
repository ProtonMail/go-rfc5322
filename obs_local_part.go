package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type obsLocalPart struct {
	words []string
}

func (p *obsLocalPart) withWord(word *word) {
	p.words = append(p.words, word.value)
}

func (w *walker) EnterObsLocalPart(ctx *parser.ObsLocalPartContext) {
	w.enter(&obsLocalPart{})
}

func (w *walker) ExitObsLocalPart(ctx *parser.ObsLocalPartContext) {
	type withObsLocalPart interface {
		withObsLocalPart(*obsLocalPart)
	}

	res := w.exit().(*obsLocalPart)

	if parent, ok := w.parent().(withObsLocalPart); ok {
		parent.withObsLocalPart(res)
	}
}
