package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("currency", func() {
	BasePath("/currency")

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

		Response(OK, CurrencyResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("add", func() {
		Routing(POST("/"))
		Payload(CreateCurrencyRequest)
		Response(Created, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("update", func() {
		Routing(PUT("/:currency_uri"))
		Params(func() {
			Param("currency_uri", String, func() {
				Description(descriptionCurrencyURI)
				Example(exampleCurrencyURI)
			})
		})

		Payload(UpdateCurrencyRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("delete", func() {
		Routing(DELETE("/:currency_uri"))
		Params(func() {
			Param("currency_uri", String, func() {
				Description(descriptionCurrencyURI)
				Example(exampleCurrencyURI)
			})
		})

		Payload(DeleteCurrencyRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("addValue", func() {
		Routing(POST("/:currency_uri/add"))
		Params(func() {
			Param("currency_uri", String, func() {
				Description(descriptionCurrencyURI)
				Example(exampleCurrencyURI)
			})
		})

		Payload(CurrencyValueRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("subtractValue", func() {
		Routing(POST("/:currency_uri/subtract"))
		Params(func() {
			Param("currency_uri", String, func() {
				Description(descriptionCurrencyURI)
				Example(exampleCurrencyURI)
			})
		})

		Payload(CurrencyValueRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("transfer", func() {
		Routing(POST("/:currency_uri/transfer"))
		Params(func() {
			Param("currency_uri", String, func() {
				Description(descriptionCurrencyURI)
				Example(exampleCurrencyURI)
			})
		})

		Payload(CurrencyTransferRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})
})

var CreateCurrencyRequest = Type("CreateCurrencyRequest", func() {
	Attribute("target", String, func() {
		Pattern(`[0-9a-zA-Z-_.~]+`)
		Example(exampleTargetEncoded)
		Description("Public key of URL-encoded target's account")
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
		Pattern(`[0-9]{1,18}`)
	})
	Attribute("currency", Currency)

	Required("target", "creator_pubkey", "signature", "timestamp", "currency")
})

var UpdateCurrencyRequest = Type("UpdateCurrencyRequest", func() {
	Attribute("description", String, func() {
		Example(exampleCurrencyDescription)
		Description(descriptionCurrencyDescription)
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
		Pattern(`[0-9]{1,18}`)
	})

	Required("description", "creator_pubkey", "signature", "timestamp")
})

var DeleteCurrencyRequest = Type("DeleteCurrencyRequest", func() {
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

var CurrencyValueRequest = Type("CurrencyValueRequest", func() {
	Attribute("value", Number, func() {
		Description(descriptionCurrencyValue)
		Example(exampleCurrencyValue)
	})
	Attribute("target", String, func() {
		Pattern(`[0-9a-zA-Z-_.~]+`)
		Description("Public key of URL-encoded target's account")
		Example(exampleTargetEncoded)
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
		Pattern(`[0-9]{1,18}`)
	})

	Required("value", "target", "creator_pubkey", "signature", "timestamp")
})

var CurrencyTransferRequest = Type("CurrencyTransferRequest", func() {
	Attribute("value", Number, func() {
		Description(descriptionCurrencyValue)
		Example(exampleCurrencyValue)
	})
	Attribute("sender", String, func() {
		Description(descriptionCurrencySender)
		Example(exampleCurrencySender)
		Pattern(patternBase64)
	})
	Attribute("receiver", String, func() {
		Description(descriptionCurrencyReceiver)
		Example(exampleCurrencyReceiver)
		Pattern(patternBase64)
	})
	Attribute("target", String, func() {
		Pattern(`[0-9a-zA-Z-_.~]+`)
		Description("Public key of URL-encoded target's account")
		Example(exampleTargetEncoded)
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
		Pattern(`[0-9]{1,18}`)
	})

	Required("value", "sender", "receiver", "target", "creator_pubkey", "signature", "timestamp")

})

var CurrencyResponse = MediaType("application/vnd.currencyResponse+json", func() {
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})
		Attribute("currency", ArrayOf(Currency))

		Required("message", "code", "currency")
	})

	View("default", func() {
		Attribute("message")
		Attribute("code")
		Attribute("currency")
	})
})

var Currency = Type("Currency", func() {
	Attribute("name", String, func() {
		Example(exampleCurrencyName)
		Description(descriptionCurrencyName)
	})
	Attribute("domain_name", String, func() {
		Example(exampleCurrencyDomainName)
		Description(descriptionCurrencyDomainName)
	})
	Attribute("ledger_name", String, func() {
		Example(exampleCurrencyLedgerName)
		Description(descriptionCurrencyLedgerName)
	})
	Attribute("description", String, func() {
		Example(exampleCurrencyDescription)
		Description(descriptionCurrencyDescription)
	})
	Attribute("value", Number, func() {
		Description(descriptionCurrencyValue)
		Example(exampleCurrencyValue)
	})

	Required("name", "domain_name", "ledger_name", "value")
})
