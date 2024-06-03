package main

import (
	"context"
	"crypto/tls"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	"github.com/amir-the-h/okex/events/public"
	requests "github.com/amir-the-h/okex/requests/ws/public"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	// Start the pprof server
	go func() {
		log.Println("Starting pprof server on localhost:6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalf("could not start pprof server: %v", err)
		}
	}()

	apiKey := ""
	secretKey := ""
	passphrase := ""
	ctx := context.Background()
	client, err := api.NewClient(ctx, apiKey, secretKey, passphrase, okex.NormalServer)
	if err != nil {
		log.Fatalln(err)
	}

	orderBookRequests := []requests.OrderBook{
		{InstID: "BTC-USDT", Channel: "books"},
		{InstID: "ETH-USDT", Channel: "books"},
		{InstID: "LTC-USDT", Channel: "books"},
		{InstID: "XRP-USDT", Channel: "books"},
		{InstID: "EOS-USDT", Channel: "books"},
		{InstID: "BCH-USDT", Channel: "books"},
		{InstID: "ETC-USDT", Channel: "books"},
		{InstID: "BSV-USDT", Channel: "books"},
		{InstID: "TRX-USDT", Channel: "books"},
		{InstID: "LINK-USDT", Channel: "books"},
		{InstID: "ADA-USDT", Channel: "books"},
		{InstID: "DOT-USDT", Channel: "books"},
		{InstID: "UNI-USDT", Channel: "books"},
	}

	client.Ws.Public.SetDialer(&websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	})
	obCh := make(chan *public.OrderBook)
	err = client.Ws.Public.OrderBook(orderBookRequests, obCh)
	if err != nil {
		log.Fatalln(err)
	}

	// Listen for updates
	for update := range obCh {
		log.Printf("Received order book update: %+v\n", update)
		insId, _ := update.Arg.Get("instId")
		log.Printf("Instrument ID: %s\n", insId)
	}
}
