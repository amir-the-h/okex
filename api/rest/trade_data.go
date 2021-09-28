package rest

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/tradedata"
	responses "github.com/amir-the-h/okex/responses/trade_data"
	"net/http"
)

// TradeData
//
// https://www.okex.com/docs-v5/en/#rest-api-tradeing-data
type TradeData struct {
	client *ClientRest
}

// NewTradeData returns a pointer to a fresh TradeData
func NewTradeData(c *ClientRest) *TradeData {
	return &TradeData{c}
}

// GetSupportCoin
// Get the currency supported by the transaction big data interface
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-support-coin
func (c *TradeData) GetSupportCoin() (response responses.GetSupportCoin, err error) {
	p := "/api/v5/rubik/stat/trading-data/support-coin"
	res, err := c.client.Do(http.MethodGet, p, false)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetTakerVolume
// This is the taker volume for both buyers and sellers. This shows the influx and exit of funds in and out of {coin}.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-support-coin
func (c *TradeData) GetTakerVolume(req requests.GetTakerVolume) (response responses.GetTakerVolume, err error) {
	p := "/api/v5/rubik/stat/taker-volume"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMarginLendingRatio
// This indicator shows the ratio of cumulative data value between currency pair leverage quote currency and underlying asset over a given period of time.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-margin-lending-ratio
func (c *TradeData) GetMarginLendingRatio(req requests.GetRatio) (response responses.GetRatio, err error) {
	p := "/api/v5/rubik/stat/margin/loan-ratio"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetLongShortRatio
// This is the ratio of users with net long vs short positions. It includes data from futures and perpetual swaps.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-long-short-ratio
func (c *TradeData) GetLongShortRatio(req requests.GetRatio) (response responses.GetRatio, err error) {
	p := "/api/v5/rubik/stat/contracts/long-short-account-ratio"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetContractsOpenInterestAndVolume
// Open interest is the sum of all long and short futures and perpetual swap positions.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-contracts-open-interest-and-volume
func (c *TradeData) GetContractsOpenInterestAndVolume(req requests.GetRatio) (response responses.GetOpenInterestAndVolume, err error) {
	p := "/api/v5/rubik/stat/contracts/open-interest-volume"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetOptionsOpenInterestAndVolume
// This shows the sum of all open positions and how much total trading volume has taken place.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-options-open-interest-and-volume
func (c *TradeData) GetOptionsOpenInterestAndVolume(req requests.GetRatio) (response responses.GetOpenInterestAndVolume, err error) {
	p := "/api/v5/rubik/stat/option/open-interest-volume"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetPutCallRatio
// This shows the relative buy/sell volume for calls and puts. It shows whether traders are bullish or bearish on price and volatility.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-put-call-ratio
func (c *TradeData) GetPutCallRatio(req requests.GetRatio) (response responses.GetPutCallRatio, err error) {
	p := "/api/v5/rubik/stat/option/open-interest-volume-ratio"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetOpenInterestAndVolumeExpiry
// This shows the volume and open interest for each upcoming expiration. You can use this to see which expirations are currently the most popular to trade.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-open-interest-and-volume-expiry
func (c *TradeData) GetOpenInterestAndVolumeExpiry(req requests.GetRatio) (response responses.GetOpenInterestAndVolumeExpiry, err error) {
	p := "/api/v5/rubik/stat/option/open-interest-volume-expiry"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetOpenInterestAndVolumeStrike
// This shows what option strikes are the most popular for each expiration.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-open-interest-and-volume-strike
func (c *TradeData) GetOpenInterestAndVolumeStrike(req requests.GetOpenInterestAndVolumeStrike) (response responses.GetOpenInterestAndVolumeStrike, err error) {
	p := "/api/v5/rubik/stat/option/open-interest-volume-strike"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetTakerFlow
// This shows the relative buy/sell volume for calls and puts. It shows whether traders are bullish or bearish on price and volatility.
//
// https://www.okex.com/docs-v5/en/#rest-api-trading-data-get-taker-flow
func (c *TradeData) GetTakerFlow(req requests.GetRatio) (response responses.GetTakerFlow, err error) {
	p := "/api/v5/rubik/stat/option/taker-block-volume"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
