package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type groupList struct {
	addresses []*mail.Address
}

func (gl *groupList) withMailboxList(mailboxList *mailboxList) {
	gl.addresses = append(gl.addresses, mailboxList.addresses...)
}

func (w *walker) EnterGroupList(ctx *parser.GroupListContext) {
	w.enter(&groupList{})
}

func (w *walker) ExitGroupList(ctx *parser.GroupListContext) {
	type withGroupList interface {
		withGroupList(*groupList)
	}

	res := w.exit().(*groupList)

	if parent, ok := w.parent().(withGroupList); ok {
		parent.withGroupList(res)
	}
}
