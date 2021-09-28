package private

import "github.com/amir-the-h/okex"

type (
	Account struct {
		Ccy string `json:"ccy,omitempty"`
	}
	Position struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	Order struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	AlgoOrder struct {
		Uly      string              `json:"uly,omitempty"`
		InstID   string              `json:"instId,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
)
