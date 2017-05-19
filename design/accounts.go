package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("account", func() {

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

	Action("getAll", func() {
		Routing(GET("/accounts/"))
		Description("get all accounts")
		Response(OK, AccountsResponse)
	})

	Action("update", func() {
		Routing(PUT("/accounts/:uuid"))
		Description("update an account")
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
	Attribute("alias", String, func() {
		Description(descriptionAlias)
		Example(exampleAlias)
		Pattern(`[0-9a-zA-Z]+`)
	})
	Attribute("creator_pubkey", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("signature", String, func() {
		Description(descriptionSignature)
		Example(exampleSignature)
	})
	Attribute("timestamp", String, func() {
		Description(descriptionTimestamp)
		Example(exampleTimestamp)
		Pattern(patternTimestamp)
	})

	Required("alias", "creator_pubkey", "signature", "timestamp")
})

var DeleteAccountRequest = Type("DeleteAccountRequest", func() {
	Attribute("creator_pubkey", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("signature", String, func() {
		Description(descriptionSignature)
		Example(exampleSignature)
	})
	Attribute("timestamp", String, func() {
		Description(descriptionTimestamp)
		Example(exampleTimestamp)
		Pattern(`[0-9]{1,18}`)
	})

	Required("creator_pubkey", "signature", "timestamp")

})

var AddAccountRequest = Type("AddAccountRequest", func() {
	Attribute("creator_pubkey", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("signature", String, func() {
		Description(descriptionSignature)
		Example(exampleSignature)
	})
	Attribute("timestamp", String, func() {
		Description(descriptionTimestamp)
		Example(exampleTimestamp)
		Pattern(`[0-9]{1,18}`)
	})
	Attribute("account", Account)

	Required("creator_pubkey", "signature", "timestamp", "account")
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
