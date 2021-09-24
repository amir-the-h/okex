package market

import (
	"github.com/amir-the-h/okex/models/market"
	"github.com/amir-the-h/okex/responses"
)

type (
	Ticker struct {
		responses.RestBasic
		Tickers []*market.Ticker `json:"data,omitempty"`
	}
	IndexTicker struct {
		responses.RestBasic
		IndexTickers []*market.IndexTicker `json:"data,omitempty"`
	}
	OrderBook struct {
		responses.RestBasic
		OrderBooks []*market.OrderBook `json:"data,omitempty"`
	}
	Candle struct {
		responses.RestBasic
		Candles []*market.Candle `json:"data,omitempty"`
	}
	IndexCandle struct {
		responses.RestBasic
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	CandleMarket struct {
		responses.RestBasic
		Candles []*market.IndexCandle `json:"data,omitempty"`
	}
	Trade struct {
		responses.RestBasic
		Trades []*market.Trade `json:"data,omitempty"`
	}
	TotalVolume24H struct {
		responses.RestBasic
		TotalVolume24Hs []*market.TotalVolume24H `json:"data,omitempty"`
	}
	IndexComponent struct {
		responses.RestBasic
		IndexComponents *market.IndexComponent `json:"data,omitempty"`
	}
)
