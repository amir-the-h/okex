package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		InstID     string            `json:"instId"`
		Ccy        string            `json:"ccy,omitempty"`
		ClOrdID    string            `json:"clOrdId,omitempty"`
		Tag        string            `json:"tag,omitempty"`
		ReduceOnly bool              `json:"reduceOnly,omitempty"`
		Sz         float64           `json:"sz,string"`
		Px         float64           `json:"px,omitempty,string"`
		TdMode     okex.TradeMode    `json:"tdMode,string"`
		Side       okex.OrderSide    `json:"side,string"`
		PosSide    okex.PositionSide `json:"posSide,string"`
		OrdType    okex.OrderType    `json:"ordType,string"`
		TgtCcy     okex.QuantityType `json:"tgtCcy,string,omitempty"`
	}
	CancelOrder struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		InstID    string  `json:"instId"`
		OrdID     string  `json:"ordId,omitempty"`
		ClOrdID   string  `json:"clOrdId,omitempty"`
		ReqID     string  `json:"reqId,omitempty"`
		NewSz     float64 `json:"newSz,omitempty,string"`
		NewPx     float64 `json:"newPx,omitempty,string"`
		CxlOnFail bool    `json:"cxlOnFail,string,omitempty"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		Ccy     string            `json:"ccy,omitempty"`
		PosSide okex.PositionSide `json:"posSide,omitempty,string"`
		MgnMode okex.MarginMode   `json:"mgnMode,string"`
	}
	OrderDetails struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	OrderList struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
		OrdType  okex.OrderType      `json:"ordType,omitempty,string"`
		State    okex.OrderState     `json:"state,omitempty,string"`
	}
	TransactionDetails struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		OrdID    string              `json:"ordId,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
	}
)
