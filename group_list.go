package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type groupList struct {
	addresses []*mail.Address
}

func (gl *groupList) withMailboxList(mailboxList *mailboxList) {
	gl.addresses = append(gl.addresses, mailboxList.addresses...)
}

func (w *walker) EnterGroupList(ctx *parser.GroupListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering groupList")
	w.enter(&groupList{})
}

func (w *walker) ExitGroupList(ctx *parser.GroupListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting groupList")

	type withGroupList interface {
		withGroupList(*groupList)
	}

	res := w.exit().(*groupList)

	if parent, ok := w.parent().(withGroupList); ok {
		parent.withGroupList(res)
	}
}
