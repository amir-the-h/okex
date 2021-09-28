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
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okex.Destination) (*Client, error) {
	restURL := okex.RestUrl
	wsPubURL := okex.PublicWsUrl
	wsPriURL := okex.PrivateWsUrl
	switch destination {
	case okex.AwsServer:
		restURL = okex.AwsRestUrl
		wsPubURL = okex.AwsPublicWsUrl
		wsPriURL = okex.AwsPrivateWsUrl
	case okex.DemoServer:
		restURL = okex.DemoRestUrl
		wsPubURL = okex.DemoPublicWsUrl
		wsPriURL = okex.DemoPrivateWsUrl
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex.BaseUrl{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}
