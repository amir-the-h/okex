package sub_account

import "github.com/amir-the-h/okex"

type (
	ViewList struct {
		SubAcct string `json:"subAcct,omitempty"`
		Enable  bool   `json:"enable,omitempty"`
		After   int64  `json:"after,omitempty,string"`
		Before  int64  `json:"before,omitempty,string"`
		Limit   int64  `json:"limit,omitempty,string"`
	}
	CreateAPIKey struct {
		Pwd        string            `json:"pwd"`
		SubAcct    string            `json:"subAcct"`
		Label      string            `json:"label"`
		Passphrase string            `json:"Passphrase"`
		Ip         []string          `json:"ip,omitempty"`
		Perm       okex.APIKeyAccess `json:"perm,omitempty"`
	}
	QueryAPIKey struct {
		ApiKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	DeleteAPIKey struct {
		Pwd     string `json:"pwd"`
		ApiKey  string `json:"apiKey"`
		SubAcct string `json:"subAcct"`
	}
	GetBalance struct {
		SubAcct string `json:"subAcct"`
	}
	HistoryTransfer struct {
		Ccy     string            `json:"ccy,omitempty"`
		SubAcct string            `json:"subAcct,omitempty"`
		After   int64             `json:"after,omitempty,string"`
		Before  int64             `json:"before,omitempty,string"`
		Limit   int64             `json:"limit,omitempty,string"`
		Type    okex.TransferType `json:"type,omitempty"`
	}
	ManageTransfers struct {
		Ccy            string           `json:"ccy"`
		FromSubAccount string           `json:"fromSubAccount"`
		ToSubAccount   string           `json:"tiSubAccount"`
		Amt            float64          `json:"amt,string"`
		From           okex.AccountType `json:"from"`
		To             okex.AccountType `json:"to"`
	}
)
