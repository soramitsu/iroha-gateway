package service

import (
	"testing"

	"github.com/soramitsu/iroha-gateway/command"
	"github.com/soramitsu/iroha-gateway/model"
)

func TestCreateCurrency(t *testing.T) {
	tx := &model.Transaction{
		Command: &command.Currency{
			CurrencyName: "test_currency",
			DomainName:   "test_domain",
			LedgerName:   "test_ledger_name",
			Description:  "test_description",
			Amount:       "10039",
			Precision:    2,
		},
	}

	CreateCurrency(tx)
}
