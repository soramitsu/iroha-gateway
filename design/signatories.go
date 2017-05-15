package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("signatories", func() {
	BasePath("/accounts/:target")
	Params(func() {
		Param("target", String, func() {
			Pattern(`[0-9a-zA-Z-_.~]+`)
			Example(exampleTargetEncoded)
			Description("Public key of URL-encoded target's account")
		})
		Required("target")
	})

	Action("getAll", func() {
		Routing(GET("/signatories"))
		Params(func() {
			Param("creator", String, func() {
				Pattern(urlEncodedPattern)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded creator's account")

			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})
		})

		Response(OK, SignatoriesResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("add", func() {
		Routing(POST("/signatories"))
		Payload(SignatoriesRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("delete", func() {
		Routing(DELETE("/signatories/:sig"))
		Params(func() {
			Param("sig", String, func() {
				Description(descriptionSignatory)
				Example(exampleSignatory)
				Pattern(urlEncodedPattern)
			})
			Required("sig")
		})
		Payload(DeleteSignatoryRequest)
		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})
})

var SignatoriesRequest = Type("SignatoryRequest", func() {
	Attribute("signatories", ArrayOf(String), func() {
		Example(exampleSignatories)
		Description(descriptionSignatories)
	})
	Attribute("creator", String, func() {
		Example(exampleCreator)
		Description(descriptionCreator)
		Pattern(base64Pattern)
	})
	Attribute("signature", String, func() {
		Example(exampleSignature)
		Description(descriptionSignature)
		Pattern(base64Pattern)
	})
	Attribute("timestamp", String, func() {
		Example(exampleTimestamp)
		Description(descriptionTimestamp)
		Pattern(timestampPattern)
	})

	Required("signatories", "creator", "signature", "timestamp")
})

var SignatoriesResponse = MediaType("application/vnd.signatories+json", func() {
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
	Attribute("creator", String, func() {
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

	Required("creator", "signature", "timestamp")
})
