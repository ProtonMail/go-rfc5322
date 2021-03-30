package rfc5322

import (
	"github.com/ProtonMail/go-rfc5322/parser"
)

type displayName struct {
	value    string
	unspaced bool
}

func (n *displayName) withWord(word *word) {
	if n.unspaced {
		n.unspaced = false
	} else if len(n.value) > 0 {
		n.value += " "
	}

	n.value += word.value
}

func (n *displayName) withUnspaced(unspaced *unspaced) {
	n.unspaced = true
	n.value += unspaced.value
}

func (w *walker) EnterDisplayName(ctx *parser.DisplayNameContext) {
	w.enter(&displayName{})
}

func (w *walker) ExitDisplayName(ctx *parser.DisplayNameContext) {
	type withDisplayName interface {
		withDisplayName(*displayName)
	}

	res := w.exit().(*displayName)

	if parent, ok := w.parent().(withDisplayName); ok {
		parent.withDisplayName(res)
	}
}
