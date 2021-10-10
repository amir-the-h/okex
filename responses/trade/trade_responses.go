package trade

import (
	"github.com/amir-the-h/okex/models/trade"
	"github.com/amir-the-h/okex/responses"
)

type (
	PlaceOrder struct {
		responses.Basic
		PlaceOrders []*trade.PlaceOrder `json:"data"`
	}
	CancelOrder struct {
		responses.Basic
		CancelOrders []*trade.CancelOrder `json:"data"`
	}
	AmendOrder struct {
		responses.Basic
		AmendOrders []*trade.AmendOrder `json:"data"`
	}
	ClosePosition struct {
		responses.Basic
		ClosePositions []*trade.ClosePosition `json:"data"`
	}
	OrderList struct {
		responses.Basic
		Orders []*trade.Order `json:"data"`
	}
	TransactionDetail struct {
		responses.Basic
		TransactionDetails []*trade.TransactionDetail `json:"data"`
	}
	PlaceAlgoOrder struct {
		responses.Basic
		PlaceAlgoOrders []*trade.PlaceAlgoOrder `json:"data"`
	}
	CancelAlgoOrder struct {
		responses.Basic
		CancelAlgoOrders []*trade.CancelAlgoOrder `json:"data"`
	}
	AlgoOrderList struct {
		responses.Basic
		AlgoOrders []*trade.AlgoOrder `json:"data"`
	}
)
