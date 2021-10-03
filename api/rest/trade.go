package rest

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/trade"
	responses "github.com/amir-the-h/okex/responses/trade"
	"net/http"
)

// Trade
//
// https://www.okex.com/docs-v5/en/#rest-api-trade
type Trade struct {
	client *ClientRest
}

// NewTrade returns a pointer to a fresh Trade
func NewTrade(c *ClientRest) *Trade {
	return &Trade{c}
}

// PlaceOrder
// You can place an order only if you have sufficient funds.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-get-positions
func (c *Trade) PlaceOrder(req []requests.PlaceOrder) (response responses.PlaceOrder, err error) {
	p := "/api/v5/trade/order"
	var tmp interface{}
	tmp = req[0]
	if len(req) > 1 {
		tmp = req
		p = "/api/trade/batch-orders"
	}
	m := okex.S2M(tmp)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// PlaceMultipleOrders
// Cancel an incomplete order.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-place-multiple-orders
func (c *Trade) PlaceMultipleOrders(req []requests.PlaceOrder) (response responses.PlaceOrder, err error) {
	p := "/api/v5/trade/batch-order"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// CandleOrder
// Cancel an incomplete order.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-cancel-order
//
// Cancel incomplete orders in batches. Maximum 20 orders can be canceled at a time. Request parameters should be passed in the form of an array.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-cancel-multiple-orders
func (c *Trade) CandleOrder(req []requests.CancelOrder) (response responses.PlaceOrder, err error) {
	p := "/api/v5/trade/cancel-order"
	var tmp interface{}
	tmp = req[0]
	if len(req) > 1 {
		tmp = req
		p = "/api/trade/cancel-batch-orders"
	}
	m := okex.S2M(tmp)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// AmendOrder
// Amend an incomplete order.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-amend-order
//
// Amend incomplete orders in batches. Maximum 20 orders can be amended at a time. Request parameters should be passed in the form of an array.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-amend-multiple-orders
func (c *Trade) AmendOrder(req []requests.OrderList) (response responses.AmendOrder, err error) {
	p := "/api/v5/trade/amend-order"
	var tmp interface{}
	tmp = req[0]
	if len(req) > 1 {
		tmp = req
		p = "/api/trade/amend-batch-orders"
	}
	m := okex.S2M(tmp)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// ClosePosition
// Close all positions of an instrument via a market order.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-close-positions
func (c *Trade) ClosePosition(req requests.ClosePosition) (response responses.ClosePosition, err error) {
	p := "/api/v5/trade/close-position"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetOrderDetail
// Retrieve order details.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-get-order-details
func (c *Trade) GetOrderDetail(req requests.OrderList) (response responses.Order, err error) {
	p := "/api/v5/trade/order"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetOrderList
// Retrieve all incomplete orders under the current account.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-get-order-list
func (c *Trade) GetOrderList(req requests.OrderList) (response responses.Order, err error) {
	p := "/api/v5/trade/orders-pending"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetOrderHistory
// Retrieve the completed order data for the last 7 days, and the incomplete orders that have been cancelled are only reserved for 2 hours.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-get-order-history-last-7-days
//
// Retrieve the completed order data of the last 3 months, and the incomplete orders that have been canceled are only reserved for 2 hours.
// https://www.okex.com/docs-v5/en/#rest-api-trade-get-order-history-last-3-months
func (c *Trade) GetOrderHistory(req requests.OrderList, arch bool) (response responses.Order, err error) {
	p := "/api/v5/trade/orders-history"
	if arch {
		p = "/api/trade/orders-history-archive"
	}
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

// GetTransactionDetails
// Retrieve recently-filled transaction details in the last 3 day.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-get-order-history-last-7-days
//
// Retrieve recently-filled transaction details in the last 3 months.
//
// https://www.okex.com/docs-v5/en/#rest-api-trade-get-transaction-details-last-3-months
func (c *Trade) GetTransactionDetails(req requests.TransactionDetails, arch bool) (response responses.TransactionDetail, err error) {
	p := "/api/v5/trade/fills"
	if arch {
		p = "/api/trade/fills-history"
	}
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
