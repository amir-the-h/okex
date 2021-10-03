package account

import (
	"github.com/amir-the-h/okex"
)

type (
	Balance struct {
		TotalEq     okex.JSONFloat64  `json:"totalEq"`
		IsoEq       okex.JSONFloat64  `json:"isoEq"`
		AdjEq       okex.JSONFloat64  `json:"adjEq,omitempty"`
		OrdFroz     okex.JSONFloat64  `json:"ordFroz,omitempty"`
		Imr         okex.JSONFloat64  `json:"imr,omitempty"`
		Mmr         okex.JSONFloat64  `json:"mmr,omitempty"`
		MgnRatio    okex.JSONFloat64  `json:"mgnRatio,omitempty"`
		NotionalUsd okex.JSONFloat64  `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okex.JSONTime     `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string           `json:"ccy"`
		Eq            okex.JSONFloat64 `json:"eq"`
		CashBal       okex.JSONFloat64 `json:"cashBal"`
		IsoEq         okex.JSONFloat64 `json:"isoEq,omitempty"`
		AvailEq       okex.JSONFloat64 `json:"availEq,omitempty"`
		DisEq         okex.JSONFloat64 `json:"disEq"`
		AvailBal      okex.JSONFloat64 `json:"availBal"`
		FrozenBal     okex.JSONFloat64 `json:"frozenBal"`
		OrdFrozen     okex.JSONFloat64 `json:"ordFrozen"`
		Liab          okex.JSONFloat64 `json:"liab,omitempty"`
		Upl           okex.JSONFloat64 `json:"upl,omitempty"`
		UplLib        okex.JSONFloat64 `json:"uplLib,omitempty"`
		CrossLiab     okex.JSONFloat64 `json:"crossLiab,omitempty"`
		IsoLiab       okex.JSONFloat64 `json:"isoLiab,omitempty"`
		MgnRatio      okex.JSONFloat64 `json:"mgnRatio,omitempty"`
		Interest      okex.JSONFloat64 `json:"interest,omitempty"`
		Twap          okex.JSONFloat64 `json:"twap,omitempty"`
		MaxLoan       okex.JSONFloat64 `json:"maxLoan,omitempty"`
		EqUsd         okex.JSONFloat64 `json:"eqUsd"`
		NotionalLever okex.JSONFloat64 `json:"notionalLever,omitempty"`
		StgyEq        okex.JSONFloat64 `json:"stgyEq"`
		IsoUpl        okex.JSONFloat64 `json:"isoUpl,omitempty"`
		UTime         okex.JSONTime    `json:"uTime"`
	}
	Position struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		LiabCcy     string              `json:"liabCcy,omitempty"`
		OptVal      string              `json:"optVal,omitempty"`
		Ccy         string              `json:"ccy"`
		PosID       string              `json:"posId"`
		TradeID     string              `json:"tradeId"`
		Pos         okex.JSONFloat64    `json:"pos"`
		AvailPos    okex.JSONFloat64    `json:"availPos,omitempty"`
		AvgPx       okex.JSONFloat64    `json:"avgPx"`
		Upl         okex.JSONFloat64    `json:"upl"`
		UplRatio    okex.JSONFloat64    `json:"uplRatio"`
		Lever       okex.JSONFloat64    `json:"lever"`
		LiqPx       okex.JSONFloat64    `json:"liqPx,omitempty"`
		Imr         okex.JSONFloat64    `json:"imr,omitempty"`
		Margin      okex.JSONFloat64    `json:"margin,omitempty"`
		MgnRatio    okex.JSONFloat64    `json:"mgnRatio"`
		Mmr         okex.JSONFloat64    `json:"mmr"`
		Liab        okex.JSONFloat64    `json:"liab,omitempty"`
		Interest    okex.JSONFloat64    `json:"interest"`
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd"`
		ADL         okex.JSONFloat64    `json:"adl"`
		Last        okex.JSONFloat64    `json:"last"`
		DeltaBS     okex.JSONFloat64    `json:"deltaBS"`
		DeltaPA     okex.JSONFloat64    `json:"deltaPA"`
		GammaBS     okex.JSONFloat64    `json:"gammaBS"`
		GammaPA     okex.JSONFloat64    `json:"gammaPA"`
		ThetaBS     okex.JSONFloat64    `json:"thetaBS"`
		ThetaPA     okex.JSONFloat64    `json:"thetaPA"`
		VegaBS      okex.JSONFloat64    `json:"vegaBS"`
		VegaPA      okex.JSONFloat64    `json:"vegaPA"`
		PosSide     okex.PositionSide   `json:"posSide"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
		InstType    okex.InstrumentType `json:"instType"`
		CTime       okex.JSONTime       `json:"cTime"`
		UTime       okex.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType"`
		PTime     okex.JSONTime     `json:"pTime"`
		UTime     okex.JSONTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   okex.JSONFloat64                     `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string           `json:"ccy"`
		Eq    okex.JSONFloat64 `json:"eq"`
		DisEq okex.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy okex.JSONFloat64    `json:"notionalCcy"`
		Pos         okex.JSONFloat64    `json:"pos"`
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd"`
		PosSide     okex.PositionSide   `json:"posSide"`
		InstType    okex.InstrumentType `json:"instType"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstID    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillID    string              `json:"billId"`
		OrdID     string              `json:"ordId"`
		BalChg    okex.JSONFloat64    `json:"balChg"`
		PosBalChg okex.JSONFloat64    `json:"posBalChg"`
		Bal       okex.JSONFloat64    `json:"bal"`
		PosBal    okex.JSONFloat64    `json:"posBal"`
		Sz        okex.JSONFloat64    `json:"sz"`
		Pnl       okex.JSONFloat64    `json:"pnl"`
		Fee       okex.JSONFloat64    `json:"fee"`
		From      okex.AccountType    `json:"from,string"`
		To        okex.AccountType    `json:"to,string"`
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		Type      okex.BillType       `json:"type,string"`
		SubType   okex.BillSubType    `json:"subType,string"`
		TS        okex.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string            `json:"level"`
		LevelTmp   string            `json:"levelTmp"`
		AcctLv     string            `json:"acctLv"`
		AutoLoan   bool              `json:"autoLoan"`
		UID        string            `json:"uid"`
		GreeksType okex.GreekType    `json:"greeksType"`
		PosMode    okex.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string            `json:"instId"`
		Lever   okex.JSONFloat64  `json:"lever"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string           `json:"instId"`
		Ccy     string           `json:"ccy"`
		MaxBuy  okex.JSONFloat64 `json:"maxBuy"`
		MaxSell okex.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string           `json:"instId"`
		AvailBuy  okex.JSONFloat64 `json:"availBuy"`
		AvailSell okex.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string            `json:"instId"`
		Amt     okex.JSONFloat64  `json:"amt"`
		PosSide okex.PositionSide `json:"posSide,string"`
		Type    okex.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string           `json:"instId"`
		MgnCcy  string           `json:"mgnCcy"`
		Ccy     string           `json:"ccy"`
		MaxLoan okex.JSONFloat64 `json:"maxLoan"`
		MgnMode okex.MarginMode  `json:"mgnMode"`
		Side    okex.OrderSide   `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    okex.JSONFloat64    `json:"taker"`
		Maker    okex.JSONFloat64    `json:"maker"`
		Delivery okex.JSONFloat64    `json:"delivery,omitempty"`
		Exercise okex.JSONFloat64    `json:"exercise,omitempty"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string           `json:"instId"`
		Ccy          string           `json:"ccy"`
		Interest     okex.JSONFloat64 `json:"interest"`
		InterestRate okex.JSONFloat64 `json:"interestRate"`
		Liab         okex.JSONFloat64 `json:"liab"`
		MgnMode      okex.MarginMode  `json:"mgnMode"`
		TS           okex.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string           `json:"ccy"`
		InterestRate okex.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string           `json:"ccy"`
		MaxWd okex.JSONFloat64 `json:"maxWd"`
	}
)
