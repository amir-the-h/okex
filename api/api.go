package api

import (
	"context"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api/rest"
	"github.com/amir-the-h/okex/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest      *rest.ClientRest
	PrivateWs *ws.PrivateClient
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
	c1 := ws.NewClientWs(ctx, apiKey, secretKey, passphrase, wsPriUrl)
	//c2 := ws.NewClientWs(ctx, apiKey, secretKey, passphrase, wsPubUrl)
	priWs := ws.NewPrivateClient(c1)
	//pubWs := ws.NewPublicClient(c2)

	return &Client{r, priWs, ctx}, nil
}
