package market

import "github.com/amir-the-h/okex"

type (
	GetTickers struct {
		Uly      string              `json:"uly,string,omitempty"`
		InstType okex.InstrumentType `json:"instType,string"`
	}
	GetIndexTickers struct {
		InstID   string `json:"instId,string,omitempty"`
		QuoteCcy string `json:"quoteCcy,string,omitempty"`
	}
	GetOrderBook struct {
		InstID string `json:"instId,string"`
		Sz     int    `json:"sz,omitempty,string"`
	}
	GetCandlesticks struct {
		InstID string       `json:"instId,string"`
		After  int64        `json:"after,omitempty,string"`
		Before int64        `json:"before,omitempty,string"`
		Limit  int64        `json:"limit,omitempty,string"`
		Bar    okex.BarSize `json:"bar,omitempty,string"`
	}
	GetTrades struct {
		InstID string `json:"instId,string"`
		Limit  int64  `json:"limit,omitempty,string"`
	}
	GetIndexComponents struct {
		Index string `json:"index,string"`
	}
)
