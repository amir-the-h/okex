package public_data

import (
	"github.com/amir-the-h/okex/models/public_data"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetInstruments struct {
		responses.Basic
		Instruments []*public_data.Instrument `json:"data,omitempty"`
	}
	GetDeliveryExerciseHistory struct {
		responses.Basic
		Histories []*public_data.DeliveryExerciseHistory `json:"data,omitempty"`
	}
	GetOpenInterest struct {
		responses.Basic
		OpenInterests []*public_data.OpenInterest `json:"data,omitempty"`
	}
	GetFundingRate struct {
		responses.Basic
		FundingRates []*public_data.FundingRate `json:"data,omitempty"`
	}
	GetLimitPrice struct {
		responses.Basic
		LimitPrices []*public_data.LimitPrice `json:"data,omitempty"`
	}
	GetOptionMarketData struct {
		responses.Basic
		OptionMarketData []*public_data.OptionMarketData `json:"data,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		responses.Basic
		EstimatedDeliveryExercisePrices []*public_data.EstimatedDeliveryExercisePrice `json:"data,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		responses.Basic
		GetDiscountRateAndInterestFreeQuotas []*public_data.GetDiscountRateAndInterestFreeQuota `json:"data,omitempty"`
	}
	GetSystemTime struct {
		responses.Basic
		SystemTimes []*public_data.SystemTime `json:"data,omitempty"`
	}
	GetLiquidationOrders struct {
		responses.Basic
		LiquidationOrders []*public_data.LiquidationOrder `json:"data,omitempty"`
	}
	GetMarkPrice struct {
		responses.Basic
		MarkPrices []*public_data.MarkPrice `json:"data,omitempty"`
	}
	GetPositionTiers struct {
		responses.Basic
		PositionTiers []*public_data.PositionTier `json:"data,omitempty"`
	}
	GetInterestRateAndLoanQuota struct {
		responses.Basic
		InterestRateAndLoanQuotas []*public_data.InterestRateAndLoanQuota `json:"data,omitempty"`
	}
	GetUnderlying struct {
		responses.Basic
		Underlings [][]string `json:"data,omitempty"`
	}
)
