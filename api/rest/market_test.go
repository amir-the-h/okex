package rest

import (
	"fmt"
	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/requests/rest/market"
	"testing"
)

func TestMarket_GetIndexTickers(t *testing.T) {
	m := Market{
		client: NewClient("", "", "", okex.RestURL, okex.NormalServer),
	}
	resp, err := m.GetTicker(market.GetTicker{
		InstID: "BTC-USD-SWAP",
	})
	fmt.Println(err)
	fmt.Println(resp.Tickers[0])
}
