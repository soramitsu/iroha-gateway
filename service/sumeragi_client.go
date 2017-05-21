package service

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-gateway/iroha"
)

type sumeragiClient struct {
	// ctx     context.Context
	builder *flatbuffers.Builder
	client  iroha.SumeragiClient
}

func (c *sumeragiClient) Torii() (*TransactionResponse, error) {
	ctx := context.Background()

	res, err := c.client.Torii(ctx, c.builder)
	if err != nil {
		return nil, fmt.Errorf("failed to connect torii: %s", err)
	}

	return &TransactionResponse{
		Code:    res.Code(),
		Message: string(res.Message()),
	}, nil
}

func newSumeragiClient(builder *flatbuffers.Builder) (*sumeragiClient, error) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("gRPC dial failed: %s", err)
	}
	client := iroha.NewSumeragiClient(cc)

	return &sumeragiClient{
		builder: builder,
		client:  client,
	}, nil
}
