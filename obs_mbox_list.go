package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type obsMboxList struct {
	addresses []*mail.Address
}

func (ml *obsMboxList) withMailbox(mailbox *mailbox) {
	ml.addresses = append(ml.addresses, &mail.Address{
		Name:    mailbox.name,
		Address: mailbox.address,
	})
}

func (w *walker) EnterObsMboxList(ctx *parser.ObsMboxListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering obsMboxList")
	w.enter(&obsMboxList{})
}

func (w *walker) ExitObsMboxList(ctx *parser.ObsMboxListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting obsMboxList")

	type withObsMboxList interface {
		withObsMboxList(*obsMboxList)
	}

	res := w.exit().(*obsMboxList)

	if parent, ok := w.parent().(withObsMboxList); ok {
		parent.withObsMboxList(res)
	}
}
