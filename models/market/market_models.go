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
		InstId    string              `json:"instId"`
		Last      okex.JsonFloat64    `json:"last"`
		LastSz    okex.JsonFloat64    `json:"lastSz"`
		AskPx     okex.JsonFloat64    `json:"askPx"`
		AskSz     okex.JsonFloat64    `json:"askSz"`
		BidPx     okex.JsonFloat64    `json:"bidPx"`
		BidSz     okex.JsonFloat64    `json:"bidSz"`
		Open24h   okex.JsonFloat64    `json:"open24h"`
		High24h   okex.JsonFloat64    `json:"high24h"`
		Low24h    okex.JsonFloat64    `json:"low24h"`
		VolCcy24h okex.JsonFloat64    `json:"volCcy24h"`
		Vol24h    okex.JsonFloat64    `json:"vol24h"`
		SodUtc0   okex.JsonFloat64    `json:"sodUtc0"`
		SodUtc8   okex.JsonFloat64    `json:"sodUtc8"`
		InstType  okex.InstrumentType `json:"instType"`
		Ts        okex.JsonTime       `json:"ts"`
	}
	IndexTicker struct {
		InstId  string           `json:"instId"`
		IdxPx   okex.JsonFloat64 `json:"idxPx"`
		High24h okex.JsonFloat64 `json:"high24h"`
		Low24h  okex.JsonFloat64 `json:"low24h"`
		Open24h okex.JsonFloat64 `json:"open24h"`
		SodUtc0 okex.JsonFloat64 `json:"sodUtc0"`
		SodUtc8 okex.JsonFloat64 `json:"sodUtc8"`
		Ts      okex.JsonTime    `json:"ts"`
	}
	OrderBook struct {
		Asks []*OrderBookEntity `json:"asks"`
		Bids []*OrderBookEntity `json:"bids"`
		Ts   okex.JsonTime      `json:"ts"`
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
		Ts     okex.JsonTime
	}
	IndexCandle struct {
		O  float64
		H  float64
		L  float64
		C  float64
		Ts okex.JsonTime
	}
	Trade struct {
		InstId  string           `json:"instId"`
		TradeId okex.JsonFloat64 `json:"tradeId"`
		Px      okex.JsonFloat64 `json:"px"`
		Sz      okex.JsonFloat64 `json:"sz"`
		Side    okex.TradeSide   `json:"side"`
		Ts      okex.JsonTime    `json:"ts"`
	}
	TotalVolume24H struct {
		VolUsd okex.JsonFloat64 `json:"volUsd"`
		VolCny okex.JsonFloat64 `json:"volCny"`
		Ts     okex.JsonTime    `json:"ts"`
	}
	IndexComponent struct {
		Index      string           `json:"index"`
		Last       okex.JsonFloat64 `json:"last"`
		Components []*Component     `json:"components"`
		Ts         okex.JsonTime    `json:"ts"`
	}
	Component struct {
		Exch   string           `json:"exch"`
		Symbol string           `json:"symbol"`
		SymPx  okex.JsonFloat64 `json:"symPx"`
		Wgt    okex.JsonFloat64 `json:"wgt"`
		CnvPx  okex.JsonFloat64 `json:"cnvPx"`
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
	*(*time.Time)(&c.Ts) = time.UnixMilli(timestamp)

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
	*(*time.Time)(&c.Ts) = time.UnixMilli(timestamp)

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
