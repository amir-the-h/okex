package account

import (
	"github.com/amir-the-h/okex"
)

type (
	Balance struct {
		TotalEq     okex.JSONFloat64  `json:"totalEq,string"`
		IsoEq       okex.JSONFloat64  `json:"isoEq,string"`
		AdjEq       okex.JSONFloat64  `json:"adjEq,string,omitempty"`
		OrdFroz     okex.JSONFloat64  `json:"ordFroz,string,omitempty"`
		Imr         okex.JSONFloat64  `json:"imr,string,omitempty"`
		Mmr         okex.JSONFloat64  `json:"mmr,string,omitempty"`
		MgnRatio    okex.JSONFloat64  `json:"mgnRatio,string,omitempty"`
		NotionalUsd okex.JSONFloat64  `json:"notionalUsd,string,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okex.JSONTime     `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string           `json:"ccy,string"`
		Eq            okex.JSONFloat64 `json:"eq,string"`
		CashBal       okex.JSONFloat64 `json:"cashBal,string"`
		IsoEq         okex.JSONFloat64 `json:"isoEq,string,omitempty"`
		AvailEq       okex.JSONFloat64 `json:"availEq,string,omitempty"`
		DisEq         okex.JSONFloat64 `json:"disEq,string"`
		AvailBal      okex.JSONFloat64 `json:"availBal,string"`
		FrozenBal     okex.JSONFloat64 `json:"frozenBal,string"`
		OrdFrozen     okex.JSONFloat64 `json:"ordFrozen,string"`
		Liab          okex.JSONFloat64 `json:"liab,string,omitempty"`
		Upl           okex.JSONFloat64 `json:"upl,string,omitempty"`
		UplLib        okex.JSONFloat64 `json:"uplLib,string,omitempty"`
		CrossLiab     okex.JSONFloat64 `json:"crossLiab,string,omitempty"`
		IsoLiab       okex.JSONFloat64 `json:"isoLiab,string,omitempty"`
		MgnRatio      okex.JSONFloat64 `json:"mgnRatio,string,omitempty"`
		Interest      okex.JSONFloat64 `json:"interest,string,omitempty"`
		Twap          okex.JSONFloat64 `json:"twap,string,omitempty"`
		MaxLoan       okex.JSONFloat64 `json:"maxLoan,string,omitempty"`
		EqUsd         okex.JSONFloat64 `json:"eqUsd,string"`
		NotionalLever okex.JSONFloat64 `json:"notionalLever,string,omitempty"`
		StgyEq        okex.JSONFloat64 `json:"stgyEq,string"`
		IsoUpl        okex.JSONFloat64 `json:"isoUpl,string,omitempty"`
		UTime         okex.JSONTime    `json:"uTime"`
	}
	Position struct {
		InstID      string              `json:"instId,string"`
		PosCcy      string              `json:"posCcy,string,omitempty"`
		LiabCcy     string              `json:"liabCcy,string,omitempty"`
		OptVal      string              `json:"optVal,string,omitempty"`
		Ccy         string              `json:"ccy,string"`
		PosId       okex.JSONFloat64    `json:"posId,string"`
		Pos         okex.JSONFloat64    `json:"pos,string"`
		AvailPos    okex.JSONFloat64    `json:"availPos,string,omitempty"`
		AvgPx       okex.JSONFloat64    `json:"avgPx,string"`
		Upl         okex.JSONFloat64    `json:"upl,string"`
		UplRatio    okex.JSONFloat64    `json:"uplRatio,string"`
		Lever       okex.JSONFloat64    `json:"lever,string"`
		LiqPx       okex.JSONFloat64    `json:"liqPx,string,omitempty"`
		Imr         okex.JSONFloat64    `json:"imr,string,omitempty"`
		Margin      okex.JSONFloat64    `json:"margin,string,omitempty"`
		MgnRatio    okex.JSONFloat64    `json:"mgnRatio,string"`
		Mmr         okex.JSONFloat64    `json:"mmr,string"`
		Liab        okex.JSONFloat64    `json:"liab,string,omitempty"`
		Interest    okex.JSONFloat64    `json:"interest,string"`
		TradeId     okex.JSONFloat64    `json:"tradeId,string"`
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd,string"`
		ADL         okex.JSONFloat64    `json:"adl,string"`
		Last        okex.JSONFloat64    `json:"last,string"`
		DeltaBS     okex.JSONFloat64    `json:"deltaBS,string"`
		DeltaPA     okex.JSONFloat64    `json:"deltaPA,string"`
		GammaBS     okex.JSONFloat64    `json:"gammaBS,string"`
		GammaPA     okex.JSONFloat64    `json:"gammaPA,string"`
		ThetaBS     okex.JSONFloat64    `json:"thetaBS,string"`
		ThetaPA     okex.JSONFloat64    `json:"thetaPA,string"`
		VegaBS      okex.JSONFloat64    `json:"vegaBS,string"`
		VegaPA      okex.JSONFloat64    `json:"vegaPA,string"`
		PosSide     okex.PositionSide   `json:"posSide,string"`
		MgnMode     okex.MarginMode     `json:"mgnMode,string"`
		InstType    okex.InstrumentType `json:"instType,string"`
		CTime       okex.JSONTime       `json:"cTime"`
		UTime       okex.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType,string"`
		PTime     okex.JSONTime     `json:"pTime"`
		UTime     okex.JSONTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   okex.JSONFloat64                     `json:"adjEq,string,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string           `json:"ccy,string"`
		Eq    okex.JSONFloat64 `json:"eq,string"`
		DisEq okex.JSONFloat64 `json:"disEq,string"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId,string"`
		PosCcy      string              `json:"posCcy,string,omitempty"`
		Ccy         string              `json:"ccy,string"`
		NotionalCcy okex.JSONFloat64    `json:"notionalCcy,string"`
		Pos         okex.JSONFloat64    `json:"pos,string"`
		NotionalUsd okex.JSONFloat64    `json:"notionalUsd,string"`
		PosSide     okex.PositionSide   `json:"posSide,string"`
		InstType    okex.InstrumentType `json:"instType,string"`
		MgnMode     okex.MarginMode     `json:"mgnMode,string"`
	}
	Bill struct {
		Ccy       string              `json:"ccy,string"`
		InstID    string              `json:"instId,string"`
		Notes     string              `json:"notes,string"`
		BillID    okex.JSONFloat64    `json:"billId,string"`
		BalChg    okex.JSONFloat64    `json:"balChg,string"`
		PosBalChg okex.JSONFloat64    `json:"posBalChg,string"`
		Bal       okex.JSONFloat64    `json:"bal,string"`
		PosBal    okex.JSONFloat64    `json:"posBal,string"`
		Sz        okex.JSONFloat64    `json:"sz,string"`
		Pnl       okex.JSONFloat64    `json:"pnl,string"`
		Fee       okex.JSONFloat64    `json:"fee,string"`
		OrdId     okex.JSONFloat64    `json:"ordId,string"`
		From      okex.AccountType    `json:"from,string"`
		To        okex.AccountType    `json:"to,string"`
		InstType  okex.InstrumentType `json:"instType,string"`
		MgnMode   okex.MarginMode     `json:"MgnMode,string"`
		Type      okex.BillType       `json:"type,string"`
		SubType   okex.BillSubType    `json:"subType,string"`
		TS        okex.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string           `json:"level,string"`
		LevelTmp   string           `json:"levelTmp,string"`
		AcctLv     string           `json:"acctLv,string"`
		AutoLoan   bool             `json:"autoLoan"`
		Uid        okex.JSONFloat64 `json:"uid,string"`
		GreeksType okex.GreekType   `json:"greeksType"`
		PosMode    PositionMode     `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode,string"`
	}
	Leverage struct {
		InstID  string            `json:"instId,string"`
		Lever   okex.JSONFloat64  `json:"lever,string"`
		MgnMode okex.MarginMode   `json:"mgnMode,string"`
		PosSide okex.PositionSide `json:"posSide,string"`
	}
	MaxBuySellAmount struct {
		InstID  string           `json:"instId,string"`
		Ccy     string           `json:"ccy,string"`
		MaxBuy  okex.JSONFloat64 `json:"maxBuy,string"`
		MaxSell okex.JSONFloat64 `json:"maxSell,string"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string           `json:"instId,string"`
		AvailBuy  okex.JSONFloat64 `json:"availBuy,string"`
		AvailSell okex.JSONFloat64 `json:"availSell,string"`
	}
	MarginBalanceAmount struct {
		InstID  string            `json:"instId,string"`
		Amt     okex.JSONFloat64  `json:"amt,string"`
		PosSide okex.PositionSide `json:"posSide,string"`
		Type    okex.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string           `json:"instId,string"`
		MgnCcy  string           `json:"mgnCcy,string"`
		Ccy     string           `json:"ccy,string"`
		MaxLoan okex.JSONFloat64 `json:"maxLoan,string"`
		MgnMode okex.MarginMode  `json:"mgnMode,string"`
		Side    okex.OrderSide   `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level,string"`
		Taker    okex.JSONFloat64    `json:"taker,string"`
		Maker    okex.JSONFloat64    `json:"maker,string"`
		Delivery okex.JSONFloat64    `json:"delivery,string,omitempty"`
		Exercise okex.JSONFloat64    `json:"exercise,string,omitempty"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		TS       okex.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string           `json:"instId,string"`
		Ccy          string           `json:"ccy,string"`
		Interest     okex.JSONFloat64 `json:"interest,string"`
		InterestRate okex.JSONFloat64 `json:"interestRate,string"`
		Liab         okex.JSONFloat64 `json:"liab,string"`
		MgnMode      okex.MarginMode  `json:"mgnMode,string"`
		TS           okex.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string           `json:"ccy,string"`
		InterestRate okex.JSONFloat64 `json:"interestRate,string"`
	}
	Greek struct {
		GreeksType string `json:"greeksType,string"`
	}
	MaxWithdrawal struct {
		Ccy   string           `json:"ccy,string"`
		MaxWd okex.JSONFloat64 `json:"maxWd,string"`
	}
)
