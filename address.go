package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type address struct {
	addresses []*mail.Address
}

func (a *address) withMailbox(mailbox *mailbox) {
	a.addresses = append(a.addresses, &mail.Address{
		Name:    mailbox.name,
		Address: mailbox.address,
	})
}

func (a *address) withGroup(group *group) {
	a.addresses = append(a.addresses, group.addresses...)
}

func (w *walker) EnterAddress(ctx *parser.AddressContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering address")
	w.enter(&address{})
}

func (w *walker) ExitAddress(ctx *parser.AddressContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting address")

	type withAddress interface {
		withAddress(*address)
	}

	res := w.exit().(*address)

	if parent, ok := w.parent().(withAddress); ok {
		parent.withAddress(res)
	}
}
