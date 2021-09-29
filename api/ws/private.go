package ws

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/events"
	"github.com/amir-the-h/okex/events/private"
	requests "github.com/amir-the-h/okex/requests/ws/private"
	"time"
)

// Private
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel
type Private struct {
	*ClientWs
	aCh   chan *private.Account
	pCh   chan *private.Position
	bnpCh chan *private.BalanceAndPosition
	oCh   chan *private.Order
}

// NewPrivate returns a pointer to a fresh Private
func NewPrivate(c *ClientWs) *Private {
	return &Private{ClientWs: c}
}

// Account
// Retrieve account information. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (c *Private) Account(req requests.Account, ch ...chan *private.Account) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.aCh = ch[0]
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()

	return c.Subscribe(true, []okex.ChannelName{"account"}, m)
}

// UAccount
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (c *Private) UAccount(req requests.Account, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.aCh = nil
	}

	return c.Unsubscribe(true, []okex.ChannelName{"account"}, m)
}

// Position
// Retrieve position information. Initial snapshot will be pushed according to subscription granularity. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (c *Private) Position(req requests.Position, ch ...chan *private.Position) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.pCh = ch[0]
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()
	return c.Subscribe(true, []okex.ChannelName{"positions"}, m)
}

// UPosition
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (c *Private) UPosition(req requests.Position, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.pCh = nil
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()
	return c.Unsubscribe(true, []okex.ChannelName{"positions"}, m)
}

// BalanceAndPosition
// Retrieve account balance and position information. Data will be pushed when triggered by events such as filled order, funding transfer.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (c *Private) BalanceAndPosition(ch ...chan *private.BalanceAndPosition) error {
	m := make(map[string]string)
	if len(ch) > 0 {
		c.bnpCh = ch[0]
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()
	return c.Subscribe(true, []okex.ChannelName{"balance_and_position"}, m)
}

// UBalanceAndPosition unsubscribes a position channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (c *Private) UBalanceAndPosition(rCh ...bool) error {
	m := make(map[string]string)
	if len(rCh) > 0 && rCh[0] {
		c.bnpCh = nil
	}

	return c.Unsubscribe(true, []okex.ChannelName{"balance_and_position"}, m)
}

// Order
// Retrieve position information. Initial snapshot will be pushed according to subscription granularity. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-order-channel
func (c *Private) Order(req requests.Order, ch ...chan *private.Order) error {
	m := okex.S2M(req)
	if len(ch) > 0 {
		c.oCh = ch[0]
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()
	return c.Subscribe(true, []okex.ChannelName{"orders"}, m)
}

// UOrder
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-order-channel
func (c *Private) UOrder(req requests.Order, rCh ...bool) error {
	m := okex.S2M(req)
	if len(rCh) > 0 && rCh[0] {
		c.oCh = nil
	}

	if err := c.Login(); err != nil {
		return err
	}
	c.waitForAuthorization()
	return c.Unsubscribe(true, []okex.ChannelName{"orders"}, m)
}

func (c *Private) Process(data []byte, e *events.Basic) bool {
	if e.Event == "" && e.Arg != nil && e.Data != nil && len(e.Data) > 0 {
		ch, ok := e.Arg.Get("channel")
		if !ok {
			return false
		}
		switch ch {
		case "account":
			e := private.Account{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.aCh != nil {
					c.aCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		case "positions":
			e := private.Position{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.pCh != nil {
					c.pCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		case "balance_and_position":
			e := private.BalanceAndPosition{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.bnpCh != nil {
					c.bnpCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		case "orders":
			e := private.Order{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.oCh != nil {
					c.oCh <- &e
				}
				c.StructuredEventChan <- e
			}()
			return true
		}
	}

	return false
}

func (c *Private) waitForAuthorization() {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for range ticker.C {
		if c.AuthorizedUntil != nil && time.Since(*c.AuthorizedUntil) < PingPeriod {
			return
		}
	}
}
