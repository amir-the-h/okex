package subaccount

import "github.com/amir-the-h/okex"

type (
	ViewList struct {
		SubAcct string `json:"subAcct,string,omitempty"`
		Enable  bool   `json:"enable,omitempty"`
		After   int64  `json:"after,omitempty,string"`
		Before  int64  `json:"before,omitempty,string"`
		Limit   int64  `json:"limit,omitempty,string"`
	}
	CreateAPIKey struct {
		Pwd        string            `json:"pwd,string"`
		SubAcct    string            `json:"subAcct,string"`
		Label      string            `json:"label,string"`
		Passphrase string            `json:"Passphrase,string"`
		Ip         []string          `json:"ip,omitempty"`
		Perm       okex.APIKeyAccess `json:"perm,string,omitempty"`
	}
	QueryAPIKey struct {
		ApiKey  string `json:"apiKey,string"`
		SubAcct string `json:"subAcct,string"`
	}
	DeleteAPIKey struct {
		Pwd     string `json:"pwd,string"`
		ApiKey  string `json:"apiKey,string"`
		SubAcct string `json:"subAcct,string"`
	}
	GetBalance struct {
		SubAcct string `json:"subAcct,string"`
	}
	HistoryTransfer struct {
		Ccy     string            `json:"ccy,string,omitempty"`
		SubAcct string            `json:"subAcct,string,omitempty"`
		After   int64             `json:"after,omitempty,string"`
		Before  int64             `json:"before,omitempty,string"`
		Limit   int64             `json:"limit,omitempty,string"`
		Type    okex.TransferType `json:"type,string,omitempty"`
	}
	ManageTransfers struct {
		Ccy            string           `json:"ccy,string"`
		FromSubAccount string           `json:"fromSubAccount,string"`
		ToSubAccount   string           `json:"tiSubAccount,string"`
		Amt            float64          `json:"amt,string"`
		From           okex.AccountType `json:"from,string"`
		To             okex.AccountType `json:"to,string"`
	}
)
