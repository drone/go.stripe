package stripe

import (
	"net/url"
	"strconv"
)

// see https://stripe.com/docs/api#invoiceitem_object
type InvoiceItem struct {
	Id       string `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Customer string `json:"customer"`
	Date     int64  `json:"date"`
	Desc     String `json:"description"`
	Invoice  String `json:"invoice"`
	Livemode bool   `json:"livemode"`
}

// InvoiceItemParams is a data structure that represents the required input
// parameters for Creating and Updating Invoice Items in the system.
type InvoiceItemParams struct {
	// The ID of the customer who will be billed when this invoice item is
	// billed.
	Customer string

	// The integer amount in cents of the charge to be applied to the upcoming
	// invoice. If you want to apply a credit to the customer's account, pass a
	// negative amount.
	Amount int64

	// 3-letter ISO code for currency. Currently, only 'usd' is supported.
	Currency string

	// An arbitrary string which you can attach to the invoice item. The
	// description is displayed in the invoice for easy tracking.
	Desc string

	// The ID of an existing invoice to add this invoice item to. When left
	// blank, the invoice item will be added to the next upcoming scheduled
	// invoice.
	// 
	// You cannot add an invoice item to an invoice that has already been paid
	// or closed.
	Invoice string
}

// InvoiceItemClient provides an API for Creating, Updating, Removing and
// Querying Stripe Invoice Items.
type InvoiceItemClient struct{}

// Create adds an arbitrary charge or credit to the customer's upcoming invoice.
//
// see https://stripe.com/docs/api#invoiceitem_object
func (self *InvoiceItemClient) Create(params *InvoiceItemParams) (*InvoiceItem, error) {
	item := InvoiceItem{}
	values := url.Values{
		"amount":   {strconv.FormatInt(params.Amount, 10)},
		"currency": {params.Currency},
		"customer": {params.Customer},
	}

	// add optional parameters
	if len(params.Desc) != 0 {
		values.Add("description", params.Desc)
	}
	if len(params.Invoice) != 0 {
		values.Add("invoice", params.Invoice)
	}

	err := query("POST", "/v1/invoiceitems", values, &item)
	return &item, err
}

// Retrieves the invoice item with the given ID.
//
// see https://stripe.com/docs/api#retrieve_invoiceitem
func (self *InvoiceItemClient) Retrieve(id string) (*InvoiceItem, error) {
	item := InvoiceItem{}
	path := "/v1/invoiceitems/" + url.QueryEscape(id)
	err := query("GET", path, nil, &item)
	return &item, err
}

// Update changes the amount or description of an invoice item on an upcoming
// invoice, using the specified invoice item id.
//
// see https://stripe.com/docs/api#update_invoiceitem
func (self *InvoiceItemClient) Update(id string, params *InvoiceItemParams) (*InvoiceItem, error) {
	item := InvoiceItem{}
	values := url.Values{}

	if len(params.Desc) != 0 {
		values.Add("description", params.Desc)
	}
	if params.Amount != 0 {
		values.Add("invoice", strconv.FormatInt(params.Amount, 10))
	}

	err := query("POST", "/v1/invoiceitems/"+url.QueryEscape(id), values, &item)
	return &item, err
}

// Removes an invoice item from the upcoming invoice. Removing an invoice item
// is only possible before the invoice it's attached to is closed.
// 
// see https://stripe.com/docs/api#delete_invoiceitem
func (self *InvoiceItemClient) Delete(id string) (*InvoiceItem, error) {
	item := InvoiceItem{}
	path := "/v1/invoiceitems/" + url.QueryEscape(id)
	err := query("DELETE", path, nil, &item)
	return &item, err
}

// ListN returns a list of invoice items across all customers using the Stripe API's
// default range (count 10, offset 0). The items are returned in sorted order,
// with the most recent items appearing first.
//
// see https://stripe.com/docs/api#list_invoiceitems
func (self *InvoiceItemClient) List() ([]*InvoiceItem, error) {
	return self.list("", 10, 0)
}

// ListN returns a list of invoice items across all customers using the specified
// range. The items are returned in sorted order, with the most recent items
// appearing first.
//
// see https://stripe.com/docs/api#list_invoiceitems
func (self *InvoiceItemClient) ListN(count int, offset int) ([]*InvoiceItem, error) {
	return self.list("", count, offset)
}

// CustomerList returns a list of invoice items for the specified customer id
// using the Stripe API's default range (count 10, offset 0)
//
// see https://stripe.com/docs/api#list_invoiceitems
func (self *InvoiceItemClient) CustomerList(id string) ([]*InvoiceItem, error) {
	return self.list(id, 10, 0)
}

// CustomerListN returns a list of invoice items for the specified customer id,
// range, and count.
//
// see https://stripe.com/docs/api#list_invoiceitems
func (self *InvoiceItemClient) CustomerListN(id string, count int, offset int) ([]*InvoiceItem, error) {
	return self.list(id, count, offset)
}

func (self *InvoiceItemClient) list(id string, count int, offset int) ([]*InvoiceItem, error) {
	// define a wrapper function for the Invoice Items List, so that we can
	// cleanly parse the JSON
	type listInvoiceItemsResp struct{ Data []*InvoiceItem }
	resp := listInvoiceItemsResp{}

	// add the count and offset to the list of url values
	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	// query for customer id, if provided
	if id != "" {
		values.Add("customer", id)
	}

	err := query("GET", "/v1/invoiceitems", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
