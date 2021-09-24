package rest

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/account"
	responses "github.com/amir-the-h/okex/responses/account"
	"net/http"
)

// Account
//
// https://www.okex.com/docs-v5/en/#rest-api-account
type Account struct {
	client *ClientRest
}

// NewAccount returns a pointer to a fresh Account
func NewAccount(c *ClientRest) *Account {
	return &Account{c}
}

// GetBalance
// Retrieve a list of assets (with non-zero balance), remaining balance, and available amount in the account.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-balance
func (a *Account) GetBalance(req requests.GetBalance) (response responses.GetBalance, err error) {
	p := "/api/v5/account/balance"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetPositions
// Retrieve information on your positions. When the account is in net mode, net positions will be displayed, and when the account is in long/short mode, long or short positions will be displayed.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-positions
func (a *Account) GetPositions(req requests.GetPositions) (response responses.GetPositions, err error) {
	p := "/api/v5/account/positions"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetAccountAndPositionRisk
// Get account and position risk
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-account-and-position-risk
func (a *Account) GetAccountAndPositionRisk(req requests.GetAccountAndPositionRisk) (response responses.GetAccountAndPositionRisk, err error) {
	p := "/api/v5/account/positions"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetBills
// Retrieve the bills of the account. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with the most recent first. This endpoint can retrieve data from the last 7 days.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-bills-details-last-7-days
//
// Retrieve the account’s bills. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with most recent first. This endpoint can retrieve data from the last 3 months.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-bills-details-last-3-months
func (a *Account) GetBills(req requests.GetBills, arc bool) (response responses.GetBills, err error) {
	p := "/api/v5/account/bills"
	if arc {
		p = "/api/account/bills-archive"
	}
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetConfig
// Retrieve current account configuration.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-account-configuration
func (a *Account) GetConfig() (response responses.GetConfig, err error) {
	p := "/api/v5/account/config"
	res, err := a.client.Do(http.MethodGet, p, true)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// SetPositionMode
// FUTURES and SWAP support both long/short mode and net mode. In net mode, users can only have positions in one direction; In long/short mode, users can hold positions in long and short directions.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-set-position-mode
func (a *Account) SetPositionMode(req requests.SetPositionMode) (response responses.SetPositionMode, err error) {
	p := "/api/v5/account/set-position-mode"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// SetLeverage
// The following are the setting leverage cases for an instrument:
//
// Set leverage for isolated MARGIN at pairs level.
//
// Set leverage for cross MARGIN in Single-currency margin at pairs level.
//
// Set leverage for cross MARGIN in Multi-currency margin at currency level.
//
// Set leverage for cross/isolated FUTURES/SWAP at underlying/contract level.
// https://www.okex.com/docs-v5/en/#rest-api-account-set-leverage
func (a *Account) SetLeverage(req requests.SetLeverage) (response responses.Leverage, err error) {
	p := "/api/v5/account/set-leverage"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxBuySellAmount
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-maximum-buy-sell-amount-or-open-amount
func (a *Account) GetMaxBuySellAmount(req requests.GetMaxBuySellAmount) (response responses.GetMaxBuySellAmount, err error) {
	p := "/api/v5/account/max-size"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxAvailableTradeAmount
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-maximum-available-tradable-amount
func (a *Account) GetMaxAvailableTradeAmount(req requests.GetMaxAvailableTradeAmount) (response responses.GetMaxAvailableTradeAmount, err error) {
	p := "/api/v5/account/max-avail-size"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// IncreaseDecreaseMargin
// Increase or decrease the margin of the isolated position.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-increase-decrease-margin
func (a *Account) IncreaseDecreaseMargin(req requests.IncreaseDecreaseMargin) (response responses.IncreaseDecreaseMargin, err error) {
	p := "/api/v5/account/position/margin-balance"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetLeverage
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-leverage
func (a *Account) GetLeverage(req requests.GetLeverage) (response responses.Leverage, err error) {
	p := "/api/v5/account/leverage-info"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxLoan
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-the-maximum-loan-of-instrument
func (a *Account) GetMaxLoan(req requests.GetMaxLoan) (response responses.GetMaxLoan, err error) {
	p := "/api/v5/account/max-loan"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetFeeRates
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-fee-rates
func (a *Account) GetFeeRates(req requests.GetFeeRates) (response responses.GetFeeRates, err error) {
	p := "/api/v5/account/trade-fee"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetInterestAccrued
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-interest-accrued
func (a *Account) GetInterestAccrued(req requests.GetInterestAccrued) (response responses.GetInterestAccrued, err error) {
	p := "/api/v5/account/interest-accrued"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetInterestRates
// Get the user's current leveraged currency borrowing interest rate
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-interest-rate
func (a *Account) GetInterestRates(req requests.GetBalance) (response responses.GetInterestRates, err error) {
	p := "/api/v5/account/interest-rate"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// SetGreeks
// Set the display type of Greeks.
//
// https://www.okex.com/docs-v5/en/#rest-api-account-set-greeks-m-bs
func (a *Account) SetGreeks(req requests.SetGreeks) (response responses.SetGreeks, err error) {
	p := "/api/v5/account/set-greeks"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMaxWithdrawals
//
// https://www.okex.com/docs-v5/en/#rest-api-account-get-maximum-withdrawals
func (a *Account) GetMaxWithdrawals(req requests.GetBalance) (response responses.GetMaxWithdrawals, err error) {
	p := "/api/v5/account/max-withdrawal"
	m := okex.S2M(req)
	res, err := a.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
