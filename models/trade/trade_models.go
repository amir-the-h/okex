package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		ClOrdID string           `json:"clOrdId,string"`
		Tag     string           `json:"tag,string"`
		SMsg    string           `json:"sMsg,string"`
		SCode   okex.JSONInt64   `json:"sCode,string"`
		OrdID   okex.JSONFloat64 `json:"ordId,string"`
	}
	CancelOrder struct {
		ClOrdID string           `json:"clOrdId,string"`
		SMsg    string           `json:"sMsg,string"`
		OrdId   okex.JSONFloat64 `json:"ordId,string"`
		SCode   okex.JSONFloat64 `json:"sCode,string"`
	}
	AmendOrder struct {
		ClOrdID string           `json:"clOrdId,string"`
		SMsg    string           `json:"sMsg,string"`
		OrdID   okex.JSONFloat64 `json:"ordId,string"`
		ReqID   okex.JSONFloat64 `json:"reqId,string"`
		SCode   okex.JSONFloat64 `json:"sCode,string"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId,string"`
		PosSide okex.PositionSide `json:"posSide,string"`
	}
	Order struct {
		InstID      string              `json:"instId,string"`
		Ccy         string              `json:"ccy,string"`
		ClOrdID     string              `json:"clOrdId,string"`
		Tag         string              `json:"tag,string"`
		Px          okex.JSONFloat64    `json:"px,string"`
		Sz          okex.JSONFloat64    `json:"sz,string"`
		Pnl         okex.JSONFloat64    `json:"pnl,string"`
		AccFillSz   okex.JSONFloat64    `json:"accFillSz,string"`
		FillPx      okex.JSONFloat64    `json:"fillPx,string"`
		TradeId     okex.JSONFloat64    `json:"tradeId,string"`
		FillSz      okex.JSONFloat64    `json:"fillSz,string"`
		FillTime    okex.JSONFloat64    `json:"fillTime,string"`
		AvgPx       okex.JSONFloat64    `json:"avgPx,string"`
		Lever       okex.JSONFloat64    `json:"lever,string"`
		TpTriggerPx okex.JSONFloat64    `json:"tpTriggerPx,string"`
		TpOrdPx     okex.JSONFloat64    `json:"tpOrdPx,string"`
		SlTriggerPx okex.JSONFloat64    `json:"slTriggerPx,string"`
		SlOrdPx     okex.JSONFloat64    `json:"slOrdPx,string"`
		FeeCcy      okex.JSONFloat64    `json:"feeCcy,string"`
		Fee         okex.JSONFloat64    `json:"fee,string"`
		RebateCcy   okex.JSONFloat64    `json:"rebateCcy,string"`
		Rebate      okex.JSONFloat64    `json:"rebate,string"`
		Category    okex.JSONFloat64    `json:"category,string"`
		OrdId       okex.JSONFloat64    `json:"ordId,string"`
		State       okex.OrderState     `json:"state,string"`
		TdMode      okex.TradeMode      `json:"tdMode,string"`
		PosSide     okex.PositionSide   `json:"posSide,string"`
		Side        okex.OrderSide      `json:"side,string"`
		OrdType     okex.OrderType      `json:"ordType,string"`
		InstType    okex.InstrumentType `json:"instType,string"`
		TgtCcy      okex.QuantityType   `json:"tgtCcy,string"`
		UTime       okex.JSONTime       `json:"uTime"`
		CTime       okex.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string              `json:"instId,string"`
		TradeID  string              `json:"tradeId,string"`
		ClOrdID  string              `json:"clOrdId,string"`
		OrdId    okex.JSONFloat64    `json:"ordId,string"`
		BillID   okex.JSONFloat64    `json:"billId,string"`
		Tag      okex.JSONFloat64    `json:"tag,string"`
		FillPx   okex.JSONFloat64    `json:"fillPx,string"`
		FillSz   okex.JSONFloat64    `json:"fillSz,string"`
		FeeCcy   okex.JSONFloat64    `json:"feeCcy,string"`
		Fee      okex.JSONFloat64    `json:"fee,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		Side     okex.OrderSide      `json:"side,string"`
		PosSide  okex.PositionSide   `json:"posSide,string"`
		ExecType okex.OrderFlowType  `json:"execType,string"`
		TS       okex.JSONTime       `json:"ts"`
	}
)
