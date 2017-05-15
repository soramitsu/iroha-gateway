package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("quorum", func() {
	BasePath("/accounts/:target/quorum")
	Params(func() {
		Param("target", String, func() {
			Pattern(`[0-9a-zA-Z-_.~]+`)
			Example(exampleTargetEncoded)
			Description("Public key of URL-encoded target's account")
		})
		Required("target")
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

	Required("quorum", "creator", "signature", "timestamp")
})
