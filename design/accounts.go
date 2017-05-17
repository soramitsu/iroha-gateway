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
		Routing(GET("/:target"))
		Description("get an account by public key")
		Params(func() {
			Param("target", String, func() {
				Pattern(`[0-9a-zA-Z-_.~]+`)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded target's account")
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

			Required("target")
		})
		Response(OK, AccountResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("update", func() {
		Routing(PUT("/:target"))
		Description("update an account")
		Params(func() {
			Param("target", String, func() {
				Pattern(`[0-9a-zA-Z-_.~]+`)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded target's account")
			})
		})
		Payload(UpdateAccountRequest)
		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("delete", func() {
		Routing(DELETE("/:target"))
		Description("delete an account")
		Params(func() {
			Param("target", String, func() {
				Pattern(`[0-9a-zA-Z-_.~]+`)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded target's account")
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

var Account = Type("Account", func() { // TODO:
	Attribute("alias", String, func() {
		Description(descriptionAlias)
		Example(exampleAlias)
		Pattern(`[0-9a-zA-Z]+`)
	})
	Attribute("pubkey", String, func() {
		Description(descriptionPubkey)
		Example(examplePubkey)
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

	Required("alias", "pubkey", "signatories", "quorum")
})

var AccountFull = Type("AccountFull", func() {
	Attribute("alias", String, func() {
		Description(descriptionAlias)
		Example(exampleAlias)
		Pattern(`[0-9a-zA-Z]+`)
	})
	Attribute("pubkey", String, func() {
		Description(descriptionPubkey)
		Example(examplePubkey)
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
	Attribute("permissions", Permissions, func() {
		Description("account permissions")
	})

	Required("alias", "pubkey", "signatories", "quorum", "permissions")
})

var AccountResponse = MediaType("application/vnd.an.account+json", func() {
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})
		Attribute("account", AccountFull, func() {
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
		Attribute("accounts", ArrayOf(AccountFull), func() {
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
