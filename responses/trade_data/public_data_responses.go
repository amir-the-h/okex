package trade_data

import (
	"github.com/amir-the-h/okex/models/trade_data"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetSupportCoin struct {
		responses.Basic
		SupportCoins *trade_data.SupportCoin `json:"data,omitempty"`
	}
	GetTakerVolume struct {
		responses.Basic
		TakerVolumes []*trade_data.TakerVolume `json:"data,omitempty"`
	}
	GetRatio struct {
		responses.Basic
		Ratios []*trade_data.Ratio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolume struct {
		responses.Basic
		InterestAndVolumeRatios []*trade_data.InterestAndVolumeRatio `json:"data,omitempty"`
	}
	GetPutCallRatio struct {
		responses.Basic
		PutCallRatios []*trade_data.PutCallRatio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeExpiry struct {
		responses.Basic
		InterestAndVolumeExpires []*trade_data.InterestAndVolumeExpiry `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		responses.Basic
		InterestAndVolumeStrikes []*trade_data.InterestAndVolumeStrike `json:"data,omitempty"`
	}
	GetTakerFlow struct {
		responses.Basic
		TakerFlow *trade_data.TakerFlow `json:"data"`
	}
)
