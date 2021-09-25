package public

import "github.com/amir-the-h/okex"

type (
	GetInstruments struct {
		Uly      string              `json:"uly,omitempty"`
		InstId   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetDeliveryExerciseHistory struct {
		Uly      string              `json:"uly"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetOpenInterest struct {
		Uly      string              `json:"uly,omitempty"`
		InstId   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetFundingRate struct {
		InstId string `json:"instId"`
	}
	GetLimitPrice struct {
		InstId string `json:"instId"`
	}
	GetOptionMarketData struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetEstimatedDeliveryExcercisePrice struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Uly        string  `json:"uly"`
		Ccy        string  `json:"ccy,omitempty"`
		DiscountLv float64 `json:"discountLv,string"`
	}
	GetLiquidationOrders struct {
		InstId   string              `json:"instId,omitempty"`
		Ccy      string              `json:"ccy,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType"`
		MgnMode  okex.MarginMode     `json:"mgnMode,omitempty"`
		Alias    okex.AliasType      `json:"alias,omitempty"`
		State    okex.OrderState     `json:"state,omitempty"`
	}
	GetMarkPrice struct {
		InstId   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetPositionTiers struct {
		InstId   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
		TdMode   okex.TradeMode      `json:"tdMode"`
		Tier     okex.JsonInt64      `json:"tier,omitempty"`
	}
	GetUnderlying struct {
		InstType okex.InstrumentType `json:"instType"`
	}
)
