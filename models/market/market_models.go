package market

import (
	"encoding/json"
	"fmt"
	"github.com/amir-the-h/okex"
	"strconv"
	"time"
)

type (
	Ticker struct {
		InstID    string              `json:"instId"`
		Last      okex.JSONFloat64    `json:"last"`
		LastSz    okex.JSONFloat64    `json:"lastSz"`
		AskPx     okex.JSONFloat64    `json:"askPx"`
		AskSz     okex.JSONFloat64    `json:"askSz"`
		BidPx     okex.JSONFloat64    `json:"bidPx"`
		BidSz     okex.JSONFloat64    `json:"bidSz"`
		Open24h   okex.JSONFloat64    `json:"open24h"`
		High24h   okex.JSONFloat64    `json:"high24h"`
		Low24h    okex.JSONFloat64    `json:"low24h"`
		VolCcy24h okex.JSONFloat64    `json:"volCcy24h"`
		Vol24h    okex.JSONFloat64    `json:"vol24h"`
		SodUtc0   okex.JSONFloat64    `json:"sodUtc0"`
		SodUtc8   okex.JSONFloat64    `json:"sodUtc8"`
		InstType  okex.InstrumentType `json:"instType"`
		TS        okex.JSONTime       `json:"ts"`
	}
	IndexTicker struct {
		InstID  string           `json:"instId"`
		IdxPx   okex.JSONFloat64 `json:"idxPx"`
		High24h okex.JSONFloat64 `json:"high24h"`
		Low24h  okex.JSONFloat64 `json:"low24h"`
		Open24h okex.JSONFloat64 `json:"open24h"`
		SodUtc0 okex.JSONFloat64 `json:"sodUtc0"`
		SodUtc8 okex.JSONFloat64 `json:"sodUtc8"`
		TS      okex.JSONTime    `json:"ts"`
	}
	OrderBook struct {
		Asks []*OrderBookEntity `json:"asks"`
		Bids []*OrderBookEntity `json:"bids"`
		TS   okex.JSONTime      `json:"ts"`
	}
	OrderBookWs struct {
		Asks     []*OrderBookEntity `json:"asks"`
		Bids     []*OrderBookEntity `json:"bids"`
		Checksum int                `json:"checksum"`
		TS       okex.JSONTime      `json:"ts"`
	}
	OrderBookEntity struct {
		DepthPrice      float64
		Size            float64
		LiquidatedOrder int
		OrderNumbers    int
	}
	Candle struct {
		O      float64
		H      float64
		L      float64
		C      float64
		Vol    float64
		VolCcy float64
		TS     okex.JSONTime
	}
	IndexCandle struct {
		O  float64
		H  float64
		L  float64
		C  float64
		TS okex.JSONTime
	}
	Trade struct {
		InstID  string           `json:"instId"`
		TradeID okex.JSONFloat64 `json:"tradeId"`
		Px      okex.JSONFloat64 `json:"px"`
		Sz      okex.JSONFloat64 `json:"sz"`
		Side    okex.TradeSide   `json:"side,string"`
		TS      okex.JSONTime    `json:"ts"`
	}
	TotalVolume24H struct {
		VolUsd okex.JSONFloat64 `json:"volUsd"`
		VolCny okex.JSONFloat64 `json:"volCny"`
		TS     okex.JSONTime    `json:"ts"`
	}
	IndexComponent struct {
		Index      string           `json:"index"`
		Last       okex.JSONFloat64 `json:"last"`
		Components []*Component     `json:"components"`
		TS         okex.JSONTime    `json:"ts"`
	}
	Component struct {
		Exch   string           `json:"exch"`
		Symbol string           `json:"symbol"`
		SymPx  okex.JSONFloat64 `json:"symPx"`
		Wgt    okex.JSONFloat64 `json:"wgt"`
		CnvPx  okex.JSONFloat64 `json:"cnvPx"`
	}
)

func (o *OrderBookEntity) UnmarshalJSON(buf []byte) error {
	var (
		dp, s, lo, on string
		err           error
	)
	tmp := []interface{}{&dp, &s, &lo, &on}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in OrderBookEntity: %d != %d", g, e)
	}
	o.DepthPrice, err = strconv.ParseFloat(dp, 64)
	if err != nil {
		return err
	}
	o.Size, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	o.LiquidatedOrder, err = strconv.Atoi(lo)
	if err != nil {
		return err
	}
	o.OrderNumbers, err = strconv.Atoi(on)
	if err != nil {
		return err
	}

	return nil
}

func (c *Candle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, vol, volCcy, ts string
		err                          error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl, &vol, &volCcy}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	c.Vol, err = strconv.ParseFloat(vol, 64)
	if err != nil {
		return err
	}

	c.VolCcy, err = strconv.ParseFloat(volCcy, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *IndexCandle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, ts string
		err             error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.TS) = time.UnixMilli(timestamp)

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	return nil
}
