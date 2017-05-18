package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("quorum", func() {
	BasePath("/accounts/:guid/quorum")
	Params(func() {
		Param("guid", String, func() {
			Pattern(base64Pattern)
			Example(exampleTargetEncoded)
			Description("GUID of URL-encoded account")
		})
		Required("guid")
	})

	Action("update", func() {
		Routing(PUT("/"))
		Payload(UpdateQuorumRequest)
		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})
})

var UpdateQuorumRequest = Type("UpdateQuorumRequest", func() {
	Attribute("quorum", Integer, func() {
		Description(descriptionQuorum)
		Example(exampleQuorum)
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

	Required("quorum", "creator_pubkey", "signature", "timestamp")
})
