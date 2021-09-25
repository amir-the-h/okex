okex
====

*NOTICE:*
> PACKAGE IS CURRENTLY UNDER HEAVY DEVELOPMENT AND THERE IS NO GUARANTY FOR STABILITY UNTIL V1 RELEASE.

Okex V5 Golang API

A complete golang wrapper for [Okex](https://www.okex.com) V5 API. Pretty simple and easy to use. For more info about
Okex V5 API [read here](https://www.okex.com/docs-v5/en).

Installation
-----------------

```bash
go get github.com/amir-the-h/okex
```

Usage
-----------

```go
package main

import (
	"context"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	"github.com/amir-the-h/okex/events"
	"github.com/amir-the-h/okex/events/private"
	ws_requests "github.com/amir-the-h/okex/requests/ws/private"
	"log"
)

func main() {
	apiKey := "YOUR-API-KEY"
	secretKey := "YOUR-SECRET-KEY"
	passphrase := "YOUR-PASS-PHRASE"
	dest := okex.NormalServer
	ctx := context.Background()
	client, err := api.NewClient(ctx, apiKey, secretKey, passphrase, &dest)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Rest.Account.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Account Config %+v", response)

	errChan := make(chan *events.Error)
	subChan := make(chan *events.Subscribe)
	uSubChan := make(chan *events.Unsubscribe)
	lCh := make(chan *events.Login)
	oCh := make(chan *private.Order)
	log.Print("WS", "Logging in")
	client.Ws.Private.SetChannels(errChan, subChan, uSubChan)
	err = client.Ws.Login(lCh)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case <-lCh:
			log.Print("WS", "Logged in")
		case sub := <-subChan:
			channel, _ := sub.Arg.Get("channel")
			log.Printf("WS Subscribed on:\t%s", channel)
		case uSub := <-uSubChan:
			channel, _ := uSub.Arg.Get("channel")
			log.Printf("WS Unsubscribed from:\t%s", channel)
		case err := <-client.Ws.ErrChan:
			log.Printf("WS Error:\t%+v", err)
		case o := <-oCh:
			log.Printf("WS Event [Order]:\t %+v", o)
			for _, p := range o.Orders {
				log.Printf("\tOrder: \t%+v", p)
			}
			// Cancel the context
			// to stop all ops
			client.Ws.Cancel()

		case e := <-client.Ws.StructuredEventChan:
			log.Printf("WS Event [STRUCTED]:\t%+v", e)
		case e := <-client.Ws.RawEventChan:
			log.Printf("WS Event [RAW]:\t%+v", e)
		case b := <-client.Ws.DoneChan:
			log.Printf("WS End:\t%v", b)
			return
		}
	}
}
```

Supporting APIs
---------------

* [Rest](https://www.okex.com/docs-v5/en/#rest-api)
    * [Trade](https://www.okex.com/docs-v5/en/#rest-api-trade) (except demo special trading endpoints)
    * [Funding](https://www.okex.com/docs-v5/en/#rest-api-funding)
    * [Account](https://www.okex.com/docs-v5/en/#rest-api-account)
    * [SubAccount](https://www.okex.com/docs-v5/en/#rest-api-subaccount)
    * [Market Data](https://www.okex.com/docs-v5/en/#rest-api-market-data)
    * [Public Data](https://www.okex.com/docs-v5/en/#rest-api-public-data)
    * [Trading Data](https://www.okex.com/docs-v5/en/#rest-api-trading-data)

[comment]: <> (    * [Status]&#40;https://www.okex.com/docs-v5/en/#rest-api-status&#41;)

* [Ws](https://www.okex.com/docs-v5/en/#websocket-api)
    * [Private Channel](https://www.okex.com/docs-v5/en/#websocket-api-private-channel) (except demo special trading
      endpoints)

[comment]: <> (    * [Public Channel]&#40;https://www.okex.com/docs-v5/en/#websocket-api-public-channels&#41;)

[comment]: <> (    * [Trade]&#40;https://www.okex.com/docs-v5/en/#websocket-api-trade&#41;)

