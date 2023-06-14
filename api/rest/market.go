package rest

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/market"
	responses "github.com/amir-the-h/okex/responses/market"
	"net/http"
)

// Market
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data
type Market struct {
	client *ClientRest
}

// NewMarket returns a pointer to a fresh Market
func NewMarket(c *ClientRest) *Market {
	return &Market{c}
}

// GetTickers
// Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-tickers
func (c *Market) GetTickers(req requests.GetTickers) (response responses.Ticker, err error) {
	p := "/api/v5/market/tickers"
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

// GetTicker
// Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-ticker
func (c *Market) GetTicker(req requests.GetTicker) (response responses.Ticker, err error) {
	p := "/api/v5/market/ticker"
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

// GetIndexTickers
// Retrieve index tickers.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-index-tickers
func (c *Market) GetIndexTickers(req requests.GetIndexTickers) (response responses.Ticker, err error) {
	p := "/api/v5/market/ticker"
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

// GetOrderBook
// Retrieve a instrument is order book.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-order-book
func (c *Market) GetOrderBook(req requests.GetOrderBook) (response responses.OrderBook, err error) {
	p := "/api/v5/market/books"
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

// GetCandlesticks
// Retrieve the candlestick charts. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-candlesticks
func (c *Market) GetCandlesticks(req requests.GetCandlesticks) (response responses.Candle, err error) {
	p := "/api/v5/market/candles"
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

// GetCandlesticksHistory
// Retrieve history candlestick charts from recent years.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-candlesticks
func (c *Market) GetCandlesticksHistory(req requests.GetCandlesticks) (response responses.Candle, err error) {
	p := "/api/v5/market/history-candles"
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

// GetIndexCandlesticks
// Retrieve the candlestick charts of the index. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-index-candlesticks
func (c *Market) GetIndexCandlesticks(req requests.GetCandlesticks) (response responses.IndexCandle, err error) {
	p := "/api/v5/market/index-candles"
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

// GetMarkPriceCandlesticks
// Retrieve the candlestick charts of mark price. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-mark-price-candlesticks
func (c *Market) GetMarkPriceCandlesticks(req requests.GetCandlesticks) (response responses.CandleMarket, err error) {
	p := "/api/v5/market/mark-price-candles"
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

// GetTrades
// Retrieve the recent transactions of an instrument.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-trades
func (c *Market) GetTrades(req requests.GetTrades) (response responses.Trade, err error) {
	p := "/api/v5/market/trades"
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

// Get24HTotalVolume
// The 24-hour trading volume is calculated on a rolling basis, using USD as the pricing unit.
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-24h-total-volume
func (c *Market) Get24HTotalVolume() (response responses.TotalVolume24H, err error) {
	p := "/api/v5/market/platform-24-volume"
	res, err := c.client.Do(http.MethodGet, p, false)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetIndexComponents
// Get the index component information data on the market
//
// https://www.okex.com/docs-v5/en/#rest-api-market-data-get-index-components
func (c *Market) GetIndexComponents(req requests.GetIndexComponents) (response responses.IndexComponent, err error) {
	p := "/api/v5/market/index-components"
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
