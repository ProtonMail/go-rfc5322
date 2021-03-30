package rfc5322

import (
	"strconv"
	"time"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type year struct {
	value int
}

func (w *walker) EnterYear(ctx *parser.YearContext) {
	var text string

	for _, digit := range ctx.AllDigit() {
		text += digit.GetText()
	}

	val, err := strconv.Atoi(text)
	if err != nil {
		w.err = err
	}

	// NOTE: 2-digit years are obsolete but let's just have some simple handling anyway.
	if len(text) == 2 {
		if val > time.Now().Year()%100 {
			val += 1900
		} else {
			val += 2000
		}
	}

	w.enter(&year{
		value: val,
	})
}

func (w *walker) ExitYear(ctx *parser.YearContext) {
	type withYear interface {
		withYear(*year)
	}

	res := w.exit().(*year)

	if parent, ok := w.parent().(withYear); ok {
		parent.withYear(res)
	}
}
