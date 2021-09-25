package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/events"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

// ClientWs is the websocket api client
//
// https://www.okex.com/docs-v5/en/#websocket-api
type ClientWs struct {
	AuthorizedUntil     *time.Time
	Cancel              context.CancelFunc
	DoneChan            chan interface{}
	RawEventChan        chan *events.Basic
	ErrChan             chan *events.Error
	StructuredEventChan chan interface{}
	Private             *Private
	lCh                 chan *events.Login
	SubscribeChan       chan *events.Subscribe
	UnsubscribeCh       chan *events.Unsubscribe
	sendChan            chan []byte
	apiKey              string
	secretKey           []byte
	passphrase          string
	destination         okex.Destination
	url                 okex.BaseUrl
	lastTransmit        *time.Time
	conn                *websocket.Conn
	mu                  *sync.Mutex
	rmu                 *sync.Mutex
	ctx                 context.Context
}

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pongWait   = 30 * time.Second
	pingPeriod = (pongWait * 8) / 10
)

// NewClient returns a pointer to a fresh ClientWs
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, url okex.BaseUrl) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		apiKey:              apiKey,
		secretKey:           []byte(secretKey),
		passphrase:          passphrase,
		ctx:                 ctx,
		Cancel:              cancel,
		url:                 url,
		sendChan:            make(chan []byte, 5),
		DoneChan:            make(chan interface{}),
		RawEventChan:        make(chan *events.Basic),
		StructuredEventChan: make(chan interface{}),
		mu:                  &sync.Mutex{},
		rmu:                 &sync.Mutex{},
	}
	c.Private = NewPrivate(c)

	return c
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
			log.Println(err)
		case <-c.ctx.Done():
			return c.handleCancel("connect")
		}
	}
}

// Login
//
// https://www.okex.com/docs-v5/en/#websocket-api-login
func (c *ClientWs) Login(ch ...chan *events.Login) error {
	if c.AuthorizedUntil != nil && time.Since(*c.AuthorizedUntil) < pingPeriod {
		return nil
	}
	method := http.MethodGet
	path := "/users/self/verify"
	ts, sign := c.sign(method, path)
	args := []map[string]string{
		{
			"apiKey":     c.apiKey,
			"passphrase": c.passphrase,
			"timestamp":  ts,
			"sign":       sign,
		},
	}
	if len(ch) > 0 {
		c.lCh = ch[0]
	}

	return c.Send(okex.LoginOperation, args)
}

// Subscribe
// Users can choose to subscribe to one or more channels, and the total length of multiple channels cannot exceed 4096 bytes.
//
// https://www.okex.com/docs-v5/en/#websocket-api-subscribe
func (c *ClientWs) Subscribe(ch []okex.ChannelName, args map[string]string) error {
	tmpArgs := make([]map[string]string, len(ch))
	for i, name := range ch {
		tmpArgs[i] = map[string]string{}
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(okex.SubscribeOperation, tmpArgs)
}

// Unsubscribe into channel(s)
//
// https://www.okex.com/docs-v5/en/#websocket-api-unsubscribe
func (c *ClientWs) Unsubscribe(ch []okex.ChannelName, args map[string]string) error {
	tmpArgs := make([]map[string]string, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]string)
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}

	return c.Send(okex.UnsubscribeOperation, tmpArgs)
}

// Send message through either connections
func (c *ClientWs) Send(op okex.Operation, args []map[string]string) error {
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

// process the incoming event
func (c *ClientWs) process(data []byte, e *events.Basic) bool {
	switch e.Event {
	case "error":
		e := events.Error{}
		_ = json.Unmarshal(data, &e)
		go func() {
			c.ErrChan <- &e
		}()
		return true
	case "subscribe":
		e := events.Subscribe{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.SubscribeChan != nil {
				c.SubscribeChan <- &e
			}
			c.StructuredEventChan <- e
		}()
		return true
	case "unsubscribe":
		e := events.Unsubscribe{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.UnsubscribeCh != nil {
				c.UnsubscribeCh <- &e
			}
			c.StructuredEventChan <- e
		}()
		return true
	case "login":
		au := time.Now().Add(time.Second * 30)
		c.AuthorizedUntil = &au
		e := events.Login{}
		_ = json.Unmarshal(data, &e)
		go func() {
			if c.lCh != nil {
				c.lCh <- &e
			}
			c.StructuredEventChan <- e
		}()
		return true
	default:
		if c.Private.Process(data, e) {
			return true
		}
	}

	go func() { c.RawEventChan <- e }()

	return false
}

// SetChannels to receive certain events on separate channel
func (c *ClientWs) SetChannels(errCh chan *events.Error, subCh chan *events.Subscribe, unSub chan *events.Unsubscribe) {
	c.ErrChan = errCh
	c.SubscribeChan = subCh
	c.UnsubscribeCh = unSub
}

func (c *ClientWs) dial() error {
	c.mu.Lock()
	c.rmu.Lock()
	defer func() {
		c.mu.Unlock()
		c.rmu.Unlock()
	}()
	conn, res, err := websocket.DefaultDialer.Dial(string(c.url), nil)
	if err != nil {
		return fmt.Errorf("error %d: %s", res.StatusCode, err)
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
				e := &events.Basic{}
				if err := json.Unmarshal(data, &e); err != nil {
					return err
				}
				go func() {
					c.process(data, e)
				}()
			}
		}
	}
}
func (c *ClientWs) sign(method, path string) (string, string) {
	t := time.Now().UTC().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)

	// Get result and encode as hexadecimal string
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
func (c *ClientWs) handleCancel(msg string) error {
	go func() {
		c.DoneChan <- msg
	}()
	return fmt.Errorf("operation cancelled: %s", msg)
}
