package rfc5322

import (
	"fmt"
	"strings"
	"time"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type offset struct {
	rep   string
	value int
}

func (w *walker) EnterOffset(ctx *parser.OffsetContext) {
	text := ctx.GetText()

	// NOTE: RFC5322 date-time should always begin with + or -
	// but we relax that requirement a bit due to many messages
	// in the wild that skip the +; we add the "+" if missing.
	if !strings.HasPrefix(text, "+") && !strings.HasPrefix(text, "-") {
		text = "+" + text
	}

	sgn := text[0:1]
	hrs := text[1:3]
	min := text[3:5]

	dur, err := time.ParseDuration(fmt.Sprintf("%v%vh%vm", sgn, hrs, min))
	if err != nil {
		w.err = err
	}

	w.enter(&offset{
		rep:   text,
		value: int(dur.Seconds()),
	})
}

func (w *walker) ExitOffset(ctx *parser.OffsetContext) {
	type withOffset interface {
		withOffset(*offset)
	}

	res := w.exit().(*offset)

	if parent, ok := w.parent().(withOffset); ok {
		parent.withOffset(res)
	}
}
