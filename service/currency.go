package service

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-gateway/iroha"
	"github.com/soramitsu/iroha-gateway/model"
)

func CreateCurrency(tx *model.Transaction) (*TransactionResponse, error) {
	builder := flatbuffers.NewBuilder(0)

	currency := tx.Command.Serialize(builder)
	pubkey := builder.CreateString("rI9Bks2reclulb+3/RENiouWSNaBHbRH6wo7BUoQ1Tc=")
	hash := builder.CreateString(tx.Hash)

	iroha.TransactionStart(builder)
	iroha.TransactionAddCommandType(builder, 4)
	iroha.TransactionAddCommand(builder, currency)
	iroha.TransactionAddCreatorPubKey(builder, pubkey)
	iroha.TransactionAddSignatures(builder, 2)
	iroha.TransactionAddHash(builder, hash)
	// Add more some items ...

	transaction := iroha.TransactionEnd(builder)
	builder.Finish(transaction)

	client, err := NewSumeragiClient(builder)
	if err != nil {
		return nil, fmt.Errorf("failed to connect torii: %s", err)
	}

	// client := iroha.NewSumeragiClient(cc)
	// ctx := context.Background()
	// res, err := client.Torii(ctx, builder)
	return client.Torii()
}
