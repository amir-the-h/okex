package subaccount

import (
	"github.com/amir-the-h/okex"
)

type (
	SubAccount struct {
		SubAcct string        `json:"subAcct,string,omitempty"`
		Label   string        `json:"label,string,omitempty"`
		Mobile  string        `json:"mobile,string,omitempty"`
		GAuth   bool          `json:"gAuth"`
		Enable  bool          `json:"enable"`
		TS      okex.JSONTime `json:"ts"`
	}
	APIKey struct {
		SubAcct    string        `json:"subAcct,string,omitempty"`
		Label      string        `json:"label,string,omitempty"`
		ApiKey     string        `json:"apiKey,string,omitempty"`
		SecretKey  string        `json:"secretKey,string,omitempty"`
		Passphrase string        `json:"Passphrase,string,omitempty"`
		Perm       string        `json:"perm,string,omitempty"`
		Ip         string        `json:"ip,string,omitempty"`
		TS         okex.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string         `json:"subAcct,string,omitempty"`
		Ccy     string         `json:"ccy,string,omitempty"`
		BillId  okex.JSONInt64 `json:"billId,string,omitempty"`
		Type    okex.BillType  `json:"type,string,omitempty"`
		TS      okex.JSONTime  `json:"ts,omitempty"`
	}
	Transfer struct {
		TransId okex.JSONInt64 `json:"transId,string"`
	}
)
