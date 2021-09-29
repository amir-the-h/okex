package publicdata

import (
	"github.com/amir-the-h/okex"
)

type (
	Instrument struct {
		InstID    string               `json:"instId,string"`
		Uly       string               `json:"uly,string,omitempty"`
		BaseCcy   string               `json:"baseCcy,string,omitempty"`
		QuoteCcy  string               `json:"quoteCcy,string,omitempty"`
		SettleCcy string               `json:"settleCcy,string,omitempty"`
		CtValCcy  string               `json:"ctValCcy,string,omitempty"`
		CtVal     okex.JSONFloat64     `json:"ctVal,string,omitempty"`
		CtMult    okex.JSONFloat64     `json:"ctMult,string,omitempty"`
		Stk       okex.JSONFloat64     `json:"stk,string,omitempty"`
		TickSz    okex.JSONFloat64     `json:"tickSz,string,omitempty"`
		LotSz     okex.JSONFloat64     `json:"lotSz,string,omitempty"`
		MinSz     okex.JSONFloat64     `json:"minSz,string,omitempty"`
		Lever     okex.JSONFloat64     `json:"lever,string"`
		InstType  okex.InstrumentType  `json:"instType,string"`
		Category  okex.FeeCategory     `json:"category,string"`
		OptType   okex.OptionType      `json:"optType,string,omitempty"`
		ListTime  okex.JSONTime        `json:"listTime"`
		ExpTime   okex.JSONTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,string,omitempty"`
		Alias     okex.AliasType       `json:"alias,string,omitempty"`
		State     okex.InstrumentState `json:"state,string"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId,string"`
		Px     okex.JSONFloat64          `json:"px,string"`
		Type   okex.DeliveryExerciseType `json:"type,string"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId,string"`
		Oi       okex.JSONFloat64    `json:"oi,string"`
		OiCcy    okex.JSONFloat64    `json:"oiCcy,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId,string"`
		InstType        okex.InstrumentType `json:"instType,string"`
		FundingRate     okex.JSONFloat64    `json:"fundingRate,string"`
		NextFundingRate okex.JSONFloat64    `json:"NextFundingRate,string"`
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		BuyLmt   okex.JSONFloat64    `json:"buyLmt,string"`
		SellLmt  okex.JSONFloat64    `json:"sellLmt,string"`
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		SettlePx okex.JSONFloat64    `json:"settlePx,string"`
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId,string"`
		Uly      string              `json:"uly,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		Delta    okex.JSONFloat64    `json:"delta,string"`
		Gamma    okex.JSONFloat64    `json:"gamma,string"`
		Vega     okex.JSONFloat64    `json:"vega,string"`
		Theta    okex.JSONFloat64    `json:"theta,string"`
		DeltaBS  okex.JSONFloat64    `json:"deltaBS,string"`
		GammaBS  okex.JSONFloat64    `json:"gammaBS,string"`
		VegaBS   okex.JSONFloat64    `json:"vegaBS,string"`
		ThetaBS  okex.JSONFloat64    `json:"thetaBS,string"`
		Lever    okex.JSONFloat64    `json:"lever,string"`
		MarkVol  okex.JSONFloat64    `json:"markVol,string"`
		BidVol   okex.JSONFloat64    `json:"bidVol,string"`
		AskVol   okex.JSONFloat64    `json:"askVol,string"`
		RealVol  okex.JSONFloat64    `json:"realVol,string"`
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string           `json:"ccy,string"`
		Amt          okex.JSONFloat64 `json:"amt,string"`
		DiscountLv   okex.JSONInt64   `json:"discountLv,string"`
		DiscountInfo []*DiscountInfo  `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okex.JSONInt64 `json:"discountRate,string"`
		MaxAmt       okex.JSONInt64 `json:"maxAmt,string"`
		MinAmt       okex.JSONInt64 `json:"minAmt,string"`
	}
	SystemTime struct {
		TS okex.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId,string"`
		Uly       string                    `json:"uly,string,omitempty"`
		InstType  okex.InstrumentType       `json:"instType,string"`
		TotalLoss okex.JSONFloat64          `json:"totalLoss,string"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,string,omitempty"`
		Side    okex.OrderSide    `json:"side,string"`
		OosSide okex.PositionSide `json:"posSide,string"`
		BkPx    okex.JSONFloat64  `json:"bkPx,string"`
		Sz      okex.JSONFloat64  `json:"sz,string"`
		BkLoss  okex.JSONFloat64  `json:"bkLoss,string"`
		TS      okex.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		MarkPx   okex.JSONFloat64    `json:"markPx,string"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId,string"`
		Uly          string              `json:"uly,string,omitempty"`
		InstType     okex.InstrumentType `json:"instType,string"`
		Tier         okex.JSONInt64      `json:"tier,string"`
		MinSz        okex.JSONFloat64    `json:"minSz,string"`
		MaxSz        okex.JSONFloat64    `json:"maxSz,string"`
		Mmr          okex.JSONFloat64    `json:"mmr,string"`
		Imr          okex.JSONFloat64    `json:"imr,string"`
		OptMgnFactor okex.JSONFloat64    `json:"optMgnFactor,string,omitempty"`
		QuoteMaxLoan okex.JSONFloat64    `json:"quoteMaxLoan,string,omitempty"`
		BaseMaxLoan  okex.JSONFloat64    `json:"baseMaxLoan,string,omitempty"`
		MaxLever     okex.JSONFloat64    `json:"maxLever,string"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string           `json:"ccy,string"`
		Rate  okex.JSONFloat64 `json:"rate,string"`
		Quota okex.JSONFloat64 `json:"quota,string"`
	}
	InterestRateAndLoanUser struct {
		Level         string           `json:"level,string"`
		IrDiscount    okex.JSONFloat64 `json:"irDiscount,string"`
		LoanQuotaCoef int              `json:"loanQuotaCoef,string"`
	}
)
