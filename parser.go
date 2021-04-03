package rfc5322

import (
	"io"
	"net/mail"
	"time"

	"github.com/ProtonMail/go-rfc5322/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CharsetReader, if non-nil, defines a function to generate charset-conversion readers,
// converting from the provided charset into UTF-8, for use when decoding RFC2047 encoded words.
// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets are handled by default.
// One of the CharsetReader's result values must be non-nil.
var CharsetReader func(charset string, input io.Reader) (io.Reader, error)

// ParseAddressList parses one or more valid RFC5322 (with RFC2047) addresses.
func ParseAddressList(input string) ([]mail.Address, error) {
	if len(input) == 0 {
		return []mail.Address{}, nil
	}

	l := parser.NewRFC5322Lexer(antlr.NewInputStream(input))
	p := parser.NewRFC5322Parser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))
	w := &walker{}

	p.RemoveErrorListeners()
	p.AddErrorListener(w)

	tree := p.AddressList()
	if w.err != nil {
		return nil, w.err
	}

	antlr.ParseTreeWalkerDefault.Walk(w, tree)

	return w.res.([]mail.Address), w.err
}

// ParseDateTime parses a valid RFC5322 date-time.
func ParseDateTime(input string) (time.Time, error) {
	if len(input) == 0 {
		return time.Time{}, nil
	}

	l := parser.NewRFC5322Lexer(antlr.NewInputStream(input))
	p := parser.NewRFC5322Parser(antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel))
	w := &walker{}

	p.RemoveErrorListeners()
	p.AddErrorListener(w)

	tree := p.DateTime()
	if w.err != nil {
		return time.Time{}, w.err
	}

	antlr.ParseTreeWalkerDefault.Walk(w, tree)

	return w.res.(time.Time), w.err
}
