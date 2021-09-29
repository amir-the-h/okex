package rest

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/subaccount"
	responses "github.com/amir-the-h/okex/responses/sub_account"
	"net/http"
	"strings"
)

// SubAccount
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount
type SubAccount struct {
	client *ClientRest
}

// NewSubAccount returns a pointer to a fresh SubAccount
func NewSubAccount(c *ClientRest) *SubAccount {
	return &SubAccount{c}
}

// ViewList
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-view-sub-account-list
func (c *SubAccount) ViewList(req requests.ViewList) (response responses.ViewList, err error) {
	p := "/api/v5/users/subaccount/list"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// CreateAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-create-an-apikey-for-a-sub-account
func (c *SubAccount) CreateAPIKey(req requests.CreateAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/apikey"
	m := okex.S2M(req)
	if len(req.IP) > 0 {
		m["ip"] = strings.Join(req.IP, ",")
	}
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// QueryAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-query-the-apikey-of-a-sub-account
func (c *SubAccount) QueryAPIKey(req requests.QueryAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// ResetAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-reset-the-apikey-of-a-sub-account
func (c *SubAccount) ResetAPIKey(req requests.CreateAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/modify-apikey"
	m := okex.S2M(req)
	if len(req.IP) > 0 {
		m["ip"] = strings.Join(req.IP, ",")
	}
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// DeleteAPIKey
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-delete-the-apikey-of-sub-accounts
func (c *SubAccount) DeleteAPIKey(req requests.DeleteAPIKey) (response responses.APIKey, err error) {
	p := "/api/v5/users/subaccount/delete-apikey"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetBalance
// Query detailed balance info of Trading Account of a sub-account via the master account
// (applies to master accounts only)
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-get-sub-account-balance
func (c *SubAccount) GetBalance(req requests.GetBalance) (response responses.GetBalance, err error) {
	p := "/api/v5/account/subaccount/balances"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// HistoryTransfer
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-history-of-sub-account-transfer
func (c *SubAccount) HistoryTransfer(req requests.HistoryTransfer) (response responses.HistoryTransfer, err error) {
	p := "/api/v5/account/subaccount/bills"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// ManageTransfers
// applies to master accounts only
//
// https://www.okex.com/docs-v5/en/#rest-api-subaccount-master-accounts-manage-the-transfers-between-sub-accounts
func (c *SubAccount) ManageTransfers(req requests.ManageTransfers) (response responses.ManageTransfer, err error) {
	p := "/api/v5/account/subaccount/transfer"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
