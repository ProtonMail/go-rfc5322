package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
)

type addressList struct {
	addresses []mail.Address
}

func (a *addressList) withAddress(address *address) {
	a.addresses = append(a.addresses, address.addresses...)
}

func (w *walker) EnterAddressList(ctx *parser.AddressListContext) {
	w.enter(&addressList{})
}

func (w *walker) ExitAddressList(ctx *parser.AddressListContext) {
	w.res = w.exit().(*addressList).addresses
}
