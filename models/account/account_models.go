package account

import (
	"github.com/amir-the-h/okex"
)

type (
	Balance struct {
		TotalEq     okex.JsonFloat64  `json:"totalEq"`
		IsoEq       okex.JsonFloat64  `json:"isoEq"`
		AdjEq       okex.JsonFloat64  `json:"adjEq,omitempty"`
		OrdFroz     okex.JsonFloat64  `json:"ordFroz,omitempty"`
		Imr         okex.JsonFloat64  `json:"imr,omitempty"`
		Mmr         okex.JsonFloat64  `json:"mmr,omitempty"`
		MgnRatio    okex.JsonFloat64  `json:"mgnRatio,omitempty"`
		NotionalUsd okex.JsonFloat64  `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okex.JsonTime     `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string           `json:"ccy"`
		Eq            okex.JsonFloat64 `json:"eq"`
		CashBal       okex.JsonFloat64 `json:"cashBal"`
		IsoEq         okex.JsonFloat64 `json:"isoEq,omitempty"`
		AvailEq       okex.JsonFloat64 `json:"availEq,omitempty"`
		DisEq         okex.JsonFloat64 `json:"disEq"`
		AvailBal      okex.JsonFloat64 `json:"availBal"`
		FrozenBal     okex.JsonFloat64 `json:"frozenBal"`
		OrdFrozen     okex.JsonFloat64 `json:"ordFrozen"`
		Liab          okex.JsonFloat64 `json:"liab,omitempty"`
		Upl           okex.JsonFloat64 `json:"upl,omitempty"`
		UplLib        okex.JsonFloat64 `json:"uplLib,omitempty"`
		CrossLiab     okex.JsonFloat64 `json:"crossLiab,omitempty"`
		IsoLiab       okex.JsonFloat64 `json:"isoLiab,omitempty"`
		MgnRatio      okex.JsonFloat64 `json:"mgnRatio,omitempty"`
		Interest      okex.JsonFloat64 `json:"interest,omitempty"`
		Twap          okex.JsonFloat64 `json:"twap,omitempty"`
		MaxLoan       okex.JsonFloat64 `json:"maxLoan,omitempty"`
		EqUsd         okex.JsonFloat64 `json:"eqUsd"`
		NotionalLever okex.JsonFloat64 `json:"notionalLever,omitempty"`
		StgyEq        okex.JsonFloat64 `json:"stgyEq"`
		IsoUpl        okex.JsonFloat64 `json:"isoUpl,omitempty"`
		UTime         okex.JsonTime    `json:"uTime"`
	}
	Position struct {
		InstId      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		LiabCcy     string              `json:"liabCcy,omitempty"`
		OptVal      string              `json:"optVal,omitempty"`
		Ccy         string              `json:"ccy"`
		PosId       okex.JsonFloat64    `json:"posId"`
		Pos         okex.JsonFloat64    `json:"pos"`
		AvailPos    okex.JsonFloat64    `json:"availPos,omitempty"`
		AvgPx       okex.JsonFloat64    `json:"avgPx"`
		Upl         okex.JsonFloat64    `json:"upl"`
		UplRatio    okex.JsonFloat64    `json:"uplRatio"`
		Lever       okex.JsonFloat64    `json:"lever"`
		LiqPx       okex.JsonFloat64    `json:"liqPx,omitempty"`
		Imr         okex.JsonFloat64    `json:"imr,omitempty"`
		Margin      okex.JsonFloat64    `json:"margin,omitempty"`
		MgnRatio    okex.JsonFloat64    `json:"mgnRatio"`
		Mmr         okex.JsonFloat64    `json:"mmr"`
		Liab        okex.JsonFloat64    `json:"liab,omitempty"`
		Interest    okex.JsonFloat64    `json:"interest"`
		TradeId     okex.JsonFloat64    `json:"tradeId"`
		NotionalUsd okex.JsonFloat64    `json:"notionalUsd"`
		ADL         okex.JsonFloat64    `json:"adl"`
		Last        okex.JsonFloat64    `json:"last"`
		DeltaBS     okex.JsonFloat64    `json:"deltaBS"`
		DeltaPA     okex.JsonFloat64    `json:"deltaPA"`
		GammaBS     okex.JsonFloat64    `json:"gammaBS"`
		GammaPA     okex.JsonFloat64    `json:"gammaPA"`
		ThetaBS     okex.JsonFloat64    `json:"thetaBS"`
		ThetaPA     okex.JsonFloat64    `json:"thetaPA"`
		VegaBS      okex.JsonFloat64    `json:"vegaBS"`
		VegaPA      okex.JsonFloat64    `json:"vegaPA"`
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
		AdjEq   okex.JsonFloat64                     `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		Ts      okex.JsonTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string           `json:"ccy"`
		Eq    okex.JsonFloat64 `json:"eq"`
		DisEq okex.JsonFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstId      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy okex.JsonFloat64    `json:"notionalCcy"`
		Pos         okex.JsonFloat64    `json:"pos"`
		NotionalUsd okex.JsonFloat64    `json:"notionalUsd"`
		PosSide     okex.PositionSide   `json:"posSide"`
		InstType    okex.InstrumentType `json:"instType"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstId    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillId    okex.JsonFloat64    `json:"billId"`
		BalChg    okex.JsonFloat64    `json:"balChg"`
		PosBalChg okex.JsonFloat64    `json:"posBalChg"`
		Bal       okex.JsonFloat64    `json:"bal"`
		PosBal    okex.JsonFloat64    `json:"posBal"`
		Sz        okex.JsonFloat64    `json:"sz"`
		Pnl       okex.JsonFloat64    `json:"pnl"`
		Fee       okex.JsonFloat64    `json:"fee"`
		OrdId     okex.JsonFloat64    `json:"ordId"`
		From      okex.AccountType    `json:"from"`
		To        okex.AccountType    `json:"to"`
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		Type      okex.BillType       `json:"type"`
		SubType   okex.BillSubType    `json:"subType"`
		Ts        okex.JsonTime       `json:"ts"`
	}
	Config struct {
		Level      string           `json:"level"`
		LevelTmp   string           `json:"levelTmp"`
		Uid        okex.JsonFloat64 `json:"uid"`
		AcctLv     okex.JsonFloat64 `json:"acctLv"`
		AutoLoan   bool             `json:"autoLoan"`
		GreeksType okex.GreekType   `json:"greeksType"`
		PosMode    PositionMode     `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstId  string            `json:"instId"`
		Lever   okex.JsonFloat64  `json:"lever"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstId  string           `json:"instId"`
		Ccy     string           `json:"ccy"`
		MaxBuy  okex.JsonFloat64 `json:"maxBuy"`
		MaxSell okex.JsonFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstId    string           `json:"instId"`
		AvailBuy  okex.JsonFloat64 `json:"availBuy"`
		AvailSell okex.JsonFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstId  string            `json:"instId"`
		Amt     okex.JsonFloat64  `json:"amt"`
		PosSide okex.PositionSide `json:"posSide"`
		Type    okex.CountAction  `json:"type"`
	}
	Loan struct {
		InstId  string           `json:"instId"`
		MgnCcy  string           `json:"mgnCcy"`
		Ccy     string           `json:"ccy"`
		MaxLoan okex.JsonFloat64 `json:"maxLoan"`
		MgnMode okex.MarginMode  `json:"mgnMode"`
		Side    okex.OrderSide   `json:"side"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    okex.JsonFloat64    `json:"taker"`
		Maker    okex.JsonFloat64    `json:"maker"`
		Delivery okex.JsonFloat64    `json:"delivery,omitempty"`
		Exercise okex.JsonFloat64    `json:"exercise,omitempty"`
		Category okex.FeeCategory    `json:"category"`
		InstType okex.InstrumentType `json:"instType"`
		Ts       okex.JsonTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstId       string           `json:"instId"`
		Ccy          string           `json:"ccy"`
		Interest     okex.JsonFloat64 `json:"interest"`
		InterestRate okex.JsonFloat64 `json:"interestRate"`
		Liab         okex.JsonFloat64 `json:"liab"`
		MgnMode      okex.MarginMode  `json:"mgnMode"`
		Ts           okex.JsonTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string           `json:"ccy"`
		InterestRate okex.JsonFloat64 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string           `json:"ccy"`
		MaxWd okex.JsonFloat64 `json:"maxWd"`
	}
)
