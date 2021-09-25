package account

import (
	models "github.com/amir-the-h/okex/models/account"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetBalance struct {
		responses.Basic
		Balances []*models.Balance `json:"data,omitempty"`
	}
	GetPositions struct {
		responses.Basic
		Positions []*models.Position `json:"data"`
	}
	GetAccountAndPositionRisk struct {
		responses.Basic
		PositionAndAccountRisks []*models.PositionAndAccountRisk `json:"data"`
	}
	GetBills struct {
		responses.Basic
		Bills []*models.Bill `json:"data"`
	}
	GetConfig struct {
		responses.Basic
		Configs []*models.Config `json:"data"`
	}
	SetPositionMode struct {
		responses.Basic
		PositionModes []*models.PositionMode `json:"data"`
	}
	Leverage struct {
		responses.Basic
		Leverages []*models.Leverage `json:"data"`
	}
	GetMaxBuySellAmount struct {
		responses.Basic
		MaxBuySellAmounts []*models.MaxBuySellAmount `json:"data"`
	}
	GetMaxAvailableTradeAmount struct {
		responses.Basic
		MaxAvailableTradeAmounts []*models.MaxAvailableTradeAmount `json:"data"`
	}
	IncreaseDecreaseMargin struct {
		responses.Basic
		MarginBalanceAmounts []*models.MarginBalanceAmount `json:"data"`
	}
	GetMaxLoan struct {
		responses.Basic
		Loans []*models.Loan `json:"data"`
	}
	GetFeeRates struct {
		responses.Basic
		Fees []*models.Fee `json:"data"`
	}
	GetInterestAccrued struct {
		responses.Basic
		InterestAccrues []*models.InterestAccrued `json:"data"`
	}
	GetInterestRates struct {
		responses.Basic
		Interests []*models.InterestRate `json:"data"`
	}
	SetGreeks struct {
		responses.Basic
		Greeks []*models.Greek `json:"data"`
	}
	GetMaxWithdrawals struct {
		responses.Basic
		MaxWithdrawals []*models.MaxWithdrawal `json:"data"`
	}
)
