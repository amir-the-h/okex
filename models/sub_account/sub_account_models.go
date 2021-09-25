package sub_account

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
		Ts      okex.JsonTime `json:"ts"`
	}
	APIKey struct {
		SubAcct    string        `json:"subAcct,omitempty"`
		Label      string        `json:"label,omitempty"`
		ApiKey     string        `json:"apiKey,omitempty"`
		SecretKey  string        `json:"secretKey,omitempty"`
		Passphrase string        `json:"Passphrase,omitempty"`
		Perm       string        `json:"perm,omitempty"`
		Ip         string        `json:"ip,omitempty"`
		Ts         okex.JsonTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string        `json:"subAcct,omitempty"`
		Ccy     string        `json:"ccy,omitempty"`
		BillId  int64         `json:"billId,omitempty,string"`
		Type    okex.BillType `json:"type,omitempty"`
		Ts      okex.JsonTime `json:"ts,omitempty"`
	}
	Transfer struct {
		TransId int64 `json:"transId,string"`
	}
)
