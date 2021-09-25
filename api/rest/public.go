package rest

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/public"
	responses "github.com/amir-the-h/okex/responses/public"
	"net/http"
)

// Public
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data
type Public struct {
	client *ClientRest
}

// NewPublic returns a pointer to a fresh Public
func NewPublic(c *ClientRest) *Public {
	return &Public{c}
}

// GetInstruments
// Retrieve a list of instruments with open contracts.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-instruments
func (c *Public) GetInstruments(req requests.GetInstruments) (response responses.GetInstruments, err error) {
	p := "/api/v5/public/instruments"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetDeliveryExerciseHistory
// Retrieve the estimated delivery price, which will only have a return value one hour before the delivery/exercise.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-instruments
func (c *Public) GetDeliveryExerciseHistory(req requests.GetDeliveryExerciseHistory) (response responses.GetDeliveryExerciseHistory, err error) {
	p := "/api/v5/public/delivery-exercise-history"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetOpenInterest
// Retrieve the total open interest for contracts on OKEx.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-open-interest
func (c *Public) GetOpenInterest(req requests.GetOpenInterest) (response responses.GetOpenInterest, err error) {
	p := "/api/v5/public/open-interest"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetLimitPrice
// Retrieve the highest buy limit and lowest sell limit of the instrument.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-limit-price
func (c *Public) GetLimitPrice(req requests.GetLimitPrice) (response responses.GetLimitPrice, err error) {
	p := "/api/v5/public/price-limit"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetOptionMarketData
// Retrieve option market data.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-option-market-data
func (c *Public) GetOptionMarketData(req requests.GetOptionMarketData) (response responses.GetOptionMarketData, err error) {
	p := "/api/v5/public/opt-summary"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetEstimatedDeliveryExcercisePrice
// Retrieve the estimated delivery price which will only have a return value one hour before the delivery/exercise.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-estimated-delivery-excercise-price
func (c *Public) GetEstimatedDeliveryExcercisePrice(req requests.GetEstimatedDeliveryExcercisePrice) (response responses.GetEstimatedDeliveryExcercisePrice, err error) {
	p := "/api/v5/public/estimated-price"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetDiscountRateAndInterestFreeQuota
// Retrieve discount rate level and interest-free quota.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-discount-rate-and-interest-free-quota
func (c *Public) GetDiscountRateAndInterestFreeQuota(req requests.GetDiscountRateAndInterestFreeQuota) (response responses.GetDiscountRateAndInterestFreeQuota, err error) {
	p := "/api/v5/public/discount-rate-interest-free-quota"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetSystemTime
// Retrieve API server time.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-system-time
func (c *Public) GetSystemTime() (response responses.GetSystemTime, err error) {
	p := "/api/v5/public/estimated-price"
	res, err := c.client.Do(http.MethodGet, p, false)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetLiquidationOrders
// Retrieve information on liquidation orders in the last 7 days.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-liquidation-orders
func (c *Public) GetLiquidationOrders(req requests.GetLiquidationOrders) (response responses.GetLiquidationOrders, err error) {
	p := "/api/v5/public/liquidation-orders"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetMarkPrice
// Retrieve mark price.
//
// We set the mark price based on the SPOT index and at a reasonable basis to prevent individual users from manipulating the market and causing the contract price to fluctuate.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-mark-price
func (c *Public) GetMarkPrice(req requests.GetMarkPrice) (response responses.GetMarkPrice, err error) {
	p := "/api/v5/public/mark-price"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetPositionTiers
// Position information，Maximum leverage depends on your borrowings and margin ratio.
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-position-tiers
func (c *Public) GetPositionTiers(req requests.GetPositionTiers) (response responses.GetPositionTiers, err error) {
	p := "/api/v5/public/position-tiers"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetInterestRateAndLoanQuota
// Get margin interest rate
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-position-tiers
func (c *Public) GetInterestRateAndLoanQuota() (response responses.GetInterestRateAndLoanQuota, err error) {
	p := "/api/v5/public/interest-rate-loan-quota"
	res, err := c.client.Do(http.MethodGet, p, false)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

// GetUnderlying
//
// https://www.okex.com/docs-v5/en/#rest-api-public-data-get-underlying
func (c *Public) GetUnderlying(req requests.GetUnderlying) (response responses.GetUnderlying, err error) {
	p := "/api/v5/public/underlying"
	m := okex.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
