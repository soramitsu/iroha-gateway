package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("account", func() {

	// GET
	Action("getAll", func() {
		Routing(GET("/accounts/"))
		Description("get all accounts")
		Response(OK, AccountsResponse)
	})

	Action("getByUUID", func() {
		Routing(GET("/accounts/:uuid"))
		Description("get an account by account's UUID")
		Params(func() {
			Param("uuid", String, func() {
				Pattern(patternAccountUUID)
				Example(exampleAccountUUID)
				Description("account's UUID")
			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})
			Required("uuid")
		})
		Response(OK, AccountResponse)
	})

	Action("getByUsername", func() {
		Routing(GET("/domains/:domain_uri/accounts/:username"))
		Description("get an account by account's domain and username")
		Params(func() {
			Param("domain_uri", String, func() {
				Pattern(patternDomainURI)
				Example(exampleDomainURI)
				Description(descriptionDomainURI)
			})
			Param("username", String, func() {
				Pattern(patternAccountUsername)
				Example(exampleAccountUsername)
				Description(descriptionAccountUsername)
			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})
			Required("domain_uri", "username")
		})
		Response(OK, AccountResponse)
	})

	// TODO: Consider a good URL path when get an account with only a user name from the default domain
	Action("getByUsernameFromDefaultDomain", func() {
		Routing(GET("/domains/default/accounts/:username"))
		Description("get an account by account's username from default domain")
		Params(func() {
			Param("username", String, func() {
				Pattern(patternAccountUsername)
				Example(exampleAccountUsername)
				Description(descriptionAccountUsername)
			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})
			Required("domain_uri", "username")
		})
		Response(OK, AccountResponse)
	})

	// UPDATE
	Action("update", func() {
		Routing(PUT("/accounts/:uuid"))
		Description("update an account by account's uuid")
		Params(func() {
			Param("uuid", String, func() {
				Pattern(patternAccountUUID)
				Example(exampleAccountUUID)
				Description(descriptionAccountUUID)
			})
			Required("uuid")
		})
		Payload(UpdateAccountRequest)
		Response(OK, MessageResponse)
	})

	Action("updateByUsername", func() {
		Routing(PUT("/domains/:domain_uri/accounts/:username"))
		Description("update an account by account's domain and username")
		Params(func() {
			Param("domain_uri", String, func() {
				Pattern(patternDomainURI)
				Example(exampleDomainURI)
				Description(descriptionDomainURI)
			})
			Param("username", String, func() {
				Pattern(patternAccountUsername)
				Example(exampleAccountUsername)
				Description(descriptionAccountUsername)
			})
			Required("domain_uri", "username")
		})
		Payload(UpdateAccountRequest)
		Response(OK, MessageResponse)
	})

	// TODO: Consider a good URL path when update an account with only a user name from the default domain
	Action("updateByUsernameFromDefaultDomain", func() {
		Routing(PUT("/domains/default/accounts/:username"))
		Description("update an account by account's username from default domain")
		Params(func() {
			Param("username", String, func() {
				Pattern(patternAccountUsername)
				Example(exampleAccountUsername)
				Description(descriptionAccountUsername)
			})
			Required("username")
		})
		Payload(UpdateAccountRequest)
		Response(OK, MessageResponse)
	})

	// DELETE
	Action("delete", func() {
		Routing(DELETE("/accounts/:uuid"))
		Description("delete an account")
		Params(func() {
			Param("uuid", String, func() {
				Pattern(patternAccountUUID)
				Example(exampleAccountUUID)
				Description(descriptionAccountUUID)
			})
		})
		Payload(DeleteAccountRequest)
		Response(OK, MessageResponse)
	})

	Action("deleteByUsername", func() {
		Routing(DELETE("/domains/:domain_uri/accounts/:username"))
		Description("delete an account by account's domain and username")
		Params(func() {
			Param("domain_uri", String, func() {
				Pattern(patternDomainURI)
				Example(exampleDomainURI)
				Description(descriptionDomainURI)
			})
			Param("username", String, func() {
				Pattern(patternAccountUsername)
				Example(exampleAccountUsername)
				Description(descriptionAccountUsername)
			})
			Required("domain_uri", "username")
		})
		Payload(DeleteAccountRequest)
		Response(OK, MessageResponse)
	})

	// TODO: Consider a good URL path when delete an account with only a user name from the default domain
	Action("deleteByUsernameFromDefaultDomain", func() {
		Routing(DELETE("/domains/default/accounts/:username"))
		Description("delete an account by account's username from default domain")
		Params(func() {
			Param("username", String, func() {
				Pattern(patternAccountUsername)
				Example(exampleAccountUsername)
				Description(descriptionAccountUsername)
			})
			Required("username")
		})
		Payload(DeleteAccountRequest)
		Response(OK, MessageResponse)
	})

	// CREATE
	Action("add", func() {
		Routing(POST("/accounts"))
		Description("add an account")
		Payload(AddAccountRequest)
		Response(Created, MessageResponse)
	})

	Response(BadRequest, MessageResponse)
	Response(InternalServerError, MessageResponse)
})

var UpdateAccountRequest = Type("UpdateAccountRequest", func() {
	Attribute("username", String, func() {
		Description(descriptionAccountUsername)
		Example(exampleAccountUsername)
		Pattern(patternAccountUsername)
	})

	Attribute("meta_transaction", TransactionRequest, func() {
		Description(descriptionMetaTransaction)
	})

	Required("username", "meta_transaction")
})

var DeleteAccountRequest = Type("DeleteAccountRequest", func() {
	Attribute("meta_transaction", TransactionRequest, func() {
		Description(descriptionMetaTransaction)
	})

	Required("meta_transaction")
})

var AddAccountRequest = Type("AddAccountRequest", func() {
	Attribute("meta_transaction", TransactionRequest, func() {
		Description(descriptionMetaTransaction)
	})
	Attribute("account", Account)

	Required("meta_transaction", "account")
})

var Account = Type("Account", func() {
	Attribute("uuid", String, func() {
		Description(descriptionAccountUUID)
		Example(exampleAccountUUID)
		Pattern(patternBase64)
	})
	Attribute("signatories", ArrayOf(String), func() {
		Description(descriptionSignatories)
		Example(exampleSignatories)
	})
	Attribute("quorum", Integer, func() {
		Description(descriptionQuorum)
		Example(exampleQuorum)
		Minimum(1)
		Maximum(32)
	})

	Required("uuid", "signatories", "quorum")
})

var AccountResponse = MediaType("application/vnd.account+json", func() {
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})
		Attribute("account", Account, func() {
			Description(descriptionAccount)
		})

		Required("message", "code", "account")
	})

	View("default", func() {
		Attribute("message")
		Attribute("code")
		Attribute("account")
	})
})

var AccountsResponse = MediaType("application/vnd.accounts+json", func() {
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})
		Attribute("accounts", ArrayOf(Account), func() {
			Description(descriptionAccounts)
		})

		Required("message", "code", "accounts")
	})

	View("default", func() {
		Attribute("message")
		Attribute("code")
		Attribute("accounts")
	})
})
