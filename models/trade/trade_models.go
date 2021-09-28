package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		ClOrdId string           `json:"clOrdId"`
		Tag     string           `json:"tag"`
		SMsg    string           `json:"sMsg"`
		SCode   okex.JSONInt64   `json:"sCode"`
		OrdId   okex.JSONFloat64 `json:"ordId"`
	}
	CancelOrder struct {
		ClOrdId string           `json:"clOrdId"`
		SMsg    string           `json:"sMsg"`
		OrdId   okex.JSONFloat64 `json:"ordId"`
		SCode   okex.JSONFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		ClOrdId string           `json:"clOrdId"`
		SMsg    string           `json:"sMsg"`
		OrdId   okex.JSONFloat64 `json:"ordId"`
		ReqId   okex.JSONFloat64 `json:"reqId"`
		SCode   okex.JSONFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string              `json:"instId"`
		Ccy         string              `json:"ccy"`
		ClOrdId     string              `json:"clOrdId"`
		Tag         string              `json:"tag"`
		Px          okex.JSONFloat64    `json:"px"`
		Sz          okex.JSONFloat64    `json:"sz"`
		Pnl         okex.JSONFloat64    `json:"pnl"`
		AccFillSz   okex.JSONFloat64    `json:"accFillSz"`
		FillPx      okex.JSONFloat64    `json:"fillPx"`
		TradeId     okex.JSONFloat64    `json:"tradeId"`
		FillSz      okex.JSONFloat64    `json:"fillSz"`
		FillTime    okex.JSONFloat64    `json:"fillTime"`
		AvgPx       okex.JSONFloat64    `json:"avgPx"`
		Lever       okex.JSONFloat64    `json:"lever"`
		TpTriggerPx okex.JSONFloat64    `json:"tpTriggerPx"`
		TpOrdPx     okex.JSONFloat64    `json:"tpOrdPx"`
		SlTriggerPx okex.JSONFloat64    `json:"slTriggerPx"`
		SlOrdPx     okex.JSONFloat64    `json:"slOrdPx"`
		FeeCcy      okex.JSONFloat64    `json:"feeCcy"`
		Fee         okex.JSONFloat64    `json:"fee"`
		RebateCcy   okex.JSONFloat64    `json:"rebateCcy"`
		Rebate      okex.JSONFloat64    `json:"rebate"`
		Category    okex.JSONFloat64    `json:"category"`
		OrdId       okex.JSONFloat64    `json:"ordId"`
		State       okex.OrderState     `json:"state"`
		TdMode      okex.TradeMode      `json:"tdMode"`
		PosSide     okex.PositionSide   `json:"posSide"`
		Side        okex.OrderSide      `json:"side"`
		OrdType     okex.OrderType      `json:"ordType"`
		InstType    okex.InstrumentType `json:"instType"`
		TgtCcy      okex.QuantityType   `json:"tgtCcy"`
		UTime       okex.JSONTime       `json:"uTime"`
		CTime       okex.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string              `json:"instId"`
		TradeId  string              `json:"tradeId"`
		ClOrdId  string              `json:"clOrdId"`
		OrdId    okex.JSONFloat64    `json:"ordId"`
		BillId   okex.JSONFloat64    `json:"billId"`
		Tag      okex.JSONFloat64    `json:"tag"`
		FillPx   okex.JSONFloat64    `json:"fillPx"`
		FillSz   okex.JSONFloat64    `json:"fillSz"`
		FeeCcy   okex.JSONFloat64    `json:"feeCcy"`
		Fee      okex.JSONFloat64    `json:"fee"`
		InstType okex.InstrumentType `json:"instType"`
		Side     okex.OrderSide      `json:"side"`
		PosSide  okex.PositionSide   `json:"posSide"`
		ExecType okex.OrderFlowType  `json:"execType"`
		Ts       okex.JSONTime       `json:"ts"`
	}
)
