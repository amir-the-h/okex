package public_data

import (
	"github.com/amir-the-h/okex/models/publicdata"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetInstruments struct {
		responses.Basic
		Instruments []*publicdata.Instrument `json:"data,omitempty"`
	}
	GetDeliveryExerciseHistory struct {
		responses.Basic
		Histories []*publicdata.DeliveryExerciseHistory `json:"data,omitempty"`
	}
	GetOpenInterest struct {
		responses.Basic
		OpenInterests []*publicdata.OpenInterest `json:"data,omitempty"`
	}
	GetFundingRate struct {
		responses.Basic
		FundingRates []*publicdata.FundingRate `json:"data,omitempty"`
	}
	GetLimitPrice struct {
		responses.Basic
		LimitPrices []*publicdata.LimitPrice `json:"data,omitempty"`
	}
	GetOptionMarketData struct {
		responses.Basic
		OptionMarketData []*publicdata.OptionMarketData `json:"data,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		responses.Basic
		EstimatedDeliveryExercisePrices []*publicdata.EstimatedDeliveryExercisePrice `json:"data,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		responses.Basic
		GetDiscountRateAndInterestFreeQuotas []*publicdata.GetDiscountRateAndInterestFreeQuota `json:"data,omitempty"`
	}
	GetSystemTime struct {
		responses.Basic
		SystemTimes []*publicdata.SystemTime `json:"data,omitempty"`
	}
	GetLiquidationOrders struct {
		responses.Basic
		LiquidationOrders []*publicdata.LiquidationOrder `json:"data,omitempty"`
	}
	GetMarkPrice struct {
		responses.Basic
		MarkPrices []*publicdata.MarkPrice `json:"data,omitempty"`
	}
	GetPositionTiers struct {
		responses.Basic
		PositionTiers []*publicdata.PositionTier `json:"data,omitempty"`
	}
	GetInterestRateAndLoanQuota struct {
		responses.Basic
		InterestRateAndLoanQuotas []*publicdata.InterestRateAndLoanQuota `json:"data,omitempty"`
	}
	GetUnderlying struct {
		responses.Basic
		Underlings [][]string `json:"data,omitempty"`
	}
)
