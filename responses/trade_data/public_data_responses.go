package trade_data

import (
	"github.com/amir-the-h/okex/models/tradedata"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetSupportCoin struct {
		responses.Basic
		SupportCoins *tradedata.SupportCoin `json:"data,omitempty"`
	}
	GetTakerVolume struct {
		responses.Basic
		TakerVolumes []*tradedata.TakerVolume `json:"data,omitempty"`
	}
	GetRatio struct {
		responses.Basic
		Ratios []*tradedata.Ratio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolume struct {
		responses.Basic
		InterestAndVolumeRatios []*tradedata.InterestAndVolumeRatio `json:"data,omitempty"`
	}
	GetPutCallRatio struct {
		responses.Basic
		PutCallRatios []*tradedata.PutCallRatio `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeExpiry struct {
		responses.Basic
		InterestAndVolumeExpires []*tradedata.InterestAndVolumeExpiry `json:"data,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		responses.Basic
		InterestAndVolumeStrikes []*tradedata.InterestAndVolumeStrike `json:"data,omitempty"`
	}
	GetTakerFlow struct {
		responses.Basic
		TakerFlow *tradedata.TakerFlow `json:"data"`
	}
)
