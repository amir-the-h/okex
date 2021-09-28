package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		InstID     string            `json:"instId,string"`
		Ccy        string            `json:"ccy,string,omitempty"`
		ClOrdID    string            `json:"clOrdId,string,omitempty"`
		Tag        string            `json:"tag,string,omitempty"`
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
		InstID  string `json:"instId,string"`
		OrdID   string `json:"ordId,string,omitempty"`
		ClOrdID string `json:"clOrdId,string,omitempty"`
	}
	AmendOrder struct {
		InstID    string  `json:"instId,string"`
		OrdID     string  `json:"ordId,string,omitempty"`
		ClOrdID   string  `json:"clOrdId,string,omitempty"`
		ReqID     string  `json:"reqId,string,omitempty"`
		NewSz     float64 `json:"newSz,omitempty,string"`
		NewPx     float64 `json:"newPx,omitempty,string"`
		CxlOnFail bool    `json:"cxlOnFail,string,omitempty"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId,string"`
		Ccy     string            `json:"ccy,string,omitempty"`
		PosSide okex.PositionSide `json:"posSide,omitempty,string"`
		MgnMode okex.MarginMode   `json:"mgnMode,string"`
	}
	OrderDetails struct {
		InstID  string `json:"instId,string"`
		OrdID   string `json:"ordId,string,omitempty"`
		ClOrdID string `json:"clOrdId,string,omitempty"`
	}
	OrderList struct {
		Uly      string              `json:"uly,string,omitempty"`
		InstID   string              `json:"instId,string,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
		OrdType  okex.OrderType      `json:"ordType,omitempty,string"`
		State    okex.OrderState     `json:"state,omitempty,string"`
	}
	TransactionDetails struct {
		Uly      string              `json:"uly,string,omitempty"`
		InstID   string              `json:"instId,string,omitempty"`
		OrdID    string              `json:"ordId,string,omitempty"`
		After    float64             `json:"after,omitempty,string"`
		Before   float64             `json:"before,omitempty,string"`
		Limit    float64             `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
	}
)
