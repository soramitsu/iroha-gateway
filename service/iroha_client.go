package service

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-gateway/iroha"
)

type SumeragiClient interface {
	Torii() (*TransactionResponse, error)
}

type TestSumeragiClient struct{}

func (c *TestSumeragiClient) Torii() (*TransactionResponse, error) {
	return &TransactionResponse{}, nil
}

type SumeragiClientAdapter struct {
	// ctx     context.Context
	builder *flatbuffers.Builder
	client  iroha.SumeragiClient
}

func (c *SumeragiClientAdapter) Torii() (*TransactionResponse, error) {
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

func NewSumeragiClient(builder *flatbuffers.Builder) (SumeragiClient, error) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("gRPC dial failed: %s", err)
	}
	client := iroha.NewSumeragiClient(cc)

	return &SumeragiClientAdapter{
		builder: builder,
		client:  client,
	}, nil
}
