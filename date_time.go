package rfc5322

import (
	"time"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type dateTime struct {
	day   int
	month time.Month
	year  int

	hour, min, sec int

	loc *time.Location
}

func (dt *dateTime) withDay(day *day) {
	dt.day = day.value
}

func (dt *dateTime) withMonth(month *month) {
	dt.month = month.value
}

func (dt *dateTime) withYear(year *year) {
	dt.year = year.value
}

func (dt *dateTime) withHour(hour *hour) {
	dt.hour = hour.value
}

func (dt *dateTime) withMinute(minute *minute) {
	dt.min = minute.value
}

func (dt *dateTime) withSecond(second *second) {
	dt.sec = second.value
}

func (dt *dateTime) withZone(zone *zone) {
	dt.loc = zone.location
}

func (w *walker) EnterDateTime(ctx *parser.DateTimeContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering dateTime")

	w.enter(&dateTime{
		loc: time.UTC,
	})
}

func (w *walker) ExitDateTime(ctx *parser.DateTimeContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting dateTime")

	dt := w.exit().(*dateTime)

	w.res = time.Date(dt.year, dt.month, dt.day, dt.hour, dt.min, dt.sec, 0, dt.loc)
}
