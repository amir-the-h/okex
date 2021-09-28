package funding

import "github.com/amir-the-h/okex"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	FundsTransfer struct {
		Ccy      string            `json:"ccy,string"`
		Amt      float64           `json:"amt,string"`
		SubAcct  string            `json:"subAcct,string,omitempty"`
		InstID   string            `json:"instId,string,omitempty"`
		ToInstID string            `json:"toInstId,string,omitempty"`
		Type     okex.TransferType `json:"type,omitempty,string"`
		From     okex.AccountType  `json:"from,string"`
		To       okex.AccountType  `json:"to,string"`
	}
	AssetBillsDetails struct {
		Type   okex.BillType `json:"type,string,omitempty"`
		After  int64         `json:"after,string,omitempty"`
		Before int64         `json:"before,string,omitempty"`
		Limit  int64         `json:"limit,string,omitempty"`
	}
	GetDepositAddress struct {
		Ccy string `json:"ccy"`
	}
	GetDepositHistory struct {
		Ccy    string            `json:"ccy,string,omitempty"`
		TxID   string            `json:"txId,string,omitempty"`
		After  int64             `json:"after,omitempty,string"`
		Before int64             `json:"before,omitempty,string"`
		Limit  int64             `json:"limit,omitempty,string"`
		State  okex.DepositState `json:"state,omitempty,string"`
	}
	Withdrawal struct {
		Ccy    string                     `json:"ccy,string"`
		Chain  string                     `json:"chain,string,omitempty"`
		ToAddr string                     `json:"toAddr,string"`
		Pwd    string                     `json:"pwd,string"`
		Amt    float64                    `json:"amt,string"`
		Fee    float64                    `json:"fee,string"`
		Dest   okex.WithdrawalDestination `json:"dest,string"`
	}
	GetWithdrawalHistory struct {
		Ccy    string               `json:"ccy,string,omitempty"`
		TxID   string               `json:"txId,string,omitempty"`
		After  int64                `json:"after,omitempty,string"`
		Before int64                `json:"before,omitempty,string"`
		Limit  int64                `json:"limit,omitempty,string"`
		State  okex.WithdrawalState `json:"state,omitempty,string"`
	}
	PiggyBankPurchaseRedemption struct {
		Ccy    string               `json:"ccy,string,omitempty"`
		TxID   string               `json:"txId,string,omitempty"`
		After  int64                `json:"after,omitempty,string"`
		Before int64                `json:"before,omitempty,string"`
		Limit  int64                `json:"limit,omitempty,string"`
		State  okex.WithdrawalState `json:"state,omitempty,string"`
	}
	GetPiggyBankBalance struct {
		Ccy string `json:"ccy,omitempty"`
	}
)
