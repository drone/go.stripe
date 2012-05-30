package stripe


// Customer objects 
//
// see https://stripe.com/docs/api#customer_object
type Customer struct {
	Id         string "id"
	Desc       string "description"
	Email      string "email"
	Created    int64  "created"
	Balance    int64  "account_balance" // Current balance, if any, being stored on the customer's account. If negative, the customer has credit to apply to the next invoice. If positive, the customer has an amount owed that will be added to the next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice.
	Delinquent string "delinquent"      // Whether or not the latest charge for the customer's latest invoice has failed
	Livemode   bool   "livemode"
	Card         *Card         "active_card"
	Discount     *Discount     "discount"
	Subscription *Subscription "subscription"
}

// Discount represents the actual application of a coupon to a particular
// customer.
//
// see https://stripe.com/docs/api#discount_object
type Discount struct {
	Id       string  "id"
	Customer string  "customer"
	Start    Int64   "start"    // Date that the coupon was applied
	End      Int64   "end"      // If the coupon has a duration of once or multi-month, the date that this discount will end. If the coupon used has a forever duration, this attribute will be null.
	Coupon   *Coupon "coupon"   // the Coupon applied to create this discount
}

// CustomerClient is a wrapper around the Stripe Customer API, allowing you to
// perform recurring charges and track multiple charges that are associated with
// the same customer. The API allows you to create, delete, and update your
// customers. You can retrieve individual customers as well as a list of all
// your customers.
type CustomerClient struct { }

// Creates a new customer object.
// TODO need some type of CreateCustomerRequest object
// see https://stripe.com/docs/api#create_customer
func (self *CustomerClient) Create() (*Customer, error) {
	return nil, nil
}

// Retrieves the details of an existing customer using the provided customer
// identifier.
//
// see https://stripe.com/docs/api#retrieve_customer
func (self *CustomerClient) Retrieve(id string) (*Customer, error) {
	return nil, nil
}

// see https://stripe.com/docs/api#update_customer
func (self *CustomerClient) Update(c *Customer) (*Customer, error) {
	return nil, nil
}

// Permanently deletes a customer. It cannot be undone.
//
// see https://stripe.com/docs/api#delete_customer
func (self *CustomerClient) Delete(id string) (*Customer, error) {
	return nil, nil
}

// see https://stripe.com/docs/api#list_customers
func (self *CustomerClient) List() ([]*Customer, error) {
	return self.ListN(10, 0)
}

// see https://stripe.com/docs/api#list_customers
func (self *CustomerClient) ListN(count int, offset int) ([]*Customer, error) {
	return nil, nil
}
