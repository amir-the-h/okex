package public

import "github.com/amir-the-h/okex"

type (
	Instruments struct {
		InstType okex.InstrumentType `json:"instType"`
	}
	Tickers struct {
		InstID string `json:"instId"`
	}
	OpenInterest struct {
		InstID string `json:"instId"`
	}
	Candlesticks struct {
		InstID  string                    `json:"instId"`
		Channel okex.CandleStickWsBarSize `json:"channel"`
	}
	Trades struct {
		InstID string `json:"instId"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly,omitempty"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	MarkPrice struct {
		InstID string `json:"instId"`
	}
	MarkPriceCandlesticks struct {
		InstID  string                    `json:"instId"`
		Channel okex.CandleStickWsBarSize `json:"channel"`
	}
	PriceLimit struct {
		InstID string `json:"instId"`
	}
	OrderBook struct {
		InstID  string `json:"instId"`
		Channel string `json:"channel"`
	}
	OPTIONSummary struct {
		InstID string `json:"instId"`
		Uly    string `json:"uly"`
	}
	FundingRate struct {
		InstID string `json:"instId"`
	}
	IndexCandlesticks struct {
		InstID  string `json:"instId"`
		Channel string `json:"channel"`
	}
	IndexTickers struct {
		InstID string `json:"instId"`
	}
)
