package stripe

import (
	"net/url"
	"strconv"
)

// see https://stripe.com/docs/api#invoice_object
type Invoice struct {
	// Unique Identifier for this Invoice.
	Id string `json:"id"`

	// Final amount due at this time for this invoice. If the invoice's total is
	// smaller than the minimum charge amount, for example, or if there is an
	// account credit that can be applied to the invoice, the amount_due may be
	// zero.
	AmountDue int64 `json:"amount_due"`

	// Number of automatic payment attempts made for this invoice. Does not
	// include manual attempts to pay the invoice.
	AttemptCount int `json:"attempt_count"`

	// Whether or not an attempt has been made to pay the invoice.
	Attempted bool `json:"attempted"`

	// Whether or not the invoice is still trying to collect payment.
	Closed bool `json:"closed"`

	// Whether or not payment was successfully collected for this invoice.
	Paid bool `json:"paid"`

	// End of the usage period the invoice covers .
	PeriodEnd int64 `json:"period_end"`

	// Start of the usage period the invoice covers
	PeriodStart int64 `json:"period_start"`

	// Total of all subscriptions, invoice items, and prorations on the invoice
	// before any discount is applied
	Subtotal int64 `json:"subtotal"`

	// Total after discounts.
	Total int64 `json:"total"`

	// ID of the latest charge generated for this invoice, if any.
	Charge String `json:"charge"`

	// Customer Identifier linked to this Invoice.
	Customer string `json:"closed"`

	Date int64 `json:"date"`

	// Discount that was applied to this Invoice.
	Discount *Discount `json:"discount"`

	// The individual line items that make up the invoice
	Lines *InvoiceLines `json:"lines"`

	// Starting customer balance before attempting to pay invoice. If the
	// invoice has not been attempted yet, this will be the current customer
	// balance.
	StartingBalance int64 `json:"starting_balance"`

	// Ending customer balance after attempting to pay invoice. If the invoice
	// has not been attempted yet, this will be null.
	EndingBalance Int64 `json:"ending_balance"`

	NextPayment Int64 `json:"next_payment_attempt"`
	Livemode    bool  `json:"livemode"`
}

type InvoiceLines struct {
	InvoiceItems  []*InvoiceItem      `json:"invoiceitems"`
	Prorations    []*InvoiceItem      `json:"prorations"`
	Subscriptions []*SubscriptionItem `json:"subscriptions"`
}

type SubscriptionItem struct {
	Amount int64   `json:"amount"`
	Period *Period `json:"period"`
	Plan   *Plan   `json:"plan"`
}

type Period struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

// InvoiceClient provides an API for Querying Stripe Invoices.
type InvoiceClient struct{}

// Retrieves the invoice with the given ID.
//
// see https://stripe.com/docs/api#retrieve_invoice
func (self *InvoiceClient) Retrieve(id string) (*Invoice, error) {
	invoice := Invoice{}
	path := "/v1/invoices/" + url.QueryEscape(id)
	err := query("GET", path, nil, &invoice)
	return &invoice, err
}

// Retrieves the upcoming invoice for a customer with the given ID.
//
// see https://stripe.com/docs/api#retrieve_customer_invoice
func (self *InvoiceClient) RetrieveCustomer(cid string) (*Invoice, error) {
	invoice := Invoice{}
	values := url.Values{"customer": {cid}}
	err := query("GET", "/v1/invoices/upcoming", values, &invoice)
	return &invoice, err
}

// ListN returns a list of invoices across all customers using the Stripe API's
// default range (count 10, offset 0). The items are returned in sorted order,
// with the most recent items appearing first.
//
// see https://stripe.com/docs/api#list_customer_invoices
func (self *InvoiceClient) List() ([]*Invoice, error) {
	return self.list("", 10, 0)
}

// ListN returns a list of invoices across all customers using the specified
// range. The items are returned in sorted order, with the most recent items
// appearing first.
//
// see https://stripe.com/docs/api#list_customer_invoices
func (self *InvoiceClient) ListN(count int, offset int) ([]*Invoice, error) {
	return self.list("", count, offset)
}

// CustomerList returns a list of invoices for the specified customer id
// using the Stripe API's default range (count 10, offset 0)
//
// see https://stripe.com/docs/api#list_customer_invoices
func (self *InvoiceClient) CustomerList(id string) ([]*Invoice, error) {
	return self.list(id, 10, 0)
}

// CustomerListN returns a list of invoices for the specified customer id,
// range, and count.
//
// see https://stripe.com/docs/api#list_customer_invoices
func (self *InvoiceClient) CustomerListN(id string, count int, offset int) ([]*Invoice, error) {
	return self.list(id, count, offset)
}

func (self *InvoiceClient) list(id string, count int, offset int) ([]*Invoice, error) {
	// define a wrapper function for the Invoice List, so that we can
	// cleanly parse the JSON
	type listInvoicesResp struct{ Data []*Invoice }
	resp := listInvoicesResp{}

	// add the count and offset to the list of url values
	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	// query for customer id, if provided
	if id != "" {
		values.Add("customer", id)
	}

	err := query("GET", "/v1/invoices", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
