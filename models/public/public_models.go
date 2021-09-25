package public

import (
	"github.com/amir-the-h/okex"
)

type (
	Instrument struct {
		InstId    string               `json:"instId"`
		Uly       string               `json:"uly,omitempty"`
		BaseCcy   string               `json:"baseCcy,omitempty"`
		QuoteCcy  string               `json:"quoteCcy,omitempty"`
		SettleCcy string               `json:"settleCcy,omitempty"`
		CtValCcy  string               `json:"ctValCcy,omitempty"`
		CtVal     okex.JsonFloat64     `json:"ctVal,omitempty"`
		CtMult    okex.JsonFloat64     `json:"ctMult,omitempty"`
		Stk       okex.JsonFloat64     `json:"stk,omitempty"`
		TickSz    okex.JsonFloat64     `json:"tickSz,omitempty"`
		LotSz     okex.JsonFloat64     `json:"lotSz,omitempty"`
		MinSz     okex.JsonFloat64     `json:"minSz,omitempty"`
		Lever     okex.JsonFloat64     `json:"lever"`
		InstType  okex.InstrumentType  `json:"instType"`
		Category  okex.FeeCategory     `json:"category"`
		OptType   okex.OptionType      `json:"optType,omitempty"`
		ListTime  okex.JsonTime        `json:"listTime"`
		ExpTime   okex.JsonTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		Alias     okex.AliasType       `json:"alias,omitempty"`
		State     okex.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		Ts      okex.JsonTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstId string                     `json:"instId"`
		Px     okex.JsonFloat64           `json:"px"`
		Type   okex.DeliveryExcerciseType `json:"type"`
	}
	OpenInterest struct {
		InstId   string              `json:"instId"`
		Oi       okex.JsonFloat64    `json:"oi"`
		OiCcy    okex.JsonFloat64    `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		Ts       okex.JsonTime       `json:"ts"`
	}
	FundingRate struct {
		InstId          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     okex.JsonFloat64    `json:"fundingRate"`
		NextFundingRate okex.JsonFloat64    `json:"NextFundingRate"`
		FundingTime     okex.JsonTime       `json:"fundingTime"`
		NextFundingTime okex.JsonTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstId   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   okex.JsonFloat64    `json:"buyLmt"`
		SellLmt  okex.JsonFloat64    `json:"sellLmt"`
		Ts       okex.JsonTime       `json:"ts"`
	}
	EstimatedDeliveryExcercisePrice struct {
		InstId   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx okex.JsonFloat64    `json:"settlePx"`
		Ts       okex.JsonTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstId   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    okex.JsonFloat64    `json:"delta"`
		Gamma    okex.JsonFloat64    `json:"gamma"`
		Vega     okex.JsonFloat64    `json:"vega"`
		Theta    okex.JsonFloat64    `json:"theta"`
		DeltaBS  okex.JsonFloat64    `json:"deltaBS"`
		GammaBS  okex.JsonFloat64    `json:"gammaBS"`
		VegaBS   okex.JsonFloat64    `json:"vegaBS"`
		ThetaBS  okex.JsonFloat64    `json:"thetaBS"`
		Lever    okex.JsonFloat64    `json:"lever"`
		MarkVol  okex.JsonFloat64    `json:"markVol"`
		BidVol   okex.JsonFloat64    `json:"bidVol"`
		AskVol   okex.JsonFloat64    `json:"askVol"`
		RealVol  okex.JsonFloat64    `json:"realVol"`
		Ts       okex.JsonTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string           `json:"ccy"`
		Amt          okex.JsonFloat64 `json:"amt"`
		DiscountLv   okex.JsonInt64   `json:"discountLv"`
		DiscountInfo []*DiscountInfo  `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okex.JsonInt64 `json:"discountRate"`
		MaxAmt       okex.JsonInt64 `json:"maxAmt"`
		MinAmt       okex.JsonInt64 `json:"minAmt"`
	}
	SystemTime struct {
		Ts okex.JsonTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstId    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex.InstrumentType       `json:"instType"`
		TotalLoss okex.JsonFloat64          `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    okex.OrderSide    `json:"side"`
		OosSide okex.PositionSide `json:"posSide"`
		BkPx    okex.JsonFloat64  `json:"bkPx"`
		Sz      okex.JsonFloat64  `json:"sz"`
		BkLoss  okex.JsonFloat64  `json:"bkLoss"`
		Ts      okex.JsonTime     `json:"ts"`
	}
	MarkPrice struct {
		InstId   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   okex.JsonFloat64    `json:"markPx"`
		Ts       okex.JsonTime       `json:"ts"`
	}
	PositionTier struct {
		InstId       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         okex.JsonInt64      `json:"tier"`
		MinSz        okex.JsonFloat64    `json:"minSz"`
		MaxSz        okex.JsonFloat64    `json:"maxSz"`
		Mmr          okex.JsonFloat64    `json:"mmr"`
		Imr          okex.JsonFloat64    `json:"imr"`
		OptMgnFactor okex.JsonFloat64    `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okex.JsonFloat64    `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okex.JsonFloat64    `json:"baseMaxLoan,omitempty"`
		MaxLever     okex.JsonFloat64    `json:"maxLever"`
		Ts           okex.JsonTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string           `json:"ccy"`
		Rate  okex.JsonFloat64 `json:"rate"`
		Quota okex.JsonFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string           `json:"level"`
		IrDiscount    okex.JsonFloat64 `json:"irDiscount"`
		LoanQuotaCoef int              `json:"loanQuotaCoef"`
	}
)
