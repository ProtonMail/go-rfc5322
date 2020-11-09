package rfc5322

import (
	"strconv"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type day struct {
	value int
}

func (w *walker) EnterDay(ctx *parser.DayContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering day")

	var text string

	for _, digit := range ctx.AllDigit() {
		text += digit.GetText()
	}

	val, err := strconv.Atoi(text)
	if err != nil {
		w.err = err
	}

	w.enter(&day{
		value: val,
	})
}

func (w *walker) ExitDay(ctx *parser.DayContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting day")

	type withDay interface {
		withDay(*day)
	}

	res := w.exit().(*day)

	if parent, ok := w.parent().(withDay); ok {
		parent.withDay(res)
	}
}
