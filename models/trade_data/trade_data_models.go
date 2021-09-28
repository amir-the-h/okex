package trade_data

import (
	"encoding/json"
	"fmt"
	"github.com/amir-the-h/okex"
	"strconv"
	"time"
)

type (
	SupportCoin struct {
		Contract []string `json:"contract"`
		Option   []string `json:"option"`
		Spot     []string `json:"spot"`
	}
	TakerVolume struct {
		SellVol float64
		BuyVol  float64
		Ts      okex.JSONTime
	}
	Ratio struct {
		Ratio float64
		Ts    okex.JSONTime
	}
	InterestAndVolumeRatio struct {
		Oi  float64
		Vol float64
		Ts  okex.JSONTime
	}
	PutCallRatio struct {
		OiRatio  float64
		VolRatio float64
		Ts       okex.JSONTime
	}
	InterestAndVolumeExpiry struct {
		CallOI  float64
		PutOI   float64
		CallVol float64
		PutVol  float64
		ExpTime okex.JSONTime
		Ts      okex.JSONTime
	}
	InterestAndVolumeStrike struct {
		Strike  float64
		CallOI  float64
		PutOI   float64
		CallVol float64
		PutVol  float64
		Ts      okex.JSONTime
	}
	TakerFlow struct {
		CallBuyVol   float64
		CallSellVol  float64
		PutBuyVol    float64
		PutSellVol   float64
		CallBlockVol float64
		PutBlockVol  float64
		Ts           okex.JSONTime
	}
)

func (c *TakerVolume) UnmarshalJSON(buf []byte) error {
	var (
		sellVol, buyVol, ts string
		err                 error
	)
	tmp := []interface{}{&ts, &sellVol, &buyVol}
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

	c.SellVol, err = strconv.ParseFloat(sellVol, 64)
	if err != nil {
		return err
	}

	c.BuyVol, err = strconv.ParseFloat(buyVol, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *Ratio) UnmarshalJSON(buf []byte) error {
	var (
		ratio, ts string
		err       error
	)
	tmp := []interface{}{&ts, &ratio}
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

	c.Ratio, err = strconv.ParseFloat(ratio, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *InterestAndVolumeRatio) UnmarshalJSON(buf []byte) error {
	var (
		oi, vol, ts string
		err         error
	)
	tmp := []interface{}{&ts, &oi, &vol}
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

	if oi != "" {
		c.Oi, err = strconv.ParseFloat(oi, 64)
		if err != nil {
			return err
		}
	}
	if vol != "" {
		c.Vol, err = strconv.ParseFloat(vol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *PutCallRatio) UnmarshalJSON(buf []byte) error {
	var (
		oi, vol, ts string
		err         error
	)
	tmp := []interface{}{&ts, &oi, &vol}
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

	if oi != "" {
		c.OiRatio, err = strconv.ParseFloat(oi, 64)
		if err != nil {
			return err
		}
	}
	if vol != "" {
		c.VolRatio, err = strconv.ParseFloat(vol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *InterestAndVolumeExpiry) UnmarshalJSON(buf []byte) error {
	var (
		callOI, putOI, callVol, putVol, expTime, ts string
		err                                         error
	)
	tmp := []interface{}{&ts, &expTime, &callOI, &putOI, &callVol, &putVol}
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

	exp, err := time.Parse("20060102", expTime)
	if err != nil {
		return err
	}
	*(*time.Time)(&c.ExpTime) = exp

	if callOI != "" {
		c.CallOI, err = strconv.ParseFloat(callOI, 64)
		if err != nil {
			return err
		}
	}
	if putOI != "" {
		c.PutOI, err = strconv.ParseFloat(putOI, 64)
		if err != nil {
			return err
		}
	}
	if callVol != "" {
		c.CallVol, err = strconv.ParseFloat(callVol, 64)
		if err != nil {
			return err
		}
	}
	if putVol != "" {
		c.PutVol, err = strconv.ParseFloat(putVol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *InterestAndVolumeStrike) UnmarshalJSON(buf []byte) error {
	var (
		callOI, putOI, callVol, putVol, strike, ts string
		err                                        error
	)
	tmp := []interface{}{&ts, &strike, &callOI, &putOI, &callVol, &putVol}
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

	if callOI != "" {
		c.CallOI, err = strconv.ParseFloat(callOI, 64)
		if err != nil {
			return err
		}
	}
	if putOI != "" {
		c.PutOI, err = strconv.ParseFloat(putOI, 64)
		if err != nil {
			return err
		}
	}
	if callVol != "" {
		c.CallVol, err = strconv.ParseFloat(callVol, 64)
		if err != nil {
			return err
		}
	}
	if putVol != "" {
		c.PutVol, err = strconv.ParseFloat(putVol, 64)
		if err != nil {
			return err
		}
	}
	if strike != "" {
		c.Strike, err = strconv.ParseFloat(strike, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *TakerFlow) UnmarshalJSON(buf []byte) error {
	var (
		ts, callBuyVol, callSellVol, putBuyVol, putSellVol, callBlockVol, putBlockVol string
		err                                                                           error
	)
	tmp := []interface{}{&ts, &callBuyVol, &callSellVol, &putBuyVol, &putSellVol, &callBlockVol, &putBlockVol}
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

	if callBuyVol != "" {
		c.CallBlockVol, err = strconv.ParseFloat(callBuyVol, 64)
		if err != nil {
			return err
		}
	}
	if callSellVol != "" {
		c.CallSellVol, err = strconv.ParseFloat(callSellVol, 64)
		if err != nil {
			return err
		}
	}
	if putBuyVol != "" {
		c.PutBuyVol, err = strconv.ParseFloat(putBuyVol, 64)
		if err != nil {
			return err
		}
	}
	if putSellVol != "" {
		c.PutSellVol, err = strconv.ParseFloat(putSellVol, 64)
		if err != nil {
			return err
		}
	}
	if callBlockVol != "" {
		c.CallBuyVol, err = strconv.ParseFloat(callBlockVol, 64)
		if err != nil {
			return err
		}
	}
	if putBlockVol != "" {
		c.PutBuyVol, err = strconv.ParseFloat(putBlockVol, 64)
		if err != nil {
			return err
		}
	}

	return nil
}
