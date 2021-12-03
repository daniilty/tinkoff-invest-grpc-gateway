package server

import (
	"time"

	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

const (
	interval1Min = 60
	interval2Min = interval1Min * 2
	interval3Min = interval1Min * 3
	interval5Min = interval1Min * 5

	interval10Min = interval1Min * 10
	interval15Min = interval1Min * 15
	interval30Min = interval1Min * 30

	interval1Hour = interval1Min * 60

	intervalDay = interval1Hour * 24

	intervalWeek = intervalDay * 7

	timeDay  = 24 * time.Hour
	timeWeek = 7 * timeDay
	timeYear = 365 * timeDay
)

var (
	intervals = map[uint32]invest.CandleInterval{
		interval1Min: invest.CandleInterval1Min,
		interval2Min: invest.CandleInterval2Min,
		interval3Min: invest.CandleInterval3Min,
		interval5Min: invest.CandleInterval5Min,

		interval10Min: invest.CandleInterval10Min,
		interval15Min: invest.CandleInterval15Min,
		interval30Min: invest.CandleInterval30Min,

		interval1Hour: invest.CandleInterval1Hour,

		intervalDay: invest.CandleInterval1Day,

		intervalWeek: invest.CandleInterval1Week,

		intervalDay * 28: invest.CandleInterval1Month,
		intervalDay * 29: invest.CandleInterval1Month,
		intervalDay * 30: invest.CandleInterval1Month,
		intervalDay * 31: invest.CandleInterval1Month,
	}

	maxIntervals = map[uint32]time.Duration{
		interval1Min: timeDay,
		interval2Min: timeDay,
		interval3Min: timeDay,
		interval5Min: timeDay,

		interval10Min: timeDay,
		interval15Min: timeDay,
		interval30Min: timeDay,

		interval1Hour: timeWeek,

		intervalDay:  timeYear,
		intervalWeek: 2 * timeYear,

		intervalDay * 28: 10 * timeYear,
		intervalDay * 29: 10 * timeYear,
		intervalDay * 30: 10 * timeYear,
		intervalDay * 31: 10 * timeYear,
	}
)

func getMaxInterval(interval uint32) time.Duration {
	return maxIntervals[interval]
}
