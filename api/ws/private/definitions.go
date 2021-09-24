package private

import (
	"github.com/amir-the-h/okex/models/account"
	"github.com/amir-the-h/okex/responses"
)

type (
	Login struct {
		Event string `json:"event"`
		Code  string `json:"code"`
		Msg   string `json:"msg"`
	}
	Account struct {
		Arg      *responses.WsArgument `json:"arg"`
		Balances []*account.Balance    `json:"data"`
	}
	Position struct {
		Arg       *responses.WsArgument `json:"arg"`
		Positions []*account.Position   `json:"data"`
	}
	BalanceAndPosition struct {
		Arg                 *responses.WsArgument         `json:"arg"`
		BalanceAndPositions []*account.BalanceAndPosition `json:"data"`
	}
)
