package public_data

import (
	"github.com/amir-the-h/okex/models/public_data"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetInstruments struct {
		responses.RestBasic
		Instruments []*public_data.Instrument `json:"data,omitempty"`
	}
	GetDeliveryExerciseHistory struct {
		responses.RestBasic
		Histories []*public_data.DeliveryExerciseHistory `json:"data,omitempty"`
	}
	GetOpenInterest struct {
		responses.RestBasic
		OpenInterests []*public_data.OpenInterest `json:"data,omitempty"`
	}
	GetFundingRate struct {
		responses.RestBasic
		FundingRates []*public_data.FundingRate `json:"data,omitempty"`
	}
	GetLimitPrice struct {
		responses.RestBasic
		LimitPrices []*public_data.LimitPrice `json:"data,omitempty"`
	}
	GetOptionMarketData struct {
		responses.RestBasic
		OptionMarketData []*public_data.OptionMarketData `json:"data,omitempty"`
	}
	GetEstimatedDeliveryExcercisePrice struct {
		responses.RestBasic
		EstimatedDeliveryExcercisePrices []*public_data.EstimatedDeliveryExcercisePrice `json:"data,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		responses.RestBasic
		GetDiscountRateAndInterestFreeQuotas []*public_data.GetDiscountRateAndInterestFreeQuota `json:"data,omitempty"`
	}
	GetSystemTime struct {
		responses.RestBasic
		SystemTimes []*public_data.SystemTime `json:"data,omitempty"`
	}
	GetLiquidationOrders struct {
		responses.RestBasic
		LiquidationOrders []*public_data.LiquidationOrder `json:"data,omitempty"`
	}
	GetMarkPrice struct {
		responses.RestBasic
		MarkPrices []*public_data.MarkPrice `json:"data,omitempty"`
	}
	GetPositionTiers struct {
		responses.RestBasic
		PositionTiers []*public_data.PositionTier `json:"data,omitempty"`
	}
	GetInterestRateAndLoanQuota struct {
		responses.RestBasic
		InterestRateAndLoanQuotas []*public_data.InterestRateAndLoanQuota `json:"data,omitempty"`
	}
	GetUnderlying struct {
		responses.RestBasic
		Underlings [][]string `json:"data,omitempty"`
	}
)
