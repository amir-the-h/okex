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
		InstID    string              `json:"instId,string"`
		Last      okex.JSONFloat64    `json:"last,string"`
		LastSz    okex.JSONFloat64    `json:"lastSz,string"`
		AskPx     okex.JSONFloat64    `json:"askPx,string"`
		AskSz     okex.JSONFloat64    `json:"askSz,string"`
		BidPx     okex.JSONFloat64    `json:"bidPx,string"`
		BidSz     okex.JSONFloat64    `json:"bidSz,string"`
		Open24h   okex.JSONFloat64    `json:"open24h,string"`
		High24h   okex.JSONFloat64    `json:"high24h,string"`
		Low24h    okex.JSONFloat64    `json:"low24h,string"`
		VolCcy24h okex.JSONFloat64    `json:"volCcy24h,string"`
		Vol24h    okex.JSONFloat64    `json:"vol24h,string"`
		SodUtc0   okex.JSONFloat64    `json:"sodUtc0,string"`
		SodUtc8   okex.JSONFloat64    `json:"sodUtc8,string"`
		InstType  okex.InstrumentType `json:"instType,string"`
		TS        okex.JSONTime       `json:"ts"`
	}
	IndexTicker struct {
		InstID  string           `json:"instId,string"`
		IdxPx   okex.JSONFloat64 `json:"idxPx,string"`
		High24h okex.JSONFloat64 `json:"high24h,string"`
		Low24h  okex.JSONFloat64 `json:"low24h,string"`
		Open24h okex.JSONFloat64 `json:"open24h,string"`
		SodUtc0 okex.JSONFloat64 `json:"sodUtc0,string"`
		SodUtc8 okex.JSONFloat64 `json:"sodUtc8,string"`
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
		Checksum int                `json:"checksum,string"`
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
		InstID  string           `json:"instId,string"`
		TradeID okex.JSONFloat64 `json:"tradeId,string"`
		Px      okex.JSONFloat64 `json:"px,string"`
		Sz      okex.JSONFloat64 `json:"sz,string"`
		Side    okex.TradeSide   `json:"side,string"`
		TS      okex.JSONTime    `json:"ts"`
	}
	TotalVolume24H struct {
		VolUsd okex.JSONFloat64 `json:"volUsd,string"`
		VolCny okex.JSONFloat64 `json:"volCny,string"`
		TS     okex.JSONTime    `json:"ts"`
	}
	IndexComponent struct {
		Index      string           `json:"index,string"`
		Last       okex.JSONFloat64 `json:"last,string"`
		Components []*Component     `json:"components"`
		TS         okex.JSONTime    `json:"ts"`
	}
	Component struct {
		Exch   string           `json:"exch,string"`
		Symbol string           `json:"symbol,string"`
		SymPx  okex.JSONFloat64 `json:"symPx,string"`
		Wgt    okex.JSONFloat64 `json:"wgt,string"`
		CnvPx  okex.JSONFloat64 `json:"cnvPx,string"`
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
