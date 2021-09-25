package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		InstId     string            `json:"instId"`
		Ccy        string            `json:"ccy,omitempty"`
		ClOrdId    string            `json:"clOrdId,omitempty"`
		Tag        string            `json:"tag,omitempty"`
		ReduceOnly bool              `json:"reduceOnly,omitempty"`
		Sz         float64           `json:"sz,string"`
		Px         float64           `json:"px,omitempty,string"`
		TdMode     okex.TradeMode    `json:"tdMode"`
		Side       okex.OrderSide    `json:"side"`
		PosSide    okex.PositionSide `json:"posSide"`
		OrdType    okex.OrderType    `json:"ordType"`
		TgtCcy     okex.QuantityType `json:"tgtCcy,omitempty"`
	}
	CancelOrder struct {
		InstId  string `json:"instId"`
		OrdId   string `json:"ordId,omitempty"`
		ClOrdId string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		InstId    string  `json:"instId"`
		OrdId     string  `json:"ordId,omitempty"`
		ClOrdId   string  `json:"clOrdId,omitempty"`
		ReqId     string  `json:"reqId,omitempty"`
		NewSz     float64 `json:"newSz,omitempty,string"`
		NewPx     float64 `json:"newPx,omitempty,string"`
		CxlOnFail bool    `json:"cxlOnFail,omitempty"`
	}
	ClosePosition struct {
		InstId  string            `json:"instId"`
		Ccy     string            `json:"ccy,omitempty"`
		PosSide okex.PositionSide `json:"posSide,omitempty,string"`
		MgnMode okex.MarginMode   `json:"mgnMode,string"`
	}
	OrderDetails struct {
		InstId  string `json:"instId"`
		OrdId   string `json:"ordId,omitempty"`
		ClOrdId string `json:"clOrdId,omitempty"`
	}
	OrderList struct {
		Uly      string              `json:"uly,omitempty"`
		InstId   string              `json:"instId,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
		OrdType  okex.OrderType      `json:"ordType,omitempty,string"`
		State    okex.OrderState     `json:"state,omitempty,string"`
	}
	TransactionDetails struct {
		Uly      string              `json:"uly,omitempty"`
		InstId   string              `json:"instId,omitempty"`
		OrdId    string              `json:"ordId,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
	}
)
