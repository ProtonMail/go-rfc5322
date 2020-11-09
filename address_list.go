package rfc5322

import (
	"net/mail"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/sirupsen/logrus"
)

type addressList struct {
	addresses []*mail.Address
}

func (a *addressList) withAddress(address *address) {
	a.addresses = append(a.addresses, address.addresses...)
}

func (w *walker) EnterAddressList(ctx *parser.AddressListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Entering addressList")
	w.enter(&addressList{})
}

func (w *walker) ExitAddressList(ctx *parser.AddressListContext) {
	logrus.WithField("text", ctx.GetText()).Trace("Exiting addressList")
	w.res = w.exit().(*addressList).addresses
}
