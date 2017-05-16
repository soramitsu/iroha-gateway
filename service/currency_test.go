package service

import (
	"strings"
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
		Hash: "0xdeadbeef",
	}

	res, err := CreateCurrency(tx)
	if err != nil {
		t.Fatalf("CreateCurrency: %s", err)
	}

	if !strings.Contains(res.Message, "OK") {
		t.Fatalf("CreateCurrency response: code: %d, message: %s", res.Code, res.Message)
	}
}
