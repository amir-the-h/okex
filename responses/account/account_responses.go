package account

import (
	models "github.com/amir-the-h/okex/models/account"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetBalance struct {
		responses.RestBasic
		Balances []*models.Balance `json:"data,omitempty"`
	}
	GetPositions struct {
		responses.RestBasic
		Positions []*models.Position `json:"data"`
	}
	GetAccountAndPositionRisk struct {
		responses.RestBasic
		PositionAndAccountRisks []*models.PositionAndAccountRisk `json:"data"`
	}
	GetBills struct {
		responses.RestBasic
		Bills []*models.Bill `json:"data"`
	}
	GetConfig struct {
		responses.RestBasic
		Configs []*models.Config `json:"data"`
	}
	SetPositionMode struct {
		responses.RestBasic
		PositionModes []*models.PositionMode `json:"data"`
	}
	Leverage struct {
		responses.RestBasic
		Leverages []*models.Leverage `json:"data"`
	}
	GetMaxBuySellAmount struct {
		responses.RestBasic
		MaxBuySellAmounts []*models.MaxBuySellAmount `json:"data"`
	}
	GetMaxAvailableTradeAmount struct {
		responses.RestBasic
		MaxAvailableTradeAmounts []*models.MaxAvailableTradeAmount `json:"data"`
	}
	IncreaseDecreaseMargin struct {
		responses.RestBasic
		MarginBalanceAmounts []*models.MarginBalanceAmount `json:"data"`
	}
	GetMaxLoan struct {
		responses.RestBasic
		Loans []*models.Loan `json:"data"`
	}
	GetFeeRates struct {
		responses.RestBasic
		Fees []*models.Fee `json:"data"`
	}
	GetInterestAccrued struct {
		responses.RestBasic
		InterestAccrues []*models.InterestAccrued `json:"data"`
	}
	GetInterestRates struct {
		responses.RestBasic
		Interests []*models.InterestRate `json:"data"`
	}
	SetGreeks struct {
		responses.RestBasic
		Greeks []*models.Greek `json:"data"`
	}
	GetMaxWithdrawals struct {
		responses.RestBasic
		MaxWithdrawals []*models.MaxWithdrawal `json:"data"`
	}
)
