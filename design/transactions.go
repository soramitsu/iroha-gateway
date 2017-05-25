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

		Response(OK, Transactions)
		Response(BadRequest, Message)
		Response(InternalServerError, Message)
	})

})

var Transaction = Type("Transaction", func() {
	Attribute("creator_signature", String, func() {
		Description(descriptionTransactionSignature)
		Example(exampleSignature)
	})

	Attribute("signatures", ArrayOf(Signature))

	Attribute("created_at", String, func() {
		Description(descriptionTransactionTimestamp)
		Example(exampleTimestamp)
		Pattern(patternTimestamp)
	})

	Attribute("nonce", Integer, func() {
		Description(descriptionTransactionNonce)
		Example(exampleTransactionNonce)
	})

	Required("creator_signature", "signatures", "created_at", "nonce")
})

var Signature = Type("Signature", func() {
	Attribute("pubkey", String, func() {
		Description(descriptionTransactionPublicKey)
		Example(exampleTransactionPublicKey)
		Pattern(patternBase64)
	})
	Attribute("signature", String, func() {
		Description(descriptionSignature)
		Example(exampleSignature)
	})

	Required("pubkey", "signature")
})

var TransactionRequest = Type("TransactionRequest", func() {
	Attribute("pubkey", String, func() {
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

	Required("pubkey", "signature", "timestamp")

})

var Transactions = MediaType("application/vnd.transactions+json", func() {
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
