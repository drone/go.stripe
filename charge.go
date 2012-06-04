package stripe

import (
	"net/url"
	"strconv"
)

// ISO 3-digit Currency Codes for major currencies (not the full list).
const (
	USD = "usd" // US Dollar ($)
	EUR = "eur" // Euro (€)
	GBP = "gbp" // British Pound Sterling (UK£)
	JPY = "jpy" // Japanese Yen (¥)
	CAD = "cad" // Canadian Dollar (CA$)
	HKD = "hkd" // Hong Kong Dollar (HK$)
	CNY = "cny" // Chinese Yuan (CN¥)
	AUD = "aud" // Australian Dollar (A$)
)

// see https://stripe.com/docs/api#charge_object
type Charge struct {
	Id   string `json:"id"`
	Desc String `json:"description"`

	// Amount charged in cents
	Amount int64 `json:"amount"`

	// Credit Card used to make the charge
	Card *Card `json:"card"`

	// Three-letter ISO code representing the currency of the charge
	Currency string `json:"currency"`
	Created  int64  `json:"created"`

	// ID of the customer this charge is for if one exists
	Customer String `json:"customer"`

	// ID of the invoice this charge is for if one exists
	Invoice String `json:"invoice"`

	// Fees (in cents) paid for this charge
	Fee  int64 `json:"fee"`
	Paid bool  `json:"paid"`

	// Detailed breakdown of fees (in cents) paid for this charge
	Details []*FeeDetails `json:"fee_details"`

	// Whether or not the charge has been fully refunded. If the charge is only
	// partially refunded, this attribute will still be false.
	Refunded bool `json:"refunded"`

	// Amount in cents refunded (can be less than the amount attribute on the
	// charge if a partial refund was issued)
	AmountRefunded Int64 `json:"amount_refunded"`

	// Message to user further explaining reason for charge failure if available
	FailureMessage String `json:"failure_message"`

	// Whether or not the charge has been disputed by the customer
	Disputed bool `json:"disputed"`
	Livemode bool `json:"livemode"`
}

type FeeDetails struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Type        string `json:"type"`
	Application String `json:"application"`
}

type ChargeParams struct {
	// A positive integer in cents representing how much to charge the card.
	// The minimum amount is 50 cents.
	Amount int64

	// 3-letter ISO code for currency. Currently, only 'usd' is supported.
	Currency string

	// (optional) Either customer or card is required, but not both The ID of an
	// existing customer that will be charged in this request.
	Customer string

	// (optional) Either card or customer is required, but not both A card to be
	// charged. The card can either be a token, like the ones returned by
	// Stripe.js, or a hash containing a user's credit card details, with the
	// options described below
	Card *CardParams

	// An arbitrary string which you can attach to a charge object. It is
	// displayed when in the web interface alongside the charge. It's often a
	// good idea to use an email address as a description for tracking later.
	Desc string
}

type ChargeClient struct{}

// To charge a credit card, you create a new charge object.
//
// Returns a charge object if the charge succeeded. An error will be returned
// if something goes wrong. A common source of error is an invalid or expired
// card, or a valid card with insufficient available balance.
//
// see https://stripe.com/docs/api#create_charge
func (self *ChargeClient) Create(params *ChargeParams) (*Charge, error) {
	charge := Charge{}
	values := url.Values{
		"amount":      {strconv.FormatInt(params.Amount, 10)},
		"currency":    {params.Currency},
		"description": {params.Desc},
	}

	// add optional credit card details, if specified
	if params.Card != nil {
		appendCardParamsToValues(params.Card, &values)
	} else {
		values.Add("customer", params.Customer)
	}

	err := query("POST", "/v1/charges", values, &charge)
	return &charge, err
}

// Retrieves the details of a charge that has previously been created.
//
// see https://stripe.com/docs/api#retrieve_charge
func (self *ChargeClient) Retrieve(id string) (*Charge, error) {
	charge := Charge{}
	path := "/v1/charges/" + url.QueryEscape(id)
	err := query("GET", path, nil, &charge)
	return &charge, err
}

// Refunds a charge that has previously been created but not yet refunded,
// for the full amount.
//
// see https://stripe.com/docs/api#refund_charge
func (self *ChargeClient) Refund(id string) (*Charge, error) {
	values := url.Values{}
	charge := Charge{}
	path := "/v1/charges/" + url.QueryEscape(id) + "/refund"
	err := query("POST", path, values, &charge)
	return &charge, err
}

// Refunds a charge that has previously been created but not yet refunded,
// for the specified amount.
//
// see https://stripe.com/docs/api#refund_charge
func (self *ChargeClient) RefundAmount(id string, amt int64) (*Charge, error) {
	values := url.Values{
		"amount": {strconv.FormatInt(amt, 10)},
	}
	charge := Charge{}
	path := "/v1/charges/" + url.QueryEscape(id) + "/refund"
	err := query("POST", path, values, &charge)
	return &charge, err
}

// ListN returns a list of charges across all customers using the Stripe API's
// default range (count 10, offset 0). The charges are returned in sorted order,
// with the most recent charges appearing first.
//
// see https://stripe.com/docs/api#list_charges
func (self *ChargeClient) List() ([]*Charge, error) {
	return self.list("", 10, 0)
}

// ListN returns a list of charges across all customers using the specified
// range. The charges are returned in sorted order, with the most recent charges
// appearing first.
//
// see https://stripe.com/docs/api#list_charges
func (self *ChargeClient) ListN(count int, offset int) ([]*Charge, error) {
	return self.list("", count, offset)
}

// CustomerList returns a list of charges for the specified customer id using
// the Stripe API's default range (count 10, offset 0)
//
// see https://stripe.com/docs/api#list_charges
func (self *ChargeClient) CustomerList(id string) ([]*Charge, error) {
	return self.list(id, 10, 0)
}

// CustomerListN returns a list of charges for the specified customer id and
// range, and count.
//
// see https://stripe.com/docs/api#list_charges
func (self *ChargeClient) CustomerListN(id string, count int, offset int) ([]*Charge, error) {
	return self.list(id, count, offset)
}

func (self *ChargeClient) list(id string, count int, offset int) ([]*Charge, error) {
	// define a wrapper function for the Charge List, so that we can
	// cleanly parse the JSON
	type listChargesResp struct{ Data []*Charge }
	resp := listChargesResp{}

	// add the count and offset to the list of url values
	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	// query for customer id, if provided
	if id != "" {
		values.Add("customer", id)
	}

	err := query("GET", "/v1/charges", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
