package rfc5322

import (
	"strconv"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type second struct {
	value int
}

func (w *walker) EnterSecond(ctx *parser.SecondContext) {
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
	type withSecond interface {
		withSecond(*second)
	}

	res := w.exit().(*second)

	if parent, ok := w.parent().(withSecond); ok {
		parent.withSecond(res)
	}
}
