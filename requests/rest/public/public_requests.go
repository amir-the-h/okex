package public

import "github.com/amir-the-h/okex"

type (
	GetInstruments struct {
		Uly      string              `json:"uly,string,omitempty"`
		InstID   string              `json:"instId,string,omitempty"`
		InstType okex.InstrumentType `json:"instType,string"`
	}
	GetDeliveryExerciseHistory struct {
		Uly      string              `json:"uly,string"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,string"`
	}
	GetOpenInterest struct {
		Uly      string              `json:"uly,string,omitempty"`
		InstID   string              `json:"instId,string,omitempty"`
		InstType okex.InstrumentType `json:"instType,string"`
	}
	GetFundingRate struct {
		InstID string `json:"instId,string"`
	}
	GetLimitPrice struct {
		InstID string `json:"instId,string"`
	}
	GetOptionMarketData struct {
		Uly     string `json:"uly,string"`
		ExpTime string `json:"expTime,string,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		Uly     string `json:"uly,string"`
		ExpTime string `json:"expTime,string,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Uly        string  `json:"uly,string"`
		Ccy        string  `json:"ccy,string,omitempty"`
		DiscountLv float64 `json:"discountLv,string"`
	}
	GetLiquidationOrders struct {
		InstID   string              `json:"instId,string,omitempty"`
		Ccy      string              `json:"ccy,string,omitempty"`
		Uly      string              `json:"uly,string,omitempty"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		MgnMode  okex.MarginMode     `json:"mgnMode,string,omitempty"`
		Alias    okex.AliasType      `json:"alias,string,omitempty"`
		State    okex.OrderState     `json:"state,string,omitempty"`
	}
	GetMarkPrice struct {
		InstID   string              `json:"instId,string,omitempty"`
		Uly      string              `json:"uly,string,omitempty"`
		InstType okex.InstrumentType `json:"instType,string"`
	}
	GetPositionTiers struct {
		InstID   string              `json:"instId,string,omitempty"`
		Uly      string              `json:"uly,string,omitempty"`
		InstType okex.InstrumentType `json:"instType,string"`
		TdMode   okex.TradeMode      `json:"tdMode,string"`
		Tier     okex.JSONInt64      `json:"tier,string,omitempty"`
	}
	GetUnderlying struct {
		InstType okex.InstrumentType `json:"instType,string"`
	}
)
