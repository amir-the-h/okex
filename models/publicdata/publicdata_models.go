package publicdata

import (
	"github.com/marperia/okex"
)

type (
	Instrument struct {
		InstID    string               `json:"instId"`
		Uly       string               `json:"uly,omitempty"`
		BaseCcy   string               `json:"baseCcy,omitempty"`
		QuoteCcy  string               `json:"quoteCcy,omitempty"`
		SettleCcy string               `json:"settleCcy,omitempty"`
		CtValCcy  string               `json:"ctValCcy,omitempty"`
		CtVal     okex.JSONFloat64     `json:"ctVal,omitempty"`
		CtMult    okex.JSONFloat64     `json:"ctMult,omitempty"`
		Stk       okex.JSONFloat64     `json:"stk,omitempty"`
		TickSz    okex.JSONFloat64     `json:"tickSz,omitempty"`
		LotSz     okex.JSONFloat64     `json:"lotSz,omitempty"`
		MinSz     okex.JSONFloat64     `json:"minSz,omitempty"`
		Lever     okex.JSONFloat64     `json:"lever"`
		InstType  okex.InstrumentType  `json:"instType"`
		Category  okex.FeeCategory     `json:"category,string"`
		OptType   okex.OptionType      `json:"optType,omitempty"`
		ListTime  okex.JSONTime        `json:"listTime"`
		ExpTime   okex.JSONTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		Alias     okex.AliasType       `json:"alias,omitempty"`
		State     okex.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     okex.JSONFloat64          `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       okex.JSONFloat64    `json:"oi"`
		OiCcy    okex.JSONFloat64    `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     okex.JSONFloat64    `json:"fundingRate"`
		NextFundingRate okex.JSONFloat64    `json:"NextFundingRate"`
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   okex.JSONFloat64    `json:"buyLmt"`
		SellLmt  okex.JSONFloat64    `json:"sellLmt"`
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx okex.JSONFloat64    `json:"settlePx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    okex.JSONFloat64    `json:"delta"`
		Gamma    okex.JSONFloat64    `json:"gamma"`
		Vega     okex.JSONFloat64    `json:"vega"`
		Theta    okex.JSONFloat64    `json:"theta"`
		DeltaBS  okex.JSONFloat64    `json:"deltaBS"`
		GammaBS  okex.JSONFloat64    `json:"gammaBS"`
		VegaBS   okex.JSONFloat64    `json:"vegaBS"`
		ThetaBS  okex.JSONFloat64    `json:"thetaBS"`
		Lever    okex.JSONFloat64    `json:"lever"`
		MarkVol  okex.JSONFloat64    `json:"markVol"`
		BidVol   okex.JSONFloat64    `json:"bidVol"`
		AskVol   okex.JSONFloat64    `json:"askVol"`
		RealVol  okex.JSONFloat64    `json:"realVol"`
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string           `json:"ccy"`
		Amt          okex.JSONFloat64 `json:"amt"`
		DiscountLv   okex.JSONInt64   `json:"discountLv"`
		DiscountInfo []*DiscountInfo  `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okex.JSONInt64 `json:"discountRate"`
		MaxAmt       okex.JSONInt64 `json:"maxAmt"`
		MinAmt       okex.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		TS okex.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex.InstrumentType       `json:"instType"`
		TotalLoss okex.JSONFloat64          `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    okex.OrderSide    `json:"side"`
		OosSide okex.PositionSide `json:"posSide"`
		BkPx    okex.JSONFloat64  `json:"bkPx"`
		Sz      okex.JSONFloat64  `json:"sz"`
		BkLoss  okex.JSONFloat64  `json:"bkLoss"`
		TS      okex.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   okex.JSONFloat64    `json:"markPx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         okex.JSONInt64      `json:"tier"`
		MinSz        okex.JSONFloat64    `json:"minSz"`
		MaxSz        okex.JSONFloat64    `json:"maxSz"`
		Mmr          okex.JSONFloat64    `json:"mmr"`
		Imr          okex.JSONFloat64    `json:"imr"`
		OptMgnFactor okex.JSONFloat64    `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okex.JSONFloat64    `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okex.JSONFloat64    `json:"baseMaxLoan,omitempty"`
		MaxLever     okex.JSONFloat64    `json:"maxLever"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string           `json:"ccy"`
		Rate  okex.JSONFloat64 `json:"rate"`
		Quota okex.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string           `json:"level"`
		IrDiscount    okex.JSONFloat64 `json:"irDiscount"`
		LoanQuotaCoef int              `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string        `json:"title"`
		State       string        `json:"state"`
		Href        string        `json:"href"`
		ServiceType string        `json:"serviceType"`
		System      string        `json:"system"`
		ScheDesc    string        `json:"scheDesc"`
		Begin       okex.JSONTime `json:"begin"`
		End         okex.JSONTime `json:"end"`
	}
)
