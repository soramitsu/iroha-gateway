package action

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-gateway/iroha"
)

type Currency struct {
	CurrencyName string
	DomainName   string
	LedgerName   string

	Description string

	Amount    string
	Precision byte
}

func (c *Currency) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	name := builder.CreateString(c.CurrencyName)
	domainName := builder.CreateString(c.DomainName)
	lederName := builder.CreateString(c.LedgerName)
	description := builder.CreateString(c.Description)
	amount := builder.CreateString(c.Amount)
	precision := c.Precision

	iroha.CurrencyStart(builder)
	iroha.CurrencyAddCurrencyName(builder, name)
	iroha.CurrencyAddDomainName(builder, domainName)
	iroha.CurrencyAddLedgerName(builder, lederName)
	iroha.CurrencyAddDescription(builder, description)
	iroha.CurrencyAddAmount(builder, amount)
	iroha.CurrencyAddPrecision(builder, precision)

	return iroha.CurrencyEnd(builder)
}
