package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
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
	w.enter(&address{})
}

func (w *walker) ExitAddress(ctx *parser.AddressContext) {
	type withAddress interface {
		withAddress(*address)
	}

	res := w.exit().(*address)

	if parent, ok := w.parent().(withAddress); ok {
		parent.withAddress(res)
	}
}
