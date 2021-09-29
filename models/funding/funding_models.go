package funding

import "github.com/amir-the-h/okex"

type (
	Currency struct {
		Ccy         string `json:"ccy,string"`
		Name        string `json:"name,string"`
		Chain       string `json:"chain,string"`
		MinWd       string `json:"minWd,string"`
		MinFee      string `json:"minFee,string"`
		MaxFee      string `json:"maxFee,string"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy,string"`
		Bal       string `json:"bal,string"`
		FrozenBal string `json:"frozenBal,string"`
		AvailBal  string `json:"availBal,string"`
	}
	Transfer struct {
		TransID string           `json:"transId,string"`
		Ccy     string           `json:"ccy,string"`
		Amt     okex.JSONFloat64 `json:"amt,string"`
		From    okex.AccountType `json:"from,string"`
		To      okex.AccountType `json:"to,string"`
	}
	Bill struct {
		BillID string           `json:"billId,string"`
		Ccy    string           `json:"ccy,string"`
		Bal    okex.JSONFloat64 `json:"bal,string"`
		BalChg okex.JSONFloat64 `json:"balChg,string"`
		Type   okex.BillType    `json:"type,string"`
		TS     okex.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr,string"`
		Tag      string           `json:"tag,string,omitempty"`
		Memo     string           `json:"memo,string,omitempty"`
		PmtID    string           `json:"pmtId,string,omitempty"`
		Ccy      string           `json:"ccy,string"`
		Chain    string           `json:"chain,string"`
		CtAddr   string           `json:"ctAddr,string"`
		Selected bool             `json:"selected"`
		To       okex.AccountType `json:"to,string"`
		TS       okex.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string            `json:"ccy,string"`
		Chain string            `json:"chain,string"`
		TxID  string            `json:"txId,string"`
		From  string            `json:"from,string"`
		To    string            `json:"to,string"`
		DepId string            `json:"depId,string"`
		Amt   okex.JSONFloat64  `json:"amt,string"`
		State okex.DepositState `json:"state,string"`
		TS    okex.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string           `json:"ccy,string"`
		Chain string           `json:"chain,string"`
		WdID  okex.JSONInt64   `json:"wdId,string"`
		Amt   okex.JSONFloat64 `json:"amt,string"`
	}
	WithdrawalHistory struct {
		Ccy   string               `json:"ccy,string"`
		Chain string               `json:"chain,string"`
		TxID  string               `json:"txId,string"`
		From  string               `json:"from,string"`
		To    string               `json:"to,string"`
		Tag   string               `json:"tag,string,omitempty"`
		PmtID string               `json:"pmtId,string,omitempty"`
		Memo  string               `json:"memo,string,omitempty"`
		Amt   okex.JSONFloat64     `json:"amt,string"`
		Fee   okex.JSONInt64       `json:"fee,string"`
		WdID  okex.JSONInt64       `json:"wdId,string"`
		State okex.WithdrawalState `json:"state,string"`
		TS    okex.JSONTime        `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string           `json:"ccy,string"`
		Amt  okex.JSONFloat64 `json:"amt,string"`
		Side okex.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string           `json:"ccy,string"`
		Amt      okex.JSONFloat64 `json:"amt,string"`
		Earnings okex.JSONFloat64 `json:"earnings,string"`
	}
)
