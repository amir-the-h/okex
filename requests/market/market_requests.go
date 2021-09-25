package market

import "github.com/amir-the-h/okex"

type (
	GetTickers struct {
		Uly      string              `json:"uly,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetIndexTickers struct {
		InstId   string `json:"instId,omitempty"`
		QuoteCcy string `json:"quoteCcy,omitempty"`
	}
	GetOrderBook struct {
		InstId string `json:"instId"`
		Sz     int    `json:"sz,omitempty,string"`
	}
	GetCandlesticks struct {
		InstId string       `json:"instId"`
		After  int64        `json:"after,omitempty,string"`
		Before int64        `json:"before,omitempty,string"`
		Limit  int64        `json:"limit,omitempty,string"`
		Bar    okex.BarSize `json:"bar,omitempty,string"`
	}
	GetTrades struct {
		InstId string `json:"instId"`
		Limit  int64  `json:"limit,omitempty,string"`
	}
	GetIndexComponents struct {
		Index string `json:"index"`
	}
)
