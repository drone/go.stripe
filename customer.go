package stripe

import (
	"net/url"
	"strconv"
)

// Customer objects 
//
// see https://stripe.com/docs/api#customer_object
type Customer struct {
	// Customer's Unique Identifier within the Stripe database.
	Id   string `json:"id"`
	Desc string `json:"description"`

	// Customer's Email address
	Email   string `json:"email"`
	Created int64  `json:"created"`

	// Current balance, if any, being stored on the customer's account. If
	// negative, the customer has credit to apply to the next invoice. If
	// positive, the customer has an amount owed that will be added to the next
	// invoice. The balance does not refer to any unpaid invoices; it solely
	// takes into account amounts that have yet to be successfully applied to
	// any invoice.
	Balance int64 `json:"account_balance"`

	// Whether or not the latest charge for the customer's latest invoice has
	// failed.
	Delinquent string `json:"delinquent"`

	// Describes the active credit card for the customer, if there is one.
	Card *Card `json:"active_card"`

	// Describes the current discount active for the customer, if there is one.
	Discount *Discount `json:"discount"`

	// Describes the current subscription on the customer, if there is one. If
	// the customer has no current subscription, this will be null.
	Subscription *Subscription `json:"subscription"`

	Livemode bool `json:"livemode"`
}

// Discount represents the actual application of a coupon to a particular
// customer.
//
// see https://stripe.com/docs/api#discount_object
type Discount struct {
	// Discount's Unique Identifier within the Stripe database.
	Id string `json:"id"`

	// Customer's Unique Identifier that has received this discount.
	Customer string  `json:"customer"`

	// Date that the coupon was applied
	Start Int64 `json:"start"`

	// If the coupon has a duration of once or multi-month, the date that this
	// discount will end. If the coupon used has a forever duration, this
	// attribute will be null.
	End Int64 `json:"end"`

	// the Coupon applied to create this discount
	Coupon *Coupon `json:"coupon"`
}

// CustomerParams is a data structure that represents the required input
// parameters for Creating and Updating Customer data in the system.
type CustomerParams struct {
	Email, Desc string

	// Credit Card to attach to the customer.
	Card *CardParams

	// If you provide a coupon code, the customer will have a discount applied
	// on all recurring charges.
	Coupon string

	// The identifier of the plan to subscribe the customer to. If provided,
	// the returned customer object has a 'subscription' attribute describing
	// the state of the customer's subscription.
	Plan string

	// UTC integer timestamp representing the end of the trial period the
	// customer will get before being charged for the first time.
	TrialEnd int64
}

// CustomerClient is a wrapper around the Stripe Customer API, allowing you to
// perform recurring charges and track multiple charges that are associated with
// the same customer. The API allows you to create, delete, and update your
// customers. You can retrieve individual customers as well as a list of all
// your customers.
type CustomerClient struct{}

// Creates a new customer object.
//
// see https://stripe.com/docs/api#create_customer
func (self *CustomerClient) Create(c *CustomerParams) (*Customer, error) {
	customer := Customer{}
	values := url.Values{}
	appendCustomerParamsToValues(c, &values)

	err := query("POST", "/v1/customers", values, &customer)
	return &customer, err
}

// Retrieves the details of an existing customer using the provided customer
// identifier.
//
// see https://stripe.com/docs/api#retrieve_customer
func (self *CustomerClient) Retrieve(id string) (*Customer, error) {
	customer := Customer{}
	path := "/v1/customers/" + url.QueryEscape(id)
	err := query("GET", path, nil, &customer)
	return &customer, err
}

// Updates the specified customer by setting the values of the parameters
// passed.
//
// see https://stripe.com/docs/api#update_customer
func (self *CustomerClient) Update(id string, c *CustomerParams) (*Customer, error) {
	customer := Customer{}
	values := url.Values{}
	appendCustomerParamsToValues(c, &values)

	err := query("POST", "/v1/customers/"+ url.QueryEscape(id), values, &customer)
	return &customer, err
}

// Permanently deletes a customer. It cannot be undone.
//
// see https://stripe.com/docs/api#delete_customer
func (self *CustomerClient) Delete(id string) (*Customer, error) {
	customer := Customer{}
	path := "/v1/customers/" + url.QueryEscape(id)
	err := query("DELETE", path, nil, &customer)
	return &customer, err
}

// see https://stripe.com/docs/api#list_customers
func (self *CustomerClient) List() ([]*Customer, error) {
	return self.ListN(10, 0)
}

// see https://stripe.com/docs/api#list_customers
func (self *CustomerClient) ListN(count int, offset int) ([]*Customer, error) {
	// define a wrapper function for the Customer List, so that we can
	// cleanly parse the JSON
	type listCustomerResp struct{ Data []*Customer }

	resp := listCustomerResp{}
	err := query("GET", "/v1/customers", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}


////////////////////////////////////////////////////////////////////////////////
// Helper Function(s)

func appendCustomerParamsToValues(c *CustomerParams, values *url.Values) {
	// add optional parameters, if specified
	if c.Email != "" {
		values.Add("email", c.Email)
	}
	if c.Desc != "" {
		values.Add("description", c.Desc)
	}
	if c.Coupon != "" {
		values.Add("coupon", c.Coupon)
	}
	if c.Plan != "" {
		values.Add("plan", c.Plan)
	}
	if c.TrialEnd != 0 {
		values.Add("trial_end", strconv.FormatInt(c.TrialEnd, 10))
	}

	// add optional credit card details, if specified
	if c.Card != nil {
		appendCardParamsToValues(c.Card, values)
	}
}

func appendCardParamsToValues(c *CardParams, values *url.Values) {
	values.Add("card[number]", c.Number)
	values.Add("card[exp_month]", strconv.Itoa(c.ExpMonth))
	values.Add("card[exp_year]", strconv.Itoa(c.ExpYear))
	if c.Name != "" {
		values.Add("card[name]", c.Name)
	}
	if c.CVC != "" {
		values.Add("card[cvc]", c.CVC)
	}
	if c.Address1 != "" {
		values.Add("card[address_line1]", c.Address1)
	}
	if c.Address2 != "" {
		values.Add("card[address_line2]", c.Address2)
	}
	if c.AddressZip != "" {
		values.Add("card[address_zip]", c.AddressZip)
	}
	if c.AddressState != "" {
		values.Add("card[address_state]", c.AddressState)
	}
	if c.AddressCountry != "" {
		values.Add("card[address_country]", c.AddressCountry)
	}
}
