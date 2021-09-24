package funding

import (
	models "github.com/amir-the-h/okex/models/funding"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetCurrencies struct {
		responses.RestBasic
		Currencies []*models.Currency `json:"data"`
	}
	GetBalance struct {
		responses.RestBasic
		Balances []*models.Balance `json:"data"`
	}
	FundsTransfer struct {
		responses.RestBasic
		Transfers []*models.Transfer `json:"data"`
	}
	AssetBillsDetails struct {
		responses.RestBasic
		Bills []*models.Bill `json:"data"`
	}
	GetDepositAddress struct {
		responses.RestBasic
		DepositAddresses []*models.DepositAddress `json:"data"`
	}
	GetDepositHistory struct {
		responses.RestBasic
		DepositHistories []*models.DepositHistory `json:"data"`
	}
	Withdrawal struct {
		responses.RestBasic
		Withdrawals []*models.Withdrawal `json:"data"`
	}
	GetWithdrawalHistory struct {
		responses.RestBasic
		WithdrawalHistories []*models.WithdrawalHistory `json:"data"`
	}
	PiggyBankPurchaseRedemption struct {
		responses.RestBasic
		PiggyBanks []*models.PiggyBank `json:"data"`
	}
	GetPiggyBankBalance struct {
		responses.RestBasic
		Balances []*models.PiggyBankBalance `json:"data"`
	}
)
