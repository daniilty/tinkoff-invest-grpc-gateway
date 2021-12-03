package server

import (
	"fmt"
	"strconv"
	"strings"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertIntervalToTinkoff(interval uint32) (invest.CandleInterval, error) {
	converted, ok := intervals[interval]
	if !ok {
		available := make([]string, 0, len(intervals))

		for v := range intervals {
			available = append(available, strconv.Itoa(int(v)))
		}

		availableJoined := strings.Join(available, ", ")

		return "", fmt.Errorf("invalid interval: %d, available are: %s", interval, availableJoined)
	}

	return converted, nil
}

func convertTinkoffCandlesToGRPC(candles []invest.Candle) []*schema.Candle {
	converted := make([]*schema.Candle, 0, len(candles))

	for i := range candles {
		converted = append(converted, convertTinkoffCandleToGRPC(candles[i]))
	}

	return converted
}

func convertTinkoffCandleToGRPC(candle invest.Candle) *schema.Candle {
	return &schema.Candle{
		Interval:   string(candle.Interval),
		Ts:         timestamppb.New(candle.TS),
		OpenPrice:  candle.OpenPrice,
		ClosePrice: candle.ClosePrice,
		LowPrice:   candle.LowPrice,
		HighPrice:  candle.HighPrice,
		Volume:     candle.Volume,
	}
}
