package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/responses"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type (
	// ClientWs is the websocket api client
	ClientWs struct {
		AuthorizedUntil *time.Time
		Cancel          context.CancelFunc
		DoneChan        chan interface{}
		EventChan       chan *responses.WsBasicEvent
		apiKey          string
		secretKey       []byte
		passphrase      string
		destination     okex.Destination
		url             okex.BaseUrl
		sendChan        chan []byte
		lastTransmit    *time.Time
		conn            *websocket.Conn
		mu              *sync.Mutex
		rmu             *sync.Mutex
		ctx             context.Context
		subChan         chan *responses.WsSubscribe
		uSubChan        chan *responses.WsUnsubscribe
		errChan         chan *responses.WsError
		processor       func(data []byte, e *responses.WsBasicEvent) bool
	}
)

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pongWait   = 30 * time.Second
	pingPeriod = (pongWait * 8) / 10
)

// NewClientWs returns a pointer to a fresh ClientWs
func NewClientWs(ctx context.Context, apiKey, secretKey, passphrase string, url okex.BaseUrl) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		apiKey:     apiKey,
		secretKey:  []byte(secretKey),
		passphrase: passphrase,
		ctx:        ctx,
		Cancel:     cancel,
		url:        url,
		sendChan:   make(chan []byte, 5),
		DoneChan:   make(chan interface{}),
		EventChan:  make(chan *responses.WsBasicEvent),
		mu:         &sync.Mutex{},
		rmu:        &sync.Mutex{},
	}

	return c
}

func (c *ClientWs) Process(data []byte, e *responses.WsBasicEvent) bool {
	if c.errChan != nil && e.Event == "error" {
		e := responses.WsError{}
		_ = json.Unmarshal(data, &e)
		go func() { c.errChan <- &e }()
		return true
	}
	if c.subChan != nil && e.Event == "subscribe" {
		e := responses.WsSubscribe{}
		_ = json.Unmarshal(data, &e)
		go func() { c.subChan <- &e }()
		return true
	}
	if c.uSubChan != nil && e.Event == "unsubscribe" {
		e := responses.WsUnsubscribe{}
		_ = json.Unmarshal(data, &e)
		go func() { c.uSubChan <- &e }()
		return true
	}

	return false
}

// Send message through either connections
func (c *ClientWs) Send(op okex.Operation, args []map[string]interface{}) error {
	err := c.Connect()
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"op":   op,
		"args": args,
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c.sendChan <- j

	return nil
}

// Connect into the server
//
// https://www.okex.com/docs-v5/en/#websocket-api-connect
func (c *ClientWs) Connect() error {
	if c.conn != nil {
		return nil
	}
	err := c.dial()
	if err == nil {
		return nil
	}
	ticker := time.NewTicker(redialTick)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err = c.dial()
			if err == nil {
				return nil
			}
		case <-c.ctx.Done():
			return c.handleCancel("connect")
		}
	}
}

// Subscribe into channel(s)
//
// https://www.okex.com/docs-v5/en/#websocket-api-subscribe
func (c *ClientWs) Subscribe(ch []okex.ChannelName, args map[string]interface{}) error {
	if len(ch) == 0 {
		return fmt.Errorf("not enough channels")
	}
	tmpArgs := make([]map[string]interface{}, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]interface{})
		tmpArgs[i]["channel"] = name
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(okex.SubscribeOperation, tmpArgs)
}

// Unsubscribe into channel(s)
//
// https://www.okex.com/docs-v5/en/#websocket-api-unsubscribe
func (c *ClientWs) Unsubscribe(ch []okex.ChannelName, args map[string]interface{}) error {
	if len(ch) == 0 {
		return fmt.Errorf("not enough channels")
	}
	tmpArgs := make([]map[string]interface{}, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]interface{})
		tmpArgs[i]["channel"] = name
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(okex.UnsubscribeOperation, tmpArgs)
}

// SetErrorHandler will set the error event handler
func (c *ClientWs) SetErrorHandler(ch chan *responses.WsError) {
	c.errChan = ch
}

// SetSubscribeHandler will set the error event handler
func (c *ClientWs) SetSubscribeHandler(ch chan *responses.WsSubscribe) {
	c.subChan = ch
}

// SetUnsubscribeHandler will set the error event handler
func (c *ClientWs) SetUnsubscribeHandler(ch chan *responses.WsUnsubscribe) {
	c.uSubChan = ch
}

func (c *ClientWs) Sign(method, path string) (string, string) {
	t := time.Now().UTC().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)

	// Get result and encode as hexadecimal string
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
func (c *ClientWs) dial() error {
	c.mu.Lock()
	c.rmu.Lock()
	defer func() {
		c.mu.Unlock()
		c.rmu.Unlock()
	}()
	conn, _, err := websocket.DefaultDialer.Dial(string(c.url), nil)
	if err != nil {
		return err
	}
	go func() {
		err := c.receiver()
		if err != nil {
			fmt.Printf("receiver error: %v\n", err)
			c.Cancel()
		}
	}()
	go func() {
		err := c.sender()
		if err != nil {
			fmt.Printf("sender error: %v\n", err)
			c.Cancel()
		}
	}()
	c.conn = conn

	return nil
}
func (c *ClientWs) sender() error {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for {
		select {
		case data := <-c.sendChan:
			c.mu.Lock()
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				c.mu.Unlock()
				return err
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.mu.Unlock()
				return err
			}

			if _, err = w.Write(data); err != nil {
				c.mu.Unlock()
				return err
			}

			c.mu.Unlock()
			now := time.Now()
			c.lastTransmit = &now
			if err := w.Close(); err != nil {
				return err
			}
		case <-ticker.C:
			if c.conn != nil && (c.lastTransmit == nil || (c.lastTransmit != nil && time.Since(*c.lastTransmit) > pingPeriod)) {
				c.sendChan <- []byte("ping")
			}
		case <-c.ctx.Done():
			return c.handleCancel("sender")
		}
	}
}
func (c *ClientWs) receiver() error {
	for {
		select {
		case <-c.ctx.Done():
			return c.handleCancel("receiver")
		default:
			c.rmu.Lock()
			err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
			if err != nil {
				c.rmu.Unlock()
				return err
			}

			mt, data, err := c.conn.ReadMessage()
			if err != nil {
				c.rmu.Unlock()
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					return c.conn.Close()
				}
				return err
			}

			c.rmu.Unlock()
			now := time.Now()
			c.lastTransmit = &now
			if mt == websocket.TextMessage && string(data) != "pong" {
				e := &responses.WsBasicEvent{}
				if err := json.Unmarshal(data, &e); err != nil {
					return err
				}
				if c.processor != nil {
					if c.processor(data, e) {
						continue
					}
				}
				c.EventChan <- e
			}
		}
	}
}
func (c *ClientWs) handleCancel(msg string) error {
	go func() {
		c.DoneChan <- msg
	}()
	return fmt.Errorf("operation cancelled: %s", msg)
}
