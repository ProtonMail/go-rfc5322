package rfc5322

import (
	"errors"
	"strings"
	"time"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type obsZone struct {
	location *time.Location
}

func (w *walker) EnterObsZone(ctx *parser.ObsZoneContext) {
	loc := time.UTC

	switch strings.ToLower(ctx.GetText()) {
	case "ut":
		loc = time.FixedZone(ctx.GetText(), 0)
	case "utc":
		loc = time.FixedZone(ctx.GetText(), 0)
	case "gmt":
		loc = time.FixedZone(ctx.GetText(), 0)
	case "est":
		loc = time.FixedZone(ctx.GetText(), -5*60*60)
	case "edt":
		loc = time.FixedZone(ctx.GetText(), -4*60*60)
	case "cst":
		loc = time.FixedZone(ctx.GetText(), -6*60*60)
	case "cdt":
		loc = time.FixedZone(ctx.GetText(), -5*60*60)
	case "mst":
		loc = time.FixedZone(ctx.GetText(), -7*60*60)
	case "mdt":
		loc = time.FixedZone(ctx.GetText(), -6*60*60)
	case "pst":
		loc = time.FixedZone(ctx.GetText(), -8*60*60)
	case "pdt":
		loc = time.FixedZone(ctx.GetText(), -7*60*60)
	default:
		w.err = errors.New("bad timezone")
	}

	w.enter(&obsZone{
		location: loc,
	})
}

func (w *walker) ExitObsZone(ctx *parser.ObsZoneContext) {
	type withObsZone interface {
		withObsZone(*obsZone)
	}

	res := w.exit().(*obsZone)

	if parent, ok := w.parent().(withObsZone); ok {
		parent.withObsZone(res)
	}
}
