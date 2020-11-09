package rfc5322

import (
	"strconv"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type hour struct {
	value int
}

func (w *walker) EnterHour(ctx *parser.HourContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering hour")

	var text string

	for _, digit := range ctx.AllDigit() {
		text += digit.GetText()
	}

	val, err := strconv.Atoi(text)
	if err != nil {
		w.err = err
	}

	w.enter(&hour{
		value: val,
	})
}

func (w *walker) ExitHour(ctx *parser.HourContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting hour")

	type withHour interface {
		withHour(*hour)
	}

	res := w.exit().(*hour)

	if parent, ok := w.parent().(withHour); ok {
		parent.withHour(res)
	}
}
