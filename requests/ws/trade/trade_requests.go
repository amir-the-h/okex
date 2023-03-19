package trade

import "github.com/amir-the-h/okex"

type (
	PlaceOrder struct {
		ID         string            `json:"-"`
		InstID     string            `json:"instId"`
		Ccy        string            `json:"ccy,omitempty"`
		ClOrdID    string            `json:"clOrdId,omitempty"`
		Tag        string            `json:"tag,omitempty"`
		ReduceOnly bool              `json:"reduceOnly,omitempty"`
		Sz         float64           `json:"sz,string"`
		Px         float64           `json:"px,omitempty,string"`
		TdMode     okex.TradeMode    `json:"tdMode"`
		Side       okex.OrderSide    `json:"side"`
		PosSide    okex.PositionSide `json:"posSide,omitempty"`
		OrdType    okex.OrderType    `json:"ordType"`
		TgtCcy     okex.QuantityType `json:"tgtCcy,omitempty"`
	}
	CancelOrder struct {
		ID      string `json:"-"`
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		ID        string  `json:"-"`
		InstID    string  `json:"instId"`
		OrdID     string  `json:"ordId,omitempty"`
		ClOrdID   string  `json:"clOrdId,omitempty"`
		ReqID     string  `json:"reqId,omitempty"`
		NewSz     int64   `json:"newSz,omitempty,string"`
		NewPx     float64 `json:"newPx,omitempty,string"`
		CxlOnFail bool    `json:"cxlOnFail,omitempty"`
	}
)
