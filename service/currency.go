package service

import (
	"context"
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-gateway/iroha"
	"github.com/soramitsu/iroha-gateway/model"
	"google.golang.org/grpc"
)

func CreateCurrency(tx *model.Transaction) {
	builder := flatbuffers.NewBuilder(0)

	currency := tx.Command.Serialize(builder)
	pubkey := builder.CreateString("rI9Bks2reclulb+3/RENiouWSNaBHbRH6wo7BUoQ1Tc=")

	iroha.TransactionStart(builder)
	iroha.TransactionAddCommandType(builder, 4)
	iroha.TransactionAddCommand(builder, currency)
	iroha.TransactionAddCreatorPubKey(builder, pubkey)
	iroha.TransactionAddSignatures(builder, 2)
	// Add more some items ...

	transaction := iroha.TransactionEnd(builder)
	builder.Finish(transaction)

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := iroha.NewSumeragiClient(cc)
	ctx := context.Background()
	res, err := client.Torii(ctx, builder)
	fmt.Println(res.Code())
	fmt.Println(res.Message())
}
