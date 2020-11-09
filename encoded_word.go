package rfc5322

import (
	"mime"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type encodedWord struct {
	value string
}

func (w *walker) EnterEncodedWord(ctx *parser.EncodedWordContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering encodedWord")

	dec := &mime.WordDecoder{CharsetReader: CharsetReader}

	word, err := dec.Decode(ctx.GetText())
	if err != nil {
		word = ctx.GetText()
	}

	w.enter(&encodedWord{
		value: word,
	})
}

func (w *walker) ExitEncodedWord(ctx *parser.EncodedWordContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting encodedWord")

	type withEncodedWord interface {
		withEncodedWord(*encodedWord)
	}

	res := w.exit().(*encodedWord)

	if parent, ok := w.parent().(withEncodedWord); ok {
		parent.withEncodedWord(res)
	}
}
