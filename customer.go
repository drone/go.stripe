package stripe

import (
	"net/url"
	"strconv"
)

// Customer encapsulates details about a Customer registered in Stripe.
//
// see https://stripe.com/docs/api#customer_object
type Customer struct {
	Id           string        `json:"id"`
	Desc         String        `json:"description,omitempty"`
	Email        String        `json:"email,omitempty"`
	Created      int64         `json:"created"`
	Balance      int64         `json:"account_balance"`
	Delinquent   bool          `json:"delinquent"`
	Cards        CardData      `json:"cards,omitempty"`
	Discount     *Discount     `json:"discount,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	Livemode     bool          `json:"livemode"`
	DefaultCard  String        `json:"default_card"`
}

type CardData struct {
	Object string  `json:"object"`
	Count  int     `json:"count"`
	Url    string  `json:"url"`
	Data   []*Card `json:"data"`
}

// Discount represents the actual application of a coupon to a particular
// customer.
//
// see https://stripe.com/docs/api#discount_object
type Discount struct {
	Id       string  `json:"id"`
	Customer string  `json:"customer"`
	Start    Int64   `json:"start"`
	End      Int64   `json:"end"`
	Coupon   *Coupon `json:"coupon"`
}

// CustomerParams encapsulates options for creating and updating Customers.
type CustomerParams struct {
	// (Optional) The customer's email address.
	Email string

	// (Optional) An arbitrary string which you can attach to a customer object.
	Desc string

	// (Optional) Customer's Active Credit Card
	Card *CardParams

	// (Optional) Customer's Active Credid Card, using a Card Token
	Token string

	// (Optional) If you provide a coupon code, the customer will have a
	// discount applied on all recurring charges.
	Coupon string

	// (Optional) The identifier of the plan to subscribe the customer to. If
	// provided, the returned customer object has a 'subscription' attribute
	// describing the state of the customer's subscription.
	Plan string

	// (Optional) UTC integer timestamp representing the end of the trial period
	// the customer will get before being charged for the first time.
	TrialEnd int64

	// (Optional) An integer amount in cents that is the starting account
	// balance for your customer.
	AccountBalance int64

	// (Optional) A set of key/value pairs that you can attach to a customer
	// object.
	Metadata map[string]string

	// (Optional) The quantity you’d like to apply to the subscription you’re
	// creating.
	Quantity int64
}

// CustomerClient encapsulates operations for creating, updating, deleting and
// querying customers using the Stripe REST API.
type CustomerClient struct{}

// Creates a new Customer.
//
// see https://stripe.com/docs/api#create_customer
func (self *CustomerClient) Create(c *CustomerParams) (*Customer, error) {
	customer := Customer{}
	values := url.Values{}
	appendCustomerParamsToValues(c, &values)

	err := query("POST", "/v1/customers", values, &customer)
	return &customer, err
}

// Retrieves a Customer with the given ID.
//
// see https://stripe.com/docs/api#retrieve_customer
func (self *CustomerClient) Retrieve(id string) (*Customer, error) {
	customer := Customer{}
	path := "/v1/customers/" + url.QueryEscape(id)
	err := query("GET", path, nil, &customer)
	return &customer, err
}

// Updates a Customer with the given ID.
//
// see https://stripe.com/docs/api#update_customer
func (self *CustomerClient) Update(id string, c *CustomerParams) (*Customer, error) {
	customer := Customer{}
	values := url.Values{}
	appendCustomerParamsToValues(c, &values)

	err := query("POST", "/v1/customers/"+url.QueryEscape(id), values, &customer)
	return &customer, err
}

// Deletes a Customer (permanently) with the given ID.
//
// see https://stripe.com/docs/api#delete_customer
func (self *CustomerClient) Delete(id string) (bool, error) {
	resp := DeleteResp{}
	path := "/v1/customers/" + url.QueryEscape(id)
	if err := query("DELETE", path, nil, &resp); err != nil {
		return false, err
	}
	return resp.Deleted, nil
}

// Returns a list of your Customers.
//
// see https://stripe.com/docs/api#list_customers
func (self *CustomerClient) List() ([]*Customer, error) {
	return self.ListN(10, 0)
}

// Returns a list of your Customers at the specified range.
//
// see https://stripe.com/docs/api#list_customers
func (self *CustomerClient) ListN(count int, offset int) ([]*Customer, error) {
	// define a wrapper function for the Customer List, so that we can
	// cleanly parse the JSON
	type listCustomerResp struct{ Data []*Customer }
	resp := listCustomerResp{}

	// add the count and offset to the list of url values
	values := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := query("GET", "/v1/customers", values, &resp)
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
	if c.AccountBalance != 0 {
		values.Add("account_balance", strconv.FormatInt(c.AccountBalance, 10))
	}
	if c.Quantity != 0 {
		values.Add("quantity", strconv.FormatInt(c.Quantity, 10))
	}

	// add metadata, if specified
	for k, v := range c.Metadata {
		values.Add("metadata["+k+"]", v)
	}

	// add optional credit card details, if specified
	if c.Card != nil {
		appendCardParamsToValues(c.Card, values)
	} else if len(c.Token) != 0 {
		values.Add("card", c.Token)
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
