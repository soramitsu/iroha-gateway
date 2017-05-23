package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var MessageResponse = MediaType("application/vnd.messageResponse+json", func() {
	Description("Basic response")
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})

		Required("message", "code")
	})

	View("default", func() {
		Attribute("message")
		Attribute("code")
	})

})
