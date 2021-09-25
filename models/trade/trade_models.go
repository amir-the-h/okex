package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		ClOrdId string           `json:"clOrdId"`
		Tag     string           `json:"tag"`
		SMsg    string           `json:"sMsg"`
		SCode   okex.JsonInt64   `json:"sCode"`
		OrdId   okex.JsonFloat64 `json:"ordId"`
	}
	CancelOrder struct {
		ClOrdId string           `json:"clOrdId"`
		SMsg    string           `json:"sMsg"`
		OrdId   okex.JsonFloat64 `json:"ordId"`
		SCode   okex.JsonFloat64 `json:"sCode"`
	}
	AmendOrder struct {
		ClOrdId string           `json:"clOrdId"`
		SMsg    string           `json:"sMsg"`
		OrdId   okex.JsonFloat64 `json:"ordId"`
		ReqId   okex.JsonFloat64 `json:"reqId"`
		SCode   okex.JsonFloat64 `json:"sCode"`
	}
	ClosePosition struct {
		InstId  string            `json:"instId"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	Order struct {
		InstId      string              `json:"instId"`
		Ccy         string              `json:"ccy"`
		ClOrdId     string              `json:"clOrdId"`
		Tag         string              `json:"tag"`
		Px          okex.JsonFloat64    `json:"px"`
		Sz          okex.JsonFloat64    `json:"sz"`
		Pnl         okex.JsonFloat64    `json:"pnl"`
		AccFillSz   okex.JsonFloat64    `json:"accFillSz"`
		FillPx      okex.JsonFloat64    `json:"fillPx"`
		TradeId     okex.JsonFloat64    `json:"tradeId"`
		FillSz      okex.JsonFloat64    `json:"fillSz"`
		FillTime    okex.JsonFloat64    `json:"fillTime"`
		AvgPx       okex.JsonFloat64    `json:"avgPx"`
		Lever       okex.JsonFloat64    `json:"lever"`
		TpTriggerPx okex.JsonFloat64    `json:"tpTriggerPx"`
		TpOrdPx     okex.JsonFloat64    `json:"tpOrdPx"`
		SlTriggerPx okex.JsonFloat64    `json:"slTriggerPx"`
		SlOrdPx     okex.JsonFloat64    `json:"slOrdPx"`
		FeeCcy      okex.JsonFloat64    `json:"feeCcy"`
		Fee         okex.JsonFloat64    `json:"fee"`
		RebateCcy   okex.JsonFloat64    `json:"rebateCcy"`
		Rebate      okex.JsonFloat64    `json:"rebate"`
		Category    okex.JsonFloat64    `json:"category"`
		OrdId       okex.JsonFloat64    `json:"ordId"`
		State       okex.OrderState     `json:"state"`
		TdMode      okex.TradeMode      `json:"tdMode"`
		PosSide     okex.PositionSide   `json:"posSide"`
		Side        okex.OrderSide      `json:"side"`
		OrdType     okex.OrderType      `json:"ordType"`
		InstType    okex.InstrumentType `json:"instType"`
		TgtCcy      okex.QuantityType   `json:"tgtCcy"`
		UTime       okex.JsonTime       `json:"uTime"`
		CTime       okex.JsonTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstId   string              `json:"instId"`
		TradeId  string              `json:"tradeId"`
		ClOrdId  string              `json:"clOrdId"`
		OrdId    okex.JsonFloat64    `json:"ordId"`
		BillId   okex.JsonFloat64    `json:"billId"`
		Tag      okex.JsonFloat64    `json:"tag"`
		FillPx   okex.JsonFloat64    `json:"fillPx"`
		FillSz   okex.JsonFloat64    `json:"fillSz"`
		FeeCcy   okex.JsonFloat64    `json:"feeCcy"`
		Fee      okex.JsonFloat64    `json:"fee"`
		InstType okex.InstrumentType `json:"instType"`
		Side     okex.OrderSide      `json:"side"`
		PosSide  okex.PositionSide   `json:"posSide"`
		ExecType okex.OrderFlowType  `json:"execType"`
		Ts       okex.JsonTime       `json:"ts"`
	}
)
