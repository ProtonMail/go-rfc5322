package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type group struct {
	addresses []mail.Address
}

func (g *group) withGroupList(groupList *groupList) {
	g.addresses = append(g.addresses, groupList.addresses...)
}

func (w *walker) EnterGroup(ctx *parser.GroupContext) {
	w.enter(&group{})
}

func (w *walker) ExitGroup(ctx *parser.GroupContext) {
	type withGroup interface {
		withGroup(*group)
	}

	res := w.exit().(*group)

	if parent, ok := w.parent().(withGroup); ok {
		parent.withGroup(res)
	}
}
