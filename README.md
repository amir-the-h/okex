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

Basic Usage
-----------

```go
package main

import (
	"context"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
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
	log.Printf("%+v", response)
}
```

see [examples](/examples)

Supporting APIs
---------------

- [x] [Trade](https://www.okex.com/docs-v5/en/#rest-api-trade) (except demo special trading endpoints)
- [x] [Funding](https://www.okex.com/docs-v5/en/#rest-api-funding)
- [x] [Account](https://www.okex.com/docs-v5/en/#rest-api-account)
- [x] [SubAccount](https://www.okex.com/docs-v5/en/#rest-api-subaccount)
- [x] [Market Data](https://www.okex.com/docs-v5/en/#rest-api-market-data)
- [x] [Public Data](https://www.okex.com/docs-v5/en/#rest-api-public-data)
- [ ] [Trading Data](https://www.okex.com/docs-v5/en/#rest-api-trading-data)
- [ ] [Status](https://www.okex.com/docs-v5/en/#rest-api-statusl)

