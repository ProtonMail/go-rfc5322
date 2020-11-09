package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
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
	logrus.WithField("text", ctx.GetText()).Trace("Entering word")
	w.enter(&word{})
}

func (w *walker) ExitWord(ctx *parser.WordContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting word")

	type withWord interface {
		withWord(*word)
	}

	res := w.exit().(*word)

	if parent, ok := w.parent().(withWord); ok {
		parent.withWord(res)
	}
}
