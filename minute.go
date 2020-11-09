package rfc5322

import (
	"strconv"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type minute struct {
	value int
}

func (w *walker) EnterMinute(ctx *parser.MinuteContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering minute")

	var text string

	for _, digit := range ctx.AllDigit() {
		text += digit.GetText()
	}

	val, err := strconv.Atoi(text)
	if err != nil {
		w.err = err
	}

	w.enter(&minute{
		value: val,
	})
}

func (w *walker) ExitMinute(ctx *parser.MinuteContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting minute")

	type withMinute interface {
		withMinute(*minute)
	}

	res := w.exit().(*minute)

	if parent, ok := w.parent().(withMinute); ok {
		parent.withMinute(res)
	}
}
