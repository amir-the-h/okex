package tradedata

import "github.com/amir-the-h/okex"

type (
	GetTakerVolume struct {
		Ccy      string              `json:"ccy"`
		Begin    int64               `json:"before,omitempty,string"`
		End      int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,string"`
		Period   okex.BarSize        `json:"period,string,omitempty"`
	}
	GetRatio struct {
		Ccy    string       `json:"ccy"`
		Begin  int64        `json:"before,omitempty,string"`
		End    int64        `json:"limit,omitempty,string"`
		Period okex.BarSize `json:"period,string,omitempty"`
	}
	GetOpenInterestAndVolumeStrike struct {
		Ccy     string       `json:"ccy"`
		ExpTime string       `json:"expTime"`
		Period  okex.BarSize `json:"period,string,omitempty"`
	}
)
