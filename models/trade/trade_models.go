package trade

import (
	"github.com/amir-the-h/okex"
)

type (
	PlaceOrder struct {
		ClOrdId string  `json:"clOrdId"`
		Tag     string  `json:"tag"`
		SMsg    string  `json:"sMsg"`
		SCode   int64   `json:"sCode,string"`
		OrdId   float64 `json:"ordId,string"`
	}
	CancelOrder struct {
		ClOrdId string  `json:"clOrdId"`
		SMsg    string  `json:"sMsg"`
		OrdId   float64 `json:"ordId,string"`
		SCode   float64 `json:"sCode,string"`
	}
	AmendOrder struct {
		ClOrdId string  `json:"clOrdId"`
		SMsg    string  `json:"sMsg"`
		OrdId   float64 `json:"ordId,string"`
		ReqId   float64 `json:"reqId,string"`
		SCode   float64 `json:"sCode,string"`
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
		Px          float64             `json:"px,string"`
		Sz          float64             `json:"sz,string"`
		Pnl         float64             `json:"pnl,string"`
		AccFillSz   float64             `json:"accFillSz,string"`
		FillPx      float64             `json:"fillPx,string"`
		TradeId     float64             `json:"tradeId,string"`
		FillSz      float64             `json:"fillSz,string"`
		FillTime    float64             `json:"fillTime,string"`
		AvgPx       float64             `json:"avgPx,string"`
		Lever       float64             `json:"lever,string"`
		TpTriggerPx float64             `json:"tpTriggerPx,string"`
		TpOrdPx     float64             `json:"tpOrdPx,string"`
		SlTriggerPx float64             `json:"slTriggerPx,string"`
		SlOrdPx     float64             `json:"slOrdPx,string"`
		FeeCcy      float64             `json:"feeCcy,string"`
		Fee         float64             `json:"fee,string"`
		RebateCcy   float64             `json:"rebateCcy,string"`
		Rebate      float64             `json:"rebate,string"`
		Category    float64             `json:"category,string"`
		OrdId       float64             `json:"ordId,string"`
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
		OrdId    float64             `json:"ordId,string"`
		BillId   float64             `json:"billId,string"`
		Tag      float64             `json:"tag,string"`
		FillPx   float64             `json:"fillPx,string"`
		FillSz   float64             `json:"fillSz,string"`
		FeeCcy   float64             `json:"feeCcy,string"`
		Fee      float64             `json:"fee,string"`
		InstType okex.InstrumentType `json:"instType"`
		Side     okex.OrderSide      `json:"side"`
		PosSide  okex.PositionSide   `json:"posSide"`
		ExecType okex.OrderFlowType  `json:"execType"`
		Ts       okex.JsonTime       `json:"ts"`
	}
)
