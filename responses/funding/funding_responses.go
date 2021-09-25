package funding

import (
	models "github.com/amir-the-h/okex/models/funding"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetCurrencies struct {
		responses.Basic
		Currencies []*models.Currency `json:"data"`
	}
	GetBalance struct {
		responses.Basic
		Balances []*models.Balance `json:"data"`
	}
	FundsTransfer struct {
		responses.Basic
		Transfers []*models.Transfer `json:"data"`
	}
	AssetBillsDetails struct {
		responses.Basic
		Bills []*models.Bill `json:"data"`
	}
	GetDepositAddress struct {
		responses.Basic
		DepositAddresses []*models.DepositAddress `json:"data"`
	}
	GetDepositHistory struct {
		responses.Basic
		DepositHistories []*models.DepositHistory `json:"data"`
	}
	Withdrawal struct {
		responses.Basic
		Withdrawals []*models.Withdrawal `json:"data"`
	}
	GetWithdrawalHistory struct {
		responses.Basic
		WithdrawalHistories []*models.WithdrawalHistory `json:"data"`
	}
	PiggyBankPurchaseRedemption struct {
		responses.Basic
		PiggyBanks []*models.PiggyBank `json:"data"`
	}
	GetPiggyBankBalance struct {
		responses.Basic
		Balances []*models.PiggyBankBalance `json:"data"`
	}
)
