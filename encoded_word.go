package rfc5322

import (
	"io"
	"mime"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

// CharsetReader, if non-nil, defines a function to generate
// charset-conversion readers, converting from the provided
// charset into UTF-8.
// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets
// are handled by default.
// One of the CharsetReader's result values must be non-nil.
var CharsetReader func(charset string, input io.Reader) (io.Reader, error)

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
