package public

import (
	"github.com/amir-the-h/okex/models/public"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetInstruments struct {
		responses.RestBasic
		Instruments []*public.Instrument `json:"data,omitempty"`
	}
	GetDeliveryExerciseHistory struct {
		responses.RestBasic
		Histories []*public.DeliveryExerciseHistory `json:"data,omitempty"`
	}
	GetOpenInterest struct {
		responses.RestBasic
		OpenInterests []*public.OpenInterest `json:"data,omitempty"`
	}
	GetFundingRate struct {
		responses.RestBasic
		FundingRates []*public.FundingRate `json:"data,omitempty"`
	}
	GetLimitPrice struct {
		responses.RestBasic
		LimitPrices []*public.LimitPrice `json:"data,omitempty"`
	}
	GetOptionMarketData struct {
		responses.RestBasic
		OptionMarketData []*public.OptionMarketData `json:"data,omitempty"`
	}
	GetEstimatedDeliveryExcercisePrice struct {
		responses.RestBasic
		EstimatedDeliveryExcercisePrices []*public.EstimatedDeliveryExcercisePrice `json:"data,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		responses.RestBasic
		GetDiscountRateAndInterestFreeQuotas []*public.GetDiscountRateAndInterestFreeQuota `json:"data,omitempty"`
	}
	GetSystemTime struct {
		responses.RestBasic
		SystemTimes []*public.SystemTime `json:"data,omitempty"`
	}
	GetLiquidationOrders struct {
		responses.RestBasic
		LiquidationOrders []*public.LiquidationOrder `json:"data,omitempty"`
	}
	GetMarkPrice struct {
		responses.RestBasic
		MarkPrices []*public.MarkPrice `json:"data,omitempty"`
	}
	GetPositionTiers struct {
		responses.RestBasic
		PositionTiers []*public.PositionTier `json:"data,omitempty"`
	}
	GetInterestRateAndLoanQuota struct {
		responses.RestBasic
		InterestRateAndLoanQuotas []*public.InterestRateAndLoanQuota `json:"data,omitempty"`
	}
	GetUnderlying struct {
		responses.RestBasic
		Underlings [][]string `json:"data,omitempty"`
	}
)
