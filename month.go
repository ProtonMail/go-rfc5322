package rfc5322

import (
	"errors"
	"strings"
	"time"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type month struct {
	value time.Month
}

func (w *walker) EnterMonth(ctx *parser.MonthContext) {
	var m time.Month

	switch strings.ToLower(ctx.GetText()) {
	case "jan":
		m = time.January
	case "feb":
		m = time.February
	case "mar":
		m = time.March
	case "apr":
		m = time.April
	case "may":
		m = time.May
	case "jun":
		m = time.June
	case "jul":
		m = time.July
	case "aug":
		m = time.August
	case "sep":
		m = time.September
	case "oct":
		m = time.October
	case "nov":
		m = time.November
	case "dec":
		m = time.December
	default:
		w.err = errors.New("no such month")
	}

	w.enter(&month{
		value: m,
	})
}

func (w *walker) ExitMonth(ctx *parser.MonthContext) {
	type withMonth interface {
		withMonth(*month)
	}

	res := w.exit().(*month)

	if parent, ok := w.parent().(withMonth); ok {
		parent.withMonth(res)
	}
}
