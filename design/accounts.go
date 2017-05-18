package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("accounts", func() {
	BasePath("/accounts")
	Action("getAll", func() {
		Routing(GET("/"))
		Description("get all accounts")
		Response(OK, AccountsResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("get", func() {
		Routing(GET("/:guid"))
		Description("get an account by public key")
		Params(func() {
			Param("guid", String, func() {
				Pattern(base64Pattern)
				Example(exampleAccountGUIDEncoded)
				Description("GUID of URL-encoded account")
			})
			Param("creator_pubkey", String, func() {
				Pattern(`[0-9a-zA-Z-_.~]+`)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded creator's account")

			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})

			Required("guid")
		})
		Response(OK, AccountResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("update", func() {
		Routing(PUT("/:guid"))
		Description("update an account")
		Params(func() {
			Param("guid", String, func() {
				Pattern(base64Pattern)
				Example(exampleTargetEncoded)
				Description("GUID of URL-encoded account")
			})
			Required("guid")
		})
		Payload(UpdateAccountRequest)
		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("delete", func() {
		Routing(DELETE("/:guid"))
		Description("delete an account")
		Params(func() {
			Param("guid", String, func() {
				Pattern(base64Pattern)
				Example(exampleAccountGUIDEncoded)
				Description("GUID of URL-encoded account")
			})
		})
		Payload(DeleteAccountRequest)
		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("add", func() {
		Routing(POST("/"))
		Description("add an account")
		Payload(AddAccountRequest)
		Response(Created, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})
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
		Pattern(timestampPattern)
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
	Attribute("guid", String, func() {
		Description(descriptionAccountGUID)
		Example(exampleAccountGUID)
		Pattern(base64Pattern)
	})
	Attribute("signatories", ArrayOf(String), func() {
		Description(descriptionSignatories)
		Example(exampleSignatories)
	})
	Attribute("quorum", Integer, func() {
		Description(descriptionQuorum)
		Example(exampleQuorum)
		Minimum(1)
		Maximum(65535)
	})

	Required("guid", "signatories", "quorum")
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
