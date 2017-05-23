package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("signatories", func() {
	BasePath("/accounts/:uuid")
	Params(func() {
		Param("uuid", String, func() {
			Pattern(patternAccountUUID)
			Example(exampleAccountUUID)
			Description(descriptionAccountUUID)
		})
		Required("uuid")
	})

	Action("getAll", func() {
		Routing(GET("/signatories"))
		Params(func() {
			Param("creator_pubkey", String, func() {
				Pattern(patternURLEncoded)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded creator's account")

			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})
		})

		Response(OK, Signatories)
		Response(BadRequest, Message)
		Response(InternalServerError, Message)
	})

	Action("add", func() {
		Routing(POST("/signatories"))
		Payload(SignatoriesRequest)

		Response(OK, Message)
		Response(BadRequest, Message)
		Response(InternalServerError, Message)
	})

	Action("delete", func() {
		Routing(DELETE("/signatories/:sig"))
		Params(func() {
			Param("sig", String, func() {
				Description(descriptionSignatory)
				Example(exampleSignatory)
				Pattern(patternURLEncoded)
			})
			Required("sig")
		})
		Payload(DeleteSignatoryRequest)
		Response(OK, Message)
		Response(BadRequest, Message)
		Response(InternalServerError, Message)
	})
})

var SignatoriesRequest = Type("SignatoryRequest", func() {
	Attribute("signatories", ArrayOf(String), func() {
		Example(exampleSignatories)
		Description(descriptionSignatories)
	})
	Attribute("creator_pubkey", String, func() {
		Example(exampleCreator)
		Description(descriptionCreator)
		Pattern(patternBase64)
	})
	Attribute("signature", String, func() {
		Example(exampleSignature)
		Description(descriptionSignature)
		Pattern(patternBase64)
	})
	Attribute("timestamp", String, func() {
		Example(exampleTimestamp)
		Description(descriptionTimestamp)
		Pattern(patternTimestamp)
	})

	Required("signatories", "creator_pubkey", "signature", "timestamp")
})

var Signatories = MediaType("application/vnd.signatories+json", func() {
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})
		Attribute("signatories", ArrayOf(String), func() {
			Example(exampleSignatories)
		})

		Required("message", "code", "signatories")
	})

	View("default", func() {
		Attribute("message")
		Attribute("code")
		Attribute("signatories")
	})
})

var DeleteSignatoryRequest = Type("DeleteSignatoryRequest", func() {
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
