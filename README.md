okex
====
[![Go Reference](https://pkg.go.dev/badge/github.com/amir-the-h/okex.svg)](https://pkg.go.dev/github.com/amir-the-h/okex)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/amir-the-h/okex.svg)](https://github.com/amir-the-h/okex)
[![GoReportCard example](https://goreportcard.com/badge/github.com/amir-the-h/okex)](https://goreportcard.com/report/github.com/amir-the-h/okex)
[![GitHub license](https://img.shields.io/github/license/amir-the-h/okex.svg)](https://github.com/amir-the-h/okex/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/release/amir-the-h/okex.svg)](https://GitHub.com/amir-the-h/okex/releases/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![CI](https://github.com/amir-the-h/okex/actions/workflows/main.yml/badge.svg)](https://github.com/amir-the-h/okex/actions/workflows/main.yml)
[![CodeQL](https://github.com/amir-the-h/okex/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/amir-the-h/okex/actions/workflows/codeql-analysis.yml)
[![AutoRelease](https://github.com/amir-the-h/okex/actions/workflows/release.yml/badge.svg)](https://github.com/amir-the-h/okex/actions/workflows/release.yml)

*NOTICE:*
> PACKAGE IS CURRENTLY UNDER HEAVY DEVELOPMENT AND THERE IS NO GUARANTY FOR STABILITY UNTIL V1 RELEASE.

Okex V5 Golang API

A complete golang wrapper for [Okex](https://www.okex.com) V5 API. Pretty simple and easy to use. For more info about
Okex V5 API [read here](https://www.okex.com/docs-v5/en).

Installation
-----------------

```bash
go get github.com/amir-the-h/okex@v1.0.1-alpha
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
	ws_private_requests "github.com/amir-the-h/okex/requests/ws/private"
	ws_public_requests "github.com/amir-the-h/okex/requests/ws/public"
	"log"
)

func main() {
	apiKey := "YOUR-API-KEY"
	secretKey := "YOUR-SECRET-KEY"
	passphrase := "YOUR-PASS-PHRASE"
	dest := okex.NormalServer // The main API server
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
	iCh := make(chan *public.Instruments)

	// to receive unique events individually in separated channels
	client.Ws.SetChannels(errChan, subChan, uSubChan, lCh)

	// subscribe into orders private channel
	// it will do the login process and wait until authorization confirmed
	err = client.Ws.Private.Order(ws_private_requests.Order{
		InstType: okex.SwapInstrument,
	}, oCh)
	if err != nil {
		log.Fatalln(err)
	}

	// subscribe into instruments public channel
	// it doesn't need any authorization
	err = client.Ws.Public.Instruments(ws_public_requests.Instruments{
		InstType: okex.SwapInstrument,
	}, iCh)
	if err != nil {
		log.Fatalln("Instruments", err)
	}

	// starting on listening 
	for {
		select {
		case <-lCh:
			log.Print("[Authorized]")
		case sub := <-subChan:
			channel, _ := sub.Arg.Get("channel")
			log.Printf("[Subscribed]\t%s", channel)
		case uSub := <-uSubChan:
			channel, _ := uSub.Arg.Get("channel")
			log.Printf("[Unsubscribed]\t%s", channel)
		case err := <-client.Ws.ErrChan:
			log.Printf("[Error]\t%+v", err)
		case o := <-oCh:
			log.Print("[Event]\tOrder")
			for _, p := range o.Orders {
				log.Printf("\t%+v", p)
			}
		case i := <-iCh:
			log.Print("[Event]\tInstrument")
			for _, p := range i.Instruments {
				log.Printf("\t%+v", p)
			}
		case e := <-client.Ws.StructuredEventChan:
			log.Printf("[Event] STRUCTED:\t%+v", e)
			v := reflect.TypeOf(e)
			switch v {
			case reflect.TypeOf(events.Error{}):
				log.Printf("[Error] STRUCTED:\t%+v", e)
			case reflect.TypeOf(events.Subscribe{}):
				log.Printf("[Subscribed] STRUCTED:\t%+v", e)
			case reflect.TypeOf(events.Unsubscribe{}):
				log.Printf("[Unsubscribed] STRUCTED:\t%+v", e)
			}
		case e := <-client.Ws.RawEventChan:
			log.Printf("[Event] RAW:\t%+v", e)
		case b := <-client.Ws.DoneChan:
			log.Printf("[End]:\t%v", b)
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
    * [Public Channel](https://www.okex.com/docs-v5/en/#websocket-api-public-channels)
    * [Trade](https://www.okex.com/docs-v5/en/#websocket-api-trade)

Features
--------

* All [requests](/requests), [responses](/responses), and [events](events) are well typed and will convert into the
  language built-in types instead of using API's strings. *Note that zero values will be replaced with non-existing
  data.*
* Fully automated authorization steps for both [REST](/api/rest) and [WS](/api/ws)
* To receive websocket events you can choose [RawEventChan](/api/ws/client.go#L25)
  , [StructuredEventChan](/api/ws/client.go#L28), or provide your own
  channels. [More info](https://github.com/amir-the-h/okex/wiki/Handling-WS-events) 
