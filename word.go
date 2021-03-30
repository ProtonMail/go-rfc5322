package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type word struct {
	value string
}

func (w *word) withAtom(atom *atom) {
	w.value = atom.value
}

func (w *word) withQuotedString(quotedString *quotedString) {
	w.value = quotedString.value
}

func (w *word) withEncodedWord(encodedWord *encodedWord) {
	w.value += encodedWord.value
}

func (w *walker) EnterWord(ctx *parser.WordContext) {
	w.enter(&word{})
}

func (w *walker) ExitWord(ctx *parser.WordContext) {
	type withWord interface {
		withWord(*word)
	}

	res := w.exit().(*word)

	if parent, ok := w.parent().(withWord); ok {
		parent.withWord(res)
	}
}
