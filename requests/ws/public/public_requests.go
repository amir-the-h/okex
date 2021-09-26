package public

import "github.com/amir-the-h/okex"

type (
	Instruments struct {
		InstType okex.InstrumentType `json:"instType"`
	}
	Tickers struct {
		InstId string `json:"instId"`
	}
	OpenInterest struct {
		InstId string `json:"instId"`
	}
	Candlesticks struct {
		InstId  string                    `json:"instId"`
		Channel okex.CandleStickWsBarSize `json:"channel"`
	}
	Trades struct {
		InstId string `json:"instId"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstId   string              `json:"instId"`
		Uly      string              `json:"uly,omitempty"`
		InstType okex.InstrumentType `json:"instType,omitempty"`
	}
	MarkPrice struct {
		InstId string `json:"instId"`
	}
	MarkPriceCandlesticks struct {
		InstId  string                    `json:"instId"`
		Channel okex.CandleStickWsBarSize `json:"channel"`
	}
	PriceLimit struct {
		InstId string `json:"instId"`
	}
	OrderBook struct {
		InstId  string `json:"instId"`
		Channel string `json:"channel"`
	}
	OPTIONSummary struct {
		InstId string `json:"instId"`
		Uly    string `json:"uly"`
	}
	FundingRate struct {
		InstId string `json:"instId"`
	}
	IndexCandlesticks struct {
		InstId  string `json:"instId"`
		Channel string `json:"channel"`
	}
	IndexTickers struct {
		InstId string `json:"instId"`
	}
)
