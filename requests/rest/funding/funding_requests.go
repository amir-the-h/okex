package funding

import "github.com/amir-the-h/okex"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	FundsTransfer struct {
		Ccy      string            `json:"ccy"`
		Amt      float64           `json:"amt,string"`
		SubAcct  string            `json:"subAcct,omitempty"`
		InstId   string            `json:"instId,omitempty"`
		ToInstId string            `json:"toInstId,omitempty"`
		Type     okex.TransferType `json:"type,omitempty,string"`
		From     okex.AccountType  `json:"from,string"`
		To       okex.AccountType  `json:"to,string"`
	}
	AssetBillsDetails struct {
		Type   okex.BillType `json:"type,omitempty,string"`
		After  int64         `json:"after,omitempty,string"`
		Before int64         `json:"before,omitempty,string"`
		Limit  int64         `json:"limit,omitempty,string"`
	}
	GetDepositAddress struct {
		Ccy string `json:"ccy"`
	}
	GetDepositHistory struct {
		Ccy    string            `json:"ccy,omitempty"`
		TxId   string            `json:"txId,omitempty"`
		After  int64             `json:"after,omitempty,string"`
		Before int64             `json:"before,omitempty,string"`
		Limit  int64             `json:"limit,omitempty,string"`
		State  okex.DepositState `json:"state,omitempty,string"`
	}
	Withdrawal struct {
		Ccy    string                     `json:"ccy"`
		Chain  string                     `json:"chain,omitempty"`
		ToAddr string                     `json:"toAddr"`
		Pwd    string                     `json:"pwd"`
		Amt    float64                    `json:"amt,string"`
		Fee    float64                    `json:"fee"`
		Dest   okex.WithdrawalDestination `json:"dest,string"`
	}
	GetWithdrawalHistory struct {
		Ccy    string               `json:"ccy,omitempty"`
		TxId   string               `json:"txId,omitempty"`
		After  int64                `json:"after,omitempty,string"`
		Before int64                `json:"before,omitempty,string"`
		Limit  int64                `json:"limit,omitempty,string"`
		State  okex.WithdrawalState `json:"state,omitempty,string"`
	}
	PiggyBankPurchaseRedemption struct {
		Ccy    string               `json:"ccy,omitempty"`
		TxId   string               `json:"txId,omitempty"`
		After  int64                `json:"after,omitempty,string"`
		Before int64                `json:"before,omitempty,string"`
		Limit  int64                `json:"limit,omitempty,string"`
		State  okex.WithdrawalState `json:"state,omitempty,string"`
	}
	GetPiggyBankBalance struct {
		Ccy string `json:"ccy,omitempty"`
	}
)
