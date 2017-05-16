package model

type Permissions struct {
	Root   AccountPermissionRoot
	Ledger []AccountPermissionLedger
	Domain []AccountPermissionDomain
	Asset  []AccountPermissionAsset
}

type AccountPermissionRoot struct {
	LedgerAdd    bool
	LedgerRemove bool
}

type AccountPermissionLedger struct {
	LedgerName            string
	DomainAdd             bool
	DomainRemove          bool
	PeerRead              bool
	PeerWrite             bool
	AccountAdd            bool
	AccountRemove         bool
	AccountGivePermission bool
}

type AccountPermissionDomain struct {
	DomainName            string
	LedgerName            string
	AccountGivePermission bool
	AccountAdd            bool
	AccountRemove         bool
	AssetCreate           bool
	AssetRemove           bool
	AssetUpdate           bool
}

type AccountPermissionAsset struct {
	AssetName             string
	DomainName            string
	LedgerName            string
	AccountGivePermission bool
	Transfer              bool
	Add                   bool
	Subtract              bool
	Read                  bool
}
