package server

import (
	invest "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var (
	intervals = map[uint32]invest.CandleInterval{
		60:     invest.CandleInterval1Min,
		60 * 2: invest.CandleInterval2Min,
		60 * 3: invest.CandleInterval3Min,
		60 * 5: invest.CandleInterval5Min,

		60 * 10: invest.CandleInterval10Min,
		60 * 15: invest.CandleInterval15Min,
		60 * 30: invest.CandleInterval30Min,

		60 * 60:     invest.CandleInterval1Hour,
		60 * 60 * 2: invest.CandleInterval2Hour,
		60 * 60 * 4: invest.CandleInterval4Hour,

		60 * 60 * 24: invest.CandleInterval1Day,

		60 * 60 * 24 * 7: invest.CandleInterval1Week,

		60 * 60 * 24 * 28: invest.CandleInterval1Month,
		60 * 60 * 24 * 29: invest.CandleInterval1Month,
		60 * 60 * 24 * 30: invest.CandleInterval1Month,
		60 * 60 * 24 * 31: invest.CandleInterval1Month,
	}
)
