package rfc5322

import (
	"time"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type zone struct {
	location *time.Location
}

func (z *zone) withOffset(offset *offset) {
	z.location = time.FixedZone(offset.rep, offset.value)
}

func (z *zone) withObsZone(obsZone *obsZone) {
	z.location = obsZone.location
}

func (w *walker) EnterZone(ctx *parser.ZoneContext) {
	w.enter(&zone{
		location: time.UTC,
	})
}

func (w *walker) ExitZone(ctx *parser.ZoneContext) {
	type withZone interface {
		withZone(*zone)
	}

	res := w.exit().(*zone)

	if parent, ok := w.parent().(withZone); ok {
		parent.withZone(res)
	}
}
