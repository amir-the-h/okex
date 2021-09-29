package subaccount

import (
	"github.com/amir-the-h/okex"
)

type (
	SubAccount struct {
		SubAcct string        `json:"subAcct,omitempty"`
		Label   string        `json:"label,omitempty"`
		Mobile  string        `json:"mobile,omitempty"`
		GAuth   bool          `json:"gAuth"`
		Enable  bool          `json:"enable"`
		TS      okex.JSONTime `json:"ts"`
	}
	APIKey struct {
		SubAcct    string        `json:"subAcct,omitempty"`
		Label      string        `json:"label,omitempty"`
		ApiKey     string        `json:"apiKey,omitempty"`
		SecretKey  string        `json:"secretKey,omitempty"`
		Passphrase string        `json:"Passphrase,omitempty"`
		Perm       string        `json:"perm,omitempty"`
		Ip         string        `json:"ip,omitempty"`
		TS         okex.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string         `json:"subAcct,omitempty"`
		Ccy     string         `json:"ccy,omitempty"`
		BillId  okex.JSONInt64 `json:"billId,omitempty"`
		Type    okex.BillType  `json:"type,string,omitempty"`
		TS      okex.JSONTime  `json:"ts,omitempty"`
	}
	Transfer struct {
		TransId okex.JSONInt64 `json:"transId"`
	}
)
