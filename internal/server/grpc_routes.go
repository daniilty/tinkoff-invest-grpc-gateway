package server

import (
	"context"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/daniilty/tinkoff-invest-grpc-gateway/internal/interval"
	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCandles - get candles(анальные) and split request if needed.
// TODO: выгребсти весь говнокод отсюда в отдельный корневой модуль и оставить только вызов к нему и конвертацию
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

	requestInterval := &interval.Interval{
		From: from,
		To:   to,
	}

	max := getMaxInterval(req.GetInterval())
	intervals := requestInterval.Split(max)

	candles := []sdk.Candle{}

	for i := range intervals {
		cc, err := g.client.Candles(ctx, intervals[i].From, intervals[i].To, tinkoffInterval, req.GetFigi())
		if err != nil {
			g.logger.Errorw("Failed to get candles.", "err", err)

			return nil, status.Error(codes.Internal, "failed to get candles")
		}

		candles = append(candles, cc...)
	}

	candlesGRPC := convertTinkoffCandlesToGRPC(candles)

	return &schema.CandlesResponse{
		Candles: candlesGRPC,
	}, nil
}
