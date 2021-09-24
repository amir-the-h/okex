package trade

import (
	"github.com/amir-the-h/okex/models/trade"
	"github.com/amir-the-h/okex/responses"
)

type (
	PlaceOrder struct {
		responses.RestBasic
		PlaceOrders []*trade.PlaceOrder `json:"data"`
	}
	CancelOrder struct {
		responses.RestBasic
		CancelOrders []*trade.CancelOrder `json:"data"`
	}
	AmendOrder struct {
		responses.RestBasic
		AmendOrders []*trade.AmendOrder `json:"data"`
	}
	ClosePosition struct {
		responses.RestBasic
		ClosePositions []*trade.ClosePosition `json:"data"`
	}
	Order struct {
		responses.RestBasic
		Orders []*trade.Order `json:"data"`
	}
	TransactionDetail struct {
		responses.RestBasic
		TransactionDetails []*trade.TransactionDetail `json:"data"`
	}
)
