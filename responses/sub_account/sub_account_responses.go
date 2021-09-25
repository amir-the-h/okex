package sub_account

import (
	"github.com/amir-the-h/okex/models/account"
	models "github.com/amir-the-h/okex/models/sub_account"
	"github.com/amir-the-h/okex/responses"
)

type (
	ViewList struct {
		responses.RestBasic
		SubAccounts []*models.SubAccount `json:"data,omitempty"`
	}
	APIKey struct {
		responses.RestBasic
		APIKeys []*models.APIKey `json:"data,omitempty"`
	}
	GetBalance struct {
		responses.RestBasic
		Balances []*account.Balance `json:"data,omitempty"`
	}
	HistoryTransfer struct {
		responses.RestBasic
		HistoryTransfers []*models.HistoryTransfer `json:"data,omitempty"`
	}
	ManageTransfer struct {
		responses.RestBasic
		Transfers []*models.Transfer `json:"data,omitempty"`
	}
)
