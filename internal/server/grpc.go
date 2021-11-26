package server

import (
	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"go.uber.org/zap"
)

// GRPC - grpc server.
type GRPC struct {
	schema.UnimplementedTinkoffInvestAPIGatewayServer

	client *invest.RestClient
	logger *zap.SugaredLogger
}

// NewGRPC - constructor.
func NewGRPC(client *invest.RestClient, opts ...GRPCOption) *GRPC {
	g := &GRPC{
		client: client,
	}

	for i := range opts {
		opts[i](g)
	}

	return g
}
