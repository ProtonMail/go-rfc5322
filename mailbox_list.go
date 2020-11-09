package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type mailboxList struct {
	addresses []*mail.Address
}

func (ml *mailboxList) withMailbox(mailbox *mailbox) {
	ml.addresses = append(ml.addresses, &mail.Address{
		Name:    mailbox.name,
		Address: mailbox.address,
	})
}

func (ml *mailboxList) withObsMboxList(obsMboxList *obsMboxList) {
	ml.addresses = append(ml.addresses, obsMboxList.addresses...)
}

func (w *walker) EnterMailboxList(ctx *parser.MailboxListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering mailboxList")
	w.enter(&mailboxList{})
}

func (w *walker) ExitMailboxList(ctx *parser.MailboxListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting mailboxList")

	type withMailboxList interface {
		withMailboxList(*mailboxList)
	}

	res := w.exit().(*mailboxList)

	if parent, ok := w.parent().(withMailboxList); ok {
		parent.withMailboxList(res)
	}
}
