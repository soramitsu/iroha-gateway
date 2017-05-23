// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "Iroha-Gateway Server": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/soramitsu/iroha-gateway/design
// --out=$(GOPATH)/src/github.com/soramitsu/iroha-gateway
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// Accounts media type (default view)
//
// Identifier: application/vnd.accounts+json; view=default
type Accounts struct {
	// accounts
	Accounts []*Account `form:"accounts" json:"accounts" xml:"accounts"`
	// response code
	Code int `form:"code" json:"code" xml:"code"`
	// response message
	Message string `form:"message" json:"message" xml:"message"`
}

// Validate validates the Accounts media type instance.
func (mt *Accounts) Validate() (err error) {
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}

	if mt.Accounts == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "accounts"))
	}
	for _, e := range mt.Accounts {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Basic response (default view)
//
// Identifier: application/vnd.message+json; view=default
type Message struct {
	// response code
	Code int `form:"code" json:"code" xml:"code"`
	// response message
	Message string `form:"message" json:"message" xml:"message"`
}

// Validate validates the Message media type instance.
func (mt *Message) Validate() (err error) {
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}

	return
}

// Signatories media type (default view)
//
// Identifier: application/vnd.signatories+json; view=default
type Signatories struct {
	// response code
	Code int `form:"code" json:"code" xml:"code"`
	// response message
	Message     string   `form:"message" json:"message" xml:"message"`
	Signatories []string `form:"signatories" json:"signatories" xml:"signatories"`
}

// Validate validates the Signatories media type instance.
func (mt *Signatories) Validate() (err error) {
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}

	if mt.Signatories == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "signatories"))
	}
	return
}

// Transactions media type (default view)
//
// Identifier: application/vnd.transactions+json; view=default
type Transactions struct {
	// response code
	Code int `form:"code" json:"code" xml:"code"`
	// response message
	Message      string         `form:"message" json:"message" xml:"message"`
	Transactions []*Transaction `form:"transactions" json:"transactions" xml:"transactions"`
}

// Validate validates the Transactions media type instance.
func (mt *Transactions) Validate() (err error) {
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}

	if mt.Transactions == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "transactions"))
	}
	for _, e := range mt.Transactions {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
