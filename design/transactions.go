package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("transactions", func() {
	BasePath("/transactions")
	Action("getAll", func() {
		Routing(GET("/:currency_uri"))
		Params(func() {
			Param("currency_uri", String, func() {
				Description(descriptionCurrencyURI)
				Example(exampleCurrencyURI)
			})
			Param("target", String, func() {
				Pattern(`[0-9a-zA-Z-_.~]+`)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded target's account")
			})
			Param("creator_pubkey", String, func() {
				Description(descriptionCreator)
				Example(exampleCreator)
			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})
			Required("currency_uri", "target", "creator_pubkey")
		})

		Response(OK, TransactionsResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

})

var Transaction = Type("Transaction", func() {
	Attribute("command", String, func() {
		Description(descriptionTransactionCommand)
		Example(exampleTransactionCommand)
	})
	Attribute("creator_pubkey", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("signatures", ArrayOf(Signature))
	Attribute("timestamp", String, func() {
		Description(descriptionTimestamp)
		Example(exampleTimestamp)
		Pattern(`[0-9]{1,18}`)
	})

	Required("command", "creator_pubkey", "signatures", "timestamp")
})

var Signature = Type("Signature", func() {
	Attribute("publicKey", String, func() {
		Description(descriptionTransactionPublicKey)
		Example(exampleTransactionPublicKey)
		Pattern(patternBase64)
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

	Required("publicKey", "signature", "timestamp")
})

var TransactionsResponse = MediaType("application/vnd.transactions+json", func() {
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})
		Attribute("transactions", ArrayOf(Transaction))

		Required("message", "code", "transactions")
	})

	View("default", func() {
		Attribute("message")
		Attribute("code")
		Attribute("transactions")
	})
})
