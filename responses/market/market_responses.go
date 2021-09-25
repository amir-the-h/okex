package market

import (
	"github.com/amir-the-h/okex/models/market"
	"github.com/amir-the-h/okex/responses"
)

type (
	Ticker struct {
		responses.Basic
		Tickers []*market.Ticker `json:"data,omitempty"`
	}
	IndexTicker struct {
		responses.Basic
		IndexTickers []*market.IndexTicker `json:"data,omitempty"`
	}
	OrderBook struct {
		responses.Basic
		OrderBooks []*market.OrderBook `json:"data,omitempty"`
	}
	Candle struct {
		responses.Basic
		Candles []*market.Candle `json:"data,omitempty"`
	}
	IndexCandle struct {
		responses.Basic
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	CandleMarket struct {
		responses.Basic
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	Trade struct {
		responses.Basic
		Trades []*market.Trade `json:"data,omitempty"`
	}
	TotalVolume24H struct {
		responses.Basic
		TotalVolume24Hs []*market.TotalVolume24H `json:"data,omitempty"`
	}
	IndexComponent struct {
		responses.Basic
		IndexComponents *market.IndexComponent `json:"data,omitempty"`
	}
)
