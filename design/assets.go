package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("assets", func() {
	BasePath("/assets")

	Action("getAll", func() {
		Routing(GET("/:asset_uri"))
		Params(func() {
			Param("asset_uri", String, func() {
				Description(descriptionAssetURI)
				Example(exampleAssetURI)
			})
			Param("target", String, func() {
				Pattern(`[0-9a-zA-Z-_.~]+`)
				Example(exampleTargetEncoded)
				Description("Public key of URL-encoded target's account")
			})
			Param("creator", String, func() {
				Description(descriptionCreator)
				Example(exampleCreator)
			})
			Param("is_committed", Boolean, func() {
				Description("If this value is true, you can only get transactions committed to ametsuchi")
				Example(false)
			})
			Required("asset_uri", "target", "creator")
		})

		Response(OK, AssetsResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("add", func() {
		Routing(POST("/"))
		Payload(CreateAssetRequest)
		Response(Created, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("update", func() {
		Routing(PUT("/:asset_uri"))
		Params(func() {
			Param("asset_uri", String, func() {
				Description(descriptionAssetURI)
				Example(exampleAssetURI)
			})
		})

		Payload(UpdateAssetRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("delete", func() {
		Routing(DELETE("/:asset_uri"))
		Params(func() {
			Param("asset_uri", String, func() {
				Description(descriptionAssetURI)
				Example(exampleAssetURI)
			})
		})

		Payload(DeleteAssetRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("addValue", func() {
		Routing(POST("/:asset_uri/add"))
		Params(func() {
			Param("asset_uri", String, func() {
				Description(descriptionAssetURI)
				Example(exampleAssetURI)
			})
		})

		Payload(AssetValueRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("subtractValue", func() {
		Routing(POST("/:asset_uri/subtract"))
		Params(func() {
			Param("asset_uri", String, func() {
				Description(descriptionAssetURI)
				Example(exampleAssetURI)
			})
		})

		Payload(AssetValueRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})

	Action("transfer", func() {
		Routing(POST("/:asset_uri/transfer"))
		Params(func() {
			Param("asset_uri", String, func() {
				Description(descriptionAssetURI)
				Example(exampleAssetURI)
			})
		})

		Payload(AssetTransferRequest)

		Response(OK, MessageResponse)
		Response(BadRequest, MessageResponse)
		Response(InternalServerError, MessageResponse)
	})
})

var CreateAssetRequest = Type("CreateAssetRequest", func() {
	Attribute("target", String, func() {
		Pattern(`[0-9a-zA-Z-_.~]+`)
		Example(exampleTargetEncoded)
		Description("Public key of URL-encoded target's account")
	})
	Attribute("creator", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("hash", String, func() {
		Description(descriptionHash)
		Example(exampleHash)
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
	Attribute("asset", Asset)

	Required("target", "creator", "hash", "signature", "timestamp", "asset")
})

var UpdateAssetRequest = Type("UpdateAssetRequest", func() {
	Attribute("description", String, func() {
		Example(exampleAssetDescription)
		Description(descriptionAssetDescription)
	})
	Attribute("creator", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("hash", String, func() {
		Description(descriptionHash)
		Example(exampleHash)
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

	Required("description", "creator", "hash", "signature", "timestamp")
})

var DeleteAssetRequest = Type("DeleteAssetRequest", func() {
	Attribute("creator", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("hash", String, func() {
		Description(descriptionHash)
		Example(exampleHash)
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

	Required("creator", "hash", "signature", "timestamp")
})

var AssetValueRequest = Type("AssetValueRequest", func() {
	Attribute("value", Number, func() {
		Description(descriptionAssetValue)
		Example(exampleAssetValue)
	})
	Attribute("target", String, func() {
		Pattern(`[0-9a-zA-Z-_.~]+`)
		Description("Public key of URL-encoded target's account")
		Example(exampleTargetEncoded)
	})
	Attribute("creator", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("hash", String, func() {
		Description(descriptionHash)
		Example(exampleHash)
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

	Required("value", "target", "creator", "hash", "signature", "timestamp")
})

var AssetTransferRequest = Type("AssetTransferRequest", func() {
	Attribute("value", Number, func() {
		Description(descriptionAssetValue)
		Example(exampleAssetValue)
	})
	Attribute("sender", String, func() {
		Description(descriptionAssetSender)
		Example(exampleAssetSender)
		Pattern(base64Pattern)
	})
	Attribute("receiver", String, func() {
		Description(descriptionAssetReceiver)
		Example(exampleAssetReceiver)
		Pattern(base64Pattern)
	})
	Attribute("target", String, func() {
		Pattern(`[0-9a-zA-Z-_.~]+`)
		Description("Public key of URL-encoded target's account")
		Example(exampleTargetEncoded)
	})
	Attribute("creator", String, func() {
		Description(descriptionCreator)
		Example(exampleCreator)
	})
	Attribute("hash", String, func() {
		Description(descriptionHash)
		Example(exampleHash)
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

	Required("value", "sender", "receiver", "target", "creator", "hash", "signature", "timestamp")

})

var AssetsResponse = MediaType("application/vnd.assets+json", func() {
	Attributes(func() {
		Attribute("message", String, func() {
			Description(descriptionMessage)
			Example(exampleMessage)
		})
		Attribute("code", Integer, func() {
			Description(descriptionCode)
			Example(exampleCode)
		})
		Attribute("assets", ArrayOf(Asset))

		Required("message", "code", "assets")
	})

	View("default", func() {
		Attribute("message")
		Attribute("code")
		Attribute("assets")
	})
})

var Asset = Type("Asset", func() {
	Attribute("name", String, func() {
		Example(exampleAssetName)
		Description(descriptionAssetName)
	})
	Attribute("domain_name", String, func() {
		Example(exampleAssetDomainName)
		Description(descriptionAssetDomainName)
	})
	Attribute("ledger_name", String, func() {
		Example(exampleAssetLedgerName)
		Description(descriptionAssetLedgerName)
	})
	Attribute("description", String, func() {
		Example(exampleAssetDescription)
		Description(descriptionAssetDescription)
	})

	Required("name", "domain_name", "ledger_name")
})
