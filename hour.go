package rfc5322

import (
	"strconv"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type hour struct {
	value int
}

func (w *walker) EnterHour(ctx *parser.HourContext) {
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
	type withHour interface {
		withHour(*hour)
	}

	res := w.exit().(*hour)

	if parent, ok := w.parent().(withHour); ok {
		parent.withHour(res)
	}
}
