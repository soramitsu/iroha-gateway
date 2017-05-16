package model

type Currency struct {
	CurrencyName string
	DomainName   string
	LedgerName   string

	Description string

	Amount    string
	Precision byte
}
