package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("quorum", func() {
	BasePath("/accounts/:uuid/quorum")
	Params(func() {
		Param("uuid", String, func() {
			Pattern(patternAccountUUID)
			Example(exampleAccountUUID)
			Description(descriptionAccountUUID)
		})
		Required("uuid")
	})

	Action("update", func() {
		Routing(PUT("/"))
		Payload(UpdateQuorumRequest)
		Response(OK, Message)
		Response(BadRequest, Message)
		Response(InternalServerError, Message)
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
