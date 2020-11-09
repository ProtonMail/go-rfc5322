package rfc5322

import (
	"fmt"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type mailbox struct {
	name, address string
}

func (m *mailbox) withNameAddr(nameAddr *nameAddr) {
	m.name = nameAddr.name
	m.address = nameAddr.address
}

func (m *mailbox) withAddrSpec(addrSpec *addrSpec) {
	m.address = fmt.Sprintf("%v@%v", addrSpec.localPart, addrSpec.domain)
}

func (w *walker) EnterMailbox(ctx *parser.MailboxContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering mailbox")
	w.enter(&mailbox{})
}

func (w *walker) ExitMailbox(ctx *parser.MailboxContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting mailbox")

	type withMailbox interface {
		withMailbox(*mailbox)
	}

	res := w.exit().(*mailbox)

	if parent, ok := w.parent().(withMailbox); ok {
		parent.withMailbox(res)
	}
}
