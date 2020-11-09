package rfc5322

import (
	"strconv"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type second struct {
	value int
}

func (w *walker) EnterSecond(ctx *parser.SecondContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering second")

	var text string

	for _, digit := range ctx.AllDigit() {
		text += digit.GetText()
	}

	val, err := strconv.Atoi(text)
	if err != nil {
		w.err = err
	}

	w.enter(&second{
		value: val,
	})
}

func (w *walker) ExitSecond(ctx *parser.SecondContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting second")

	type withSecond interface {
		withSecond(*second)
	}

	res := w.exit().(*second)

	if parent, ok := w.parent().(withSecond); ok {
		parent.withSecond(res)
	}
}
