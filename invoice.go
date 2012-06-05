package stripe

import (
	"net/url"
	"strconv"
)

// Invoice represents statements of what a customer owes for a particular
// billing period, including subscriptions, invoice items, and any automatic
// proration adjustments if necessary.
//
// see https://stripe.com/docs/api#invoice_object
type Invoice struct {
	Id              string        `json:"id"`
	AmountDue       int64         `json:"amount_due"`
	AttemptCount    int           `json:"attempt_count"`
	Attempted       bool          `json:"attempted"`
	Closed          bool          `json:"closed"`
	Paid            bool          `json:"paid"`
	PeriodEnd       int64         `json:"period_end"`
	PeriodStart     int64         `json:"period_start"`
	Subtotal        int64         `json:"subtotal"`
	Total           int64         `json:"total"`
	Charge          String        `json:"charge"`
	Customer        string        `json:"closed"`
	Date            int64         `json:"date"`
	Discount        *Discount     `json:"discount"`
	Lines           *InvoiceLines `json:"lines"`
	StartingBalance int64         `json:"starting_balance"`
	EndingBalance   Int64         `json:"ending_balance"`
	NextPayment     Int64         `json:"next_payment_attempt"`
	Livemode        bool          `json:"livemode"`
}

// InvoiceLines represents an individual line items that is part of an invoice.
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

// InvoiceClient encapsulates operations for querying invoices using the Stripe
// REST API.
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

// Retrieves the upcoming invoice the given customer ID.
//
// see https://stripe.com/docs/api#retrieve_customer_invoice
func (self *InvoiceClient) RetrieveCustomer(cid string) (*Invoice, error) {
	invoice := Invoice{}
	values := url.Values{"customer": {cid}}
	err := query("GET", "/v1/invoices/upcoming", values, &invoice)
	return &invoice, err
}

// Returns a list of Invoices.
//
// see https://stripe.com/docs/api#list_customer_invoices
func (self *InvoiceClient) List() ([]*Invoice, error) {
	return self.list("", 10, 0)
}

// Returns a list of Invoices at the specified range.
//
// see https://stripe.com/docs/api#list_customer_invoices
func (self *InvoiceClient) ListN(count int, offset int) ([]*Invoice, error) {
	return self.list("", count, offset)
}

// Returns a list of Invoices with the given Customer ID.
//
// see https://stripe.com/docs/api#list_customer_invoices
func (self *InvoiceClient) CustomerList(id string) ([]*Invoice, error) {
	return self.list(id, 10, 0)
}

// Returns a list of Invoices with the given Customer ID, at the specified range.
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
