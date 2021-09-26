package public

import (
	"github.com/amir-the-h/okex/events"
	"github.com/amir-the-h/okex/models/market"
	"github.com/amir-the-h/okex/models/public_data"
)

type (
	Instruments struct {
		Arg         *events.Argument          `json:"arg"`
		Instruments []*public_data.Instrument `json:"data"`
	}
	Tickers struct {
		Arg     *events.Argument `json:"arg"`
		Tickers []*market.Ticker `json:"data"`
	}
	OpenInterest struct {
		Arg           *events.Argument            `json:"arg"`
		OpenInterests []*public_data.OpenInterest `json:"data"`
	}
	Candlesticks struct {
		Arg     *events.Argument `json:"arg"`
		Candles []*market.Candle `json:"data"`
	}
	Trades struct {
		Arg    *events.Argument `json:"arg"`
		Trades []*market.Trade  `json:"data"`
	}
	EstimatedDeliveryExercisePrice struct {
		Arg                             *events.Argument                              `json:"arg"`
		EstimatedDeliveryExercisePrices []*public_data.EstimatedDeliveryExercisePrice `json:"data"`
	}
	MarkPrice struct {
		Arg    *events.Argument         `json:"arg"`
		Prices []*public_data.MarkPrice `json:"data"`
	}
	MarkPriceCandlesticks struct {
		Arg    *events.Argument      `json:"arg"`
		Prices []*market.IndexCandle `json:"data"`
	}
	PriceLimit struct {
		Arg   *events.Argument          `json:"arg"`
		Limit []*public_data.LimitPrice `json:"data"`
	}
	OrderBook struct {
		Arg   *events.Argument      `json:"arg"`
		Books []*market.OrderBookWs `json:"data"`
	}
	OPTIONSummary struct {
		Arg     *events.Argument                `json:"arg"`
		Options []*public_data.OptionMarketData `json:"data"`
	}
	FundingRate struct {
		Arg   *events.Argument           `json:"arg"`
		Rates []*public_data.FundingRate `json:"data"`
	}
	IndexCandlesticks struct {
		Arg   *events.Argument      `json:"arg"`
		Rates []*market.IndexCandle `json:"data"`
	}
	IndexTickers struct {
		Arg     *events.Argument      `json:"arg"`
		Tickers []*market.IndexTicker `json:"data"`
	}
)
