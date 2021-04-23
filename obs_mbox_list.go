package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
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
	w.enter(&obsMboxList{})
}

func (w *walker) ExitObsMboxList(ctx *parser.ObsMboxListContext) {
	type withObsMboxList interface {
		withObsMboxList(*obsMboxList)
	}

	res := w.exit().(*obsMboxList)

	if parent, ok := w.parent().(withObsMboxList); ok {
		parent.withObsMboxList(res)
	}
}
