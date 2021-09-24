package account

import (
	"github.com/amir-the-h/okex"
)

type (
	Balance struct {
		TotalEq     float64           `json:"totalEq,string"`
		IsoEq       float64           `json:"isoEq,string"`
		AdjEq       float64           `json:"adjEq,omitempty"`
		OrdFroz     float64           `json:"ordFroz,omitempty"`
		Imr         float64           `json:"imr,omitempty"`
		Mmr         float64           `json:"mmr,omitempty"`
		MgnRatio    float64           `json:"mgnRatio,omitempty"`
		NotionalUsd float64           `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okex.JsonTime     `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string        `json:"ccy"`
		Eq            float64       `json:"eq,string"`
		CashBal       float64       `json:"cashBal,string"`
		IsoEq         float64       `json:"isoEq,omitempty,string"`
		AvailEq       float64       `json:"availEq,omitempty,string"`
		DisEq         float64       `json:"disEq,string"`
		AvailBal      float64       `json:"availBal,string"`
		FrozenBal     float64       `json:"frozenBal,string"`
		OrdFrozen     float64       `json:"ordFrozen,string"`
		Liab          float64       `json:"liab,omitempty,string"`
		Upl           float64       `json:"upl,omitempty,string"`
		UplLib        float64       `json:"uplLib,omitempty,string"`
		CrossLiab     float64       `json:"crossLiab,omitempty,string"`
		IsoLiab       float64       `json:"isoLiab,omitempty,string"`
		MgnRatio      float64       `json:"mgnRatio,omitempty,string"`
		Interest      float64       `json:"interest,omitempty,string"`
		Twap          float64       `json:"twap,omitempty,string"`
		MaxLoan       float64       `json:"maxLoan,omitempty,string"`
		EqUsd         float64       `json:"eqUsd,string"`
		NotionalLever float64       `json:"notionalLever,omitempty,string"`
		StgyEq        float64       `json:"stgyEq,string"`
		IsoUpl        float64       `json:"isoUpl,omitempty,string"`
		UTime         okex.JsonTime `json:"uTime"`
	}
	Position struct {
		InstId      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		LiabCcy     string              `json:"liabCcy,omitempty"`
		OptVal      string              `json:"optVal,omitempty"`
		Ccy         string              `json:"ccy"`
		PosId       float64             `json:"posId,string"`
		Pos         float64             `json:"pos,string"`
		AvailPos    float64             `json:"availPos,omitempty,string"`
		AvgPx       float64             `json:"avgPx,string"`
		Upl         float64             `json:"upl,string"`
		UplRatio    float64             `json:"uplRatio,string"`
		Lever       float64             `json:"lever,omitempty,string"`
		LiqPx       float64             `json:"liqPx,omitempty,string"`
		Imr         float64             `json:"imr,omitempty,string"`
		Margin      float64             `json:"margin,omitempty,string"`
		MgnRatio    float64             `json:"mgnRatio,string"`
		Mmr         float64             `json:"mmr,string"`
		Liab        float64             `json:"liab,omitempty,string"`
		Interest    float64             `json:"interest,string"`
		TradeId     float64             `json:"tradeId,string"`
		NotionalUsd float64             `json:"notionalUsd,string"`
		ADL         float64             `json:"adl,string"`
		Last        float64             `json:"last,string"`
		DeltaBS     float64             `json:"deltaBS,string"`
		DeltaPA     float64             `json:"deltaPA,string"`
		GammaBS     float64             `json:"gammaBS,string"`
		GammaPA     float64             `json:"gammaPA,string"`
		ThetaBS     float64             `json:"thetaBS,string"`
		ThetaPA     float64             `json:"thetaPA,string"`
		VegaBS      float64             `json:"vegaBS,string"`
		VegaPA      float64             `json:"vegaPA,string"`
		PosSide     okex.PositionSide   `json:"posSide"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
		InstType    okex.InstrumentType `json:"instType"`
		CTime       okex.JsonTime       `json:"cTime"`
		UTime       okex.JsonTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType"`
		PTime     okex.JsonTime     `json:"pTime"`
		UTime     okex.JsonTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   float64                              `json:"adjEq,omitempty,string"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		Ts      okex.JsonTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string  `json:"ccy"`
		Eq    float64 `json:"eq,string"`
		DisEq float64 `json:"disEq,string"`
	}
	PositionAndAccountRiskPositionData struct {
		InstId      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy float64             `json:"notionalCcy,string"`
		Pos         float64             `json:"pos,string"`
		NotionalUsd float64             `json:"notionalUsd,string"`
		PosSide     okex.PositionSide   `json:"posSide"`
		InstType    okex.InstrumentType `json:"instType"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstId    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillId    float64             `json:"billId,string"`
		BalChg    float64             `json:"balChg,string"`
		PosBalChg float64             `json:"posBalChg,string"`
		Bal       float64             `json:"bal,string"`
		PosBal    float64             `json:"posBal,string"`
		Sz        float64             `json:"sz,string"`
		Pnl       float64             `json:"pnl,string"`
		Fee       float64             `json:"fee,string"`
		OrdId     float64             `json:"ordId,string"`
		From      okex.AccountType    `json:"from,string"`
		To        okex.AccountType    `json:"to,string"`
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		Type      okex.BillType       `json:"type,string"`
		SubType   okex.BillSubType    `json:"subType,string"`
		Ts        okex.JsonTime       `json:"ts"`
	}
	Config struct {
		Level      string         `json:"level"`
		LevelTmp   string         `json:"levelTmp"`
		Uid        float64        `json:"uid,string"`
		AcctLv     float64        `json:"acctLv,string"`
		AutoLoan   bool           `json:"autoLoan"`
		GreeksType okex.GreekType `json:"greeksType"`
		PosMode    PositionMode   `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstId  string            `json:"instId"`
		Lever   int64             `json:"lever,string"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstId  string  `json:"instId"`
		Ccy     string  `json:"ccy"`
		MaxBuy  float64 `json:"maxBuy,string"`
		MaxSell float64 `json:"maxSell,string"`
	}
	MaxAvailableTradeAmount struct {
		InstId    string  `json:"instId"`
		AvailBuy  float64 `json:"availBuy,string"`
		AvailSell float64 `json:"availSell,string"`
	}
	MarginBalanceAmount struct {
		InstId  string            `json:"instId"`
		Amt     float64           `json:"amt,string"`
		PosSide okex.PositionSide `json:"posSide"`
		Type    okex.CountAction  `json:"type"`
	}
	Loan struct {
		InstId  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy"`
		Ccy     string          `json:"ccy"`
		MaxLoan float64         `json:"maxLoan,string"`
		MgnMode okex.MarginMode `json:"mgnMode"`
		Side    okex.OrderSide  `json:"side"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    float64             `json:"taker,string"`
		Maker    float64             `json:"maker,string"`
		Delivery float64             `json:"delivery,string"`
		Exercise float64             `json:"exercise,string"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType"`
		Ts       okex.JsonTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstId       string          `json:"instId"`
		Ccy          string          `json:"ccy"`
		Interest     float64         `json:"interest,string"`
		InterestRate float64         `json:"interestRate,string"`
		Liab         float64         `json:"liab,string"`
		MgnMode      okex.MarginMode `json:"mgnMode"`
		Ts           okex.JsonTime   `json:"ts"`
	}
	InterestRate struct {
		Ccy          string  `json:"ccy"`
		InterestRate float64 `json:"interestRate,string"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string  `json:"ccy"`
		MaxWd float64 `json:"maxWd,string"`
	}
)
