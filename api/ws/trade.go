package ws

import (
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/trade"
	"time"
)

// Trade
//
// https://www.okex.com/docs-v5/en/#websocket-api-trade
type Trade struct {
	*ClientWs
}

// NewTrade returns a pointer to a fresh Trade
func NewTrade(c *ClientWs) *Trade {
	return &Trade{ClientWs: c}
}

// PlaceOrder
// You can place an order only if you have sufficient funds.
//
// https://www.okex.com/docs-v5/en/#websocket-api-trade-place-order
//
// Place orders in a batch. Maximum 20 orders can be placed at a time
//
// https://www.okex.com/docs-v5/en/#websocket-api-trade-place-multiple-orders
func (c *Trade) PlaceOrder(req ...requests.PlaceOrder) error {
	tmpArgs := make([]map[string]string, len(req))
	op := okex.OrderOperation
	if len(req) > 1 {
		op = okex.BatchOrderOperation
	}
	for i, order := range req {
		tmpArgs[i] = okex.S2M(order)
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()

	return c.Send(true, op, tmpArgs, map[string]string{"id": req[0].ID})
}

// CancelOrder
// Cancel an incomplete order
//
// https://www.okex.com/docs-v5/en/#websocket-api-trade-place-order
//
// Cancel incomplete orders in batches. Maximum 20 orders can be canceled at a time.
//
// https://www.okex.com/docs-v5/en/#websocket-api-trade-cancel-multiple-orders
func (c *Trade) CancelOrder(req ...requests.CancelOrder) error {
	tmpArgs := make([]map[string]string, len(req))
	op := okex.CancelOrderOperation
	if len(req) > 1 {
		op = okex.BatchCancelOrderOperation
	}
	for i, order := range req {
		tmpArgs[i] = okex.S2M(order)
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()

	return c.Send(true, op, tmpArgs, map[string]string{"id": req[0].ID})
}

// AmendOrder
// Amend an incomplete order.
//
// https://www.okex.com/docs-v5/en/#websocket-api-trade-place-order
//
// Amend incomplete orders in batches. Maximum 20 orders can be amended at a time.
//
// https://www.okex.com/docs-v5/en/#websocket-api-trade-amend-multiple-orders
func (c *Trade) AmendOrder(req ...requests.AmendOrder) error {
	tmpArgs := make([]map[string]string, len(req))
	op := okex.AmendOrderOperation
	if len(req) > 1 {
		op = okex.BatchAmendOrderOperation
	}
	for i, order := range req {
		tmpArgs[i] = okex.S2M(order)
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()

	return c.Send(true, op, tmpArgs, map[string]string{"id": req[0].ID})
}

func (c *Trade) waitForAuthorization() {
	if c.AuthorizedUntil != nil && time.Since(*c.AuthorizedUntil) < PingPeriod {
		return
	}
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for range ticker.C {
		if c.AuthorizedUntil != nil && time.Since(*c.AuthorizedUntil) < PingPeriod {
			return
		}
	}
}
