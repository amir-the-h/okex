package account

import "github.com/amir-the-h/okex"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositions struct {
		InstID   []string            `json:"instId,omitempty"`
		PosID    []string            `json:"posId,omitempty"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	GetAccountAndPositionRisk struct {
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	GetBills struct {
		Ccy      string              `json:"ccy,omitempty"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
		MgnMode  okex.MarginMode     `json:"mgnMode,omitempty"`
		CtType   okex.ContractType   `json:"ctType,omitempty"`
		Type     okex.BillType       `json:"type,omitempty"`
		SubType  okex.BillSubType    `json:"subType,omitempty"`
	}
	SetPositionMode struct {
		PositionMode okex.PositionType `json:"positionMode"`
	}
	SetLeverage struct {
		Lever   int64             `json:"lever,string"`
		InstID  string            `json:"instId,omitempty"`
		Ccy     string            `json:"ccy,omitempty"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide,omitempty"`
	}
	GetMaxBuySellAmount struct {
		Ccy    string         `json:"ccy,omitempty"`
		Px     float64        `json:"px,string,omitempty"`
		InstID []string       `json:"instId"`
		TdMode okex.TradeMode `json:"tdMode"`
	}
	GetMaxAvailableTradeAmount struct {
		Ccy        string         `json:"ccy,omitempty"`
		InstID     string         `json:"instId"`
		ReduceOnly bool           `json:"reduceOnly,omitempty"`
		TdMode     okex.TradeMode `json:"tdMode"`
	}
	IncreaseDecreaseMargin struct {
		InstID     string            `json:"instId"`
		Amt        float64           `json:"amt,string"`
		PosSide    okex.PositionSide `json:"posSide"`
		ActionType okex.CountAction  `json:"actionType"`
	}
	GetLeverage struct {
		InstID  []string        `json:"instId"`
		MgnMode okex.MarginMode `json:"mgnMode"`
	}
	GetMaxLoan struct {
		InstID  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy,omitempty"`
		MgnMode okex.MarginMode `json:"mgnMode"`
	}
	GetFeeRates struct {
		InstID   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		Category okex.FeeCategory    `json:"category,omitempty"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetInterestAccrued struct {
		InstID  string          `json:"instId,omitempty"`
		Ccy     string          `json:"ccy,omitempty"`
		After   int64           `json:"after,omitempty,string"`
		Before  int64           `json:"before,omitempty,string"`
		Limit   int64           `json:"limit,omitempty,string"`
		MgnMode okex.MarginMode `json:"mgnMode,omitempty"`
	}
	SetGreeks struct {
		GreeksType okex.GreekType `json:"greeksType"`
	}
)
