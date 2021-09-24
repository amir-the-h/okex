package funding

import "github.com/amir-the-h/okex"

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransId string           `json:"transId"`
		Ccy     string           `json:"ccy"`
		Amt     float64          `json:"amt,string"`
		From    okex.AccountType `json:"from"`
		To      okex.AccountType `json:"to"`
	}
	Bill struct {
		BillId string        `json:"billId"`
		Ccy    string        `json:"ccy"`
		Bal    float64       `json:"bal,string"`
		BalChg float64       `json:"balChg,string"`
		Type   okex.BillType `json:"type,string"`
		Ts     okex.JsonTime `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr"`
		Tag      string           `json:"tag,omitempty"`
		Memo     string           `json:"memo,omitempty"`
		PmtId    string           `json:"pmtId,omitempty"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain"`
		CtAddr   string           `json:"ctAddr"`
		Selected bool             `json:"selected"`
		To       okex.AccountType `json:"to,string"`
		Ts       okex.JsonTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string            `json:"ccy"`
		Chain string            `json:"chain"`
		TxId  string            `json:"txId"`
		From  string            `json:"from"`
		To    string            `json:"to"`
		DepId string            `json:"depId"`
		Amt   float64           `json:"amt,string"`
		State okex.DepositState `json:"state,string"`
		Ts    okex.JsonTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string  `json:"ccy"`
		Chain string  `json:"chain"`
		WdId  int64   `json:"wdId,string"`
		Amt   float64 `json:"amt,string"`
	}
	WithdrawalHistory struct {
		Ccy   string               `json:"ccy"`
		Chain string               `json:"chain"`
		TxId  string               `json:"txId"`
		From  string               `json:"from"`
		To    string               `json:"to"`
		Tag   string               `json:"tag,omitempty"`
		PmtId string               `json:"pmtId,omitempty"`
		Memo  string               `json:"memo,omitempty"`
		Amt   float64              `json:"amt,string"`
		Fee   int64                `json:"fee,string"`
		WdId  int64                `json:"wdId,string"`
		State okex.WithdrawalState `json:"state,string"`
		Ts    okex.JsonTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string          `json:"ccy"`
		Amt  float64         `json:"amt,string"`
		Side okex.ActionType `json:"side"`
	}
	PiggyBankBalance struct {
		Ccy      string  `json:"ccy"`
		Amt      float64 `json:"amt,string"`
		Earnings float64 `json:"earnings,string"`
	}
)
