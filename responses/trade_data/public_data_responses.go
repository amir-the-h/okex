package trade_data

import (
	"github.com/amir-the-h/okex/models/trade_data"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetSupportCoin struct {
		responses.RestBasic
		SupportCoins *trade_data.SupportCoin `json:"data,omitempty"`
	}
	GetTakerVolume struct {
		responses.RestBasic
		TakerVolumes []*trade_data.TakerVolume `json:"data,omitempty"`
	}
	GetRatio struct {
		responses.RestBasic
		Ratios []*trade_data.Ratio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolume struct {
		responses.RestBasic
		InterestAndVolumeRatios []*trade_data.InterestAndVolumeRatio `json:"data,omitempty"`
	}
	GetPutCallRatio struct {
		responses.RestBasic
		PutCallRatios []*trade_data.PutCallRatio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeExpiry struct {
		responses.RestBasic
		InterestAndVolumeExpires []*trade_data.InterestAndVolumeExpiry `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		responses.RestBasic
		InterestAndVolumeStrikes []*trade_data.InterestAndVolumeStrike `json:"data,omitempty"`
	}
	GetTakerFlow struct {
		responses.RestBasic
		TakerFlow *trade_data.TakerFlow `json:"data"`
	}
)
