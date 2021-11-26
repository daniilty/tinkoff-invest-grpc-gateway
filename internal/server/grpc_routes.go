package server

import (
	"context"

	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GRPC) GetCandles(ctx context.Context, req *schema.CandlesRequest) (*schema.CandlesResponse, error) {
	from := req.From.AsTime()
	to := req.To.AsTime()

	tinkoffInterval, err := convertIntervalToTinkoff(req.GetInterval())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = validateIntervals(from, to, int64(req.GetInterval()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	candles, err := g.client.Candles(ctx, from, to, tinkoffInterval, req.GetFigi())
	if err != nil {
		g.logger.Errorw("Failed to get candles.", "err", err)

		return nil, status.Error(codes.Internal, "failed to get candles")
	}

	candlesGRPC := convertTinkoffCandlesToGRPC(candles)

	return &schema.CandlesResponse{
		Candles: candlesGRPC,
	}, nil
}
