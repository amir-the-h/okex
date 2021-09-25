package account

import "github.com/amir-the-h/okex"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositions struct {
		InstId   []string            `json:"instId,omitempty"`
		PosId    []string            `json:"posId,omitempty"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
	}
	GetAccountAndPositionRisk struct {
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
	}
	GetBills struct {
		Ccy      string              `json:"ccy,omitempty"`
		After    int64               `json:"after,omitempty,string"`
		Before   int64               `json:"before,omitempty,string"`
		Limit    int64               `json:"limit,omitempty,string"`
		InstType okex.InstrumentType `json:"instType,omitempty,string"`
		MgnMode  okex.MarginMode     `json:"mgnMode,omitempty,string"`
		CtType   okex.ContractType   `json:"ctType,omitempty,string"`
		Type     okex.BillType       `json:"type,omitempty,string"`
		SubType  okex.BillSubType    `json:"subType,omitempty,string"`
	}
	SetPositionMode struct {
		PositionMode okex.PositionType `json:"positionMode,string"`
	}
	SetLeverage struct {
		Lever   int64             `json:"lever,string"`
		InstId  string            `json:"instId,omitempty"`
		Ccy     string            `json:"ccy,omitempty"`
		MgnMode okex.MarginMode   `json:"mgnMode,string"`
		PosSide okex.PositionSide `json:"posSide,omitempty,string"`
	}
	GetMaxBuySellAmount struct {
		Ccy    string         `json:"ccy,omitempty"`
		Px     float64        `json:"px,omitempty,string"`
		InstId []string       `json:"instId"`
		TdMode okex.TradeMode `json:"tdMode,string"`
	}
	GetMaxAvailableTradeAmount struct {
		Ccy        string         `json:"ccy,omitempty"`
		InstId     string         `json:"instId"`
		ReduceOnly bool           `json:"reduceOnly,omitempty"`
		TdMode     okex.TradeMode `json:"tdMode,string"`
	}
	IncreaseDecreaseMargin struct {
		InstId     string            `json:"instId"`
		Amt        float64           `json:"amt,string"`
		PosSide    okex.PositionSide `json:"posSide,string"`
		ActionType okex.CountAction  `json:"actionType,string"`
	}
	GetLeverage struct {
		InstId  []string        `json:"instId"`
		MgnMode okex.MarginMode `json:"mgnMode,string"`
	}
	GetMaxLoan struct {
		InstId  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy,omitempty"`
		MgnMode okex.MarginMode `json:"mgnMode,string"`
	}
	GetFeeRates struct {
		InstId   string              `json:"instId,omitempty"`
		Uly      string              `json:"uly,omitempty"`
		Category okex.FeeCategory    `json:"category,omitempty,string"`
		InstType okex.InstrumentType `json:"instType"`
	}
	GetInterestAccrued struct {
		InstId  string          `json:"instId,omitempty"`
		Ccy     string          `json:"ccy,omitempty"`
		After   int64           `json:"after,omitempty,string"`
		Before  int64           `json:"before,omitempty,string"`
		Limit   int64           `json:"limit,omitempty,string"`
		MgnMode okex.MarginMode `json:"mgnMode,omitempty,string"`
	}
	SetGreeks struct {
		GreeksType okex.GreekType `json:"greeksType,string"`
	}
)
