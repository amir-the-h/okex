package ws

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api/ws/private"
	"github.com/amir-the-h/okex/responses"
	"net/http"
	"time"
)

// PrivateClient is the private websocket api client
type PrivateClient struct {
	*ClientWs
	lChan   chan *private.Login
	aChan   chan *private.Account
	pChan   chan *private.Position
	bnpChan chan *private.BalanceAndPosition
}

// NewPrivateClient returns a pointer to a fresh PrivateClient
func NewPrivateClient(c *ClientWs) *PrivateClient {
	p := &PrivateClient{ClientWs: c}
	c.processor = p.process

	return p
}

// Login into the private server
//
// https://www.okex.com/docs-v5/en/#websocket-api-login
func (c *PrivateClient) Login(loginChan chan *private.Login) error {
	if c.AuthorizedUntil != nil && time.Since(*c.AuthorizedUntil) < pingPeriod {
		return nil
	}
	method := http.MethodGet
	path := "/users/self/verify"
	ts, sign := c.Sign(method, path)
	args := []map[string]interface{}{
		{
			"apiKey":     c.apiKey,
			"passphrase": c.passphrase,
			"timestamp":  ts,
			"sign":       sign,
		},
	}
	c.lChan = loginChan

	return c.Send(okex.LoginOperation, args)
}

// Account subscribes on an account channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (c *PrivateClient) Account(accountChan chan *private.Account, ccy string) error {
	args := make(map[string]interface{})
	if ccy != "" {
		args["ccy"] = ccy
	}
	c.aChan = accountChan

	return c.Subscribe([]okex.ChannelName{"account"}, args)
}

// UAccount unsubscribes an account channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-account-channel
func (c *PrivateClient) UAccount(ccy string) error {
	args := make(map[string]interface{})
	if ccy != "" {
		args["ccy"] = ccy
	}
	c.aChan = nil

	return c.Unsubscribe([]okex.ChannelName{"account"}, args)
}

// Position subscribes on a position channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (c *PrivateClient) Position(positionChan chan *private.Position, instType okex.InstrumentType, uly, instId string) error {
	args := map[string]interface{}{"instType": instType}
	if uly != "" {
		args["uly"] = uly
	}
	if instId != "" {
		args["instId"] = instId
	}
	c.pChan = positionChan

	return c.Subscribe([]okex.ChannelName{"positions"}, args)
}

// UPosition unsubscribes a position channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-positions-channel
func (c *PrivateClient) UPosition(instType okex.InstrumentType, uly, instId string) error {
	args := map[string]interface{}{"instType": instType}
	if uly != "" {
		args["uly"] = uly
	}
	if instId != "" {
		args["instId"] = instId
	}
	c.pChan = nil

	return c.Subscribe([]okex.ChannelName{"positions"}, args)
}

// BalanceAndPosition subscribes on a position channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (c *PrivateClient) BalanceAndPosition(bnpChan chan *private.BalanceAndPosition) error {
	args := map[string]interface{}{}
	c.bnpChan = bnpChan

	return c.Subscribe([]okex.ChannelName{"balance_and_position"}, args)
}

// UBalanceAndPosition unsubscribes a position channel
//
// https://www.okex.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
func (c *PrivateClient) UBalanceAndPosition() error {
	args := map[string]interface{}{}
	c.bnpChan = nil

	return c.Unsubscribe([]okex.ChannelName{"balance_and_position"}, args)
}

func (c *PrivateClient) process(data []byte, e *responses.WsBasicEvent) bool {
	if c.Process(data, e) {
		return true
	}
	if e.Event == "subscribe" || e.Event == "unsubscribe" {
		return false
	}
	if e.Event == "login" {
		au := time.Now().Add(time.Second * 30)
		c.AuthorizedUntil = &au
		if c.lChan != nil {
			r := private.Login{}
			_ = json.Unmarshal(data, &r)
			go func() { c.lChan <- &r }()
			return true
		}
	}

	if c.aChan != nil && e.Arg != nil {
		ch, ok := e.Arg.Get("channel")
		if ok && ch == "account" {
			r := private.Account{}
			_ = json.Unmarshal(data, &r)
			go func() { c.aChan <- &r }()
			return true
		}
	}

	if c.pChan != nil && e.Arg != nil {
		ch, ok := e.Arg.Get("channel")
		if ok && ch == "positions" {
			r := private.Position{}
			_ = json.Unmarshal(data, &r)
			go func() { c.pChan <- &r }()
			return true
		}
	}

	if c.bnpChan != nil && e.Arg != nil {
		ch, ok := e.Arg.Get("channel")
		if ok && ch == "balance_and_position" {
			r := private.BalanceAndPosition{}
			_ = json.Unmarshal(data, &r)
			go func() { c.bnpChan <- &r }()
			return true
		}
	}

	return false
}
