package api

import (
	"context"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api/rest"
	"github.com/amir-the-h/okex/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest *rest.ClientRest
	Ws   struct {
		*ws.Private
	}
	//PublicWs  *ws.PublicClient
	ctx context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination *okex.Destination) (*Client, error) {
	restUrl := okex.RestUrl
	//wsPubUrl := okex.PublicWsUrl
	wsPriUrl := okex.PrivateWsUrl
	switch *destination {
	case okex.AwsServer:
		restUrl = okex.AwsRestUrl
		//wsPubUrl = okex.AwsPublicWsUrl
		wsPriUrl = okex.AwsPrivateWsUrl
	case okex.DemoServer:
		restUrl = okex.DemoRestUrl
		//wsPubUrl = okex.DemoPublicWsUrl
		wsPriUrl = okex.DemoPrivateWsUrl
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restUrl, *destination)
	priC := ws.NewClient(ctx, apiKey, secretKey, passphrase, wsPriUrl)
	//pubC := ws.NewClient(ctx, apiKey, secretKey, passphrase, wsPubUrl)
	priWs := ws.NewPrivate(priC)
	w := struct {
		*ws.Private
	}{priWs}

	return &Client{r, w, ctx}, nil
}
