package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Permissions = Type("Permissions", func() {
	Attribute("root", PermissionRoot)
	Attribute("ledger", ArrayOf(PermissionLedger))
	Attribute("domain", ArrayOf(PermissionDomain))
	Attribute("asset", ArrayOf(PermissionAsset))

	Required("root", "ledger", "domain", "asset")
})

var PermissionRoot = Type("PermissionRoot", func() {
	Attribute("ledger_add", Boolean, func() {
		Description(descriptionPermissionLedgerAdd)
	})
	Attribute("ledger_remove", Boolean, func() {
		Description(descriptionPermissionLedgerRemove)
	})
})

var PermissionLedger = Type("PermissionLedger", func() {
	Attribute("ledger_name", String, func() {
		Description(descriptionPermissionLedgerName)
	})
	Attribute("domain_add", Boolean, func() {
		Description(descriptionPermissionDomainAdd)
	})
	Attribute("domain_remove", Boolean, func() {
		Description(descriptionPermissionDomainRemove)
	})
	Attribute("peer_read", Boolean, func() {
		Description(descriptionPermissionPeerRead)
	})
	Attribute("peer_write", Boolean, func() {
		Description(descriptionPermissionPeerWrite)
	})
	Attribute("account_add", Boolean, func() {
		Description(descriptionPermissionAccountAdd)
	})
	Attribute("account_remove", Boolean, func() {
		Description(descriptionPermissionAccountRemove)
	})
	Attribute("account_give_permission", Boolean, func() {
		Description(descriptionPermissionAccountGivePermission)
	})


	Required("ledger_name")
})

var PermissionDomain = Type("PermissionDomain", func() {
	Attribute("domain_name", String, func() {
		Description(descriptionPermissionDomainName)
		Example(examplePermissionDomainName)
	})
	Attribute("ledger_name", String, func() {
		Description(descriptionPermissionLedgerName)
		Example(examplePermissionLedgerName)
	})
	Attribute("account_give_permission", Boolean, func() {
		Description(descriptionPermissionAccountGivePermission)
	})
	Attribute("account_add", Boolean, func() {
		Description(descriptionPermissionAccountAdd)
	})
	Attribute("account_remove", Boolean, func() {
		Description(descriptionPermissionAccountRemove)
	})
	Attribute("asset_create", Boolean, func() {
		Description(descriptionPermissionAssetCreate)
	})
	Attribute("asset_remove", Boolean, func() {
		Description(descriptionPermissionAssetRemove)
	})
	Attribute("asset_update", Boolean, func() {
		Description(descriptionPermissionAssetUpdate)
	})

	Required("domain_name", "ledger_name")
})

var PermissionAsset = Type("PermissionAsset", func() {
	Attribute("asset_name", String, func() {
		Description(descriptionPermissionAssetName)
		Example(examplePermissionAssetName)
	})
	Attribute("domain_name", String, func() {
		Description(descriptionPermissionDomainName)
		Example(examplePermissionDomainName)
	})
	Attribute("ledger_name", String, func() {
		Description(descriptionPermissionLedgerName)
		Example(examplePermissionLedgerName)
	})
	Attribute("account_give_permission", Boolean, func() {
		Description(descriptionPermissionAccountGivePermission)
	})
	Attribute("account_give_permission", Boolean, func() {
		Description(descriptionPermissionAccountGivePermission)
	})
	Attribute("transfer", Boolean, func() {
		Description(descriptionPermissionTransfer)
	})
	Attribute("add", Boolean, func() {
		Description(descriptionPermissionAdd)
	})
	Attribute("subtract", Boolean, func() {
		Description(descriptionPermissionSubtract)
	})
	Attribute("read", Boolean, func() {
		Description(descriptionPermissionRead)
	})

	Required("asset_name", "domain_name", "ledger_name")
})

