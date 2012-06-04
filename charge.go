package stripe

// see https://stripe.com/docs/api#charge_object
type Charge struct {
	Id             string `json:"id"`
	Desc           string `json:"description"`
	Amount         int64  `json:"amount"`   // Amount charged in cents
	Currency       string `json:"currency"` // Three-letter ISO code representing the currency of the charge
	Created        int64  `json:"created"`
	Customer       string `json:"customer"` // ID of the customer this charge is for if one exists
	Invoice        string `json:"invoice"`  // ID of the invoice this charge is for if one exists
	Fee            int64  `json:"fee"`      // Fees (in cents) paid for this charge
	Paid           bool   `json:"paid"`
	Refunded       bool   `json:"refunded"`        // Whether or not the charge has been fully refunded. If the charge is only partially refunded, this attribute will still be false.
	AmountRefunded Int64  `json:"amount_refunded"` // Amount in cents refunded (can be less than the amount attribute on the charge if a partial refund was issued)
	Disputed       bool   `json:"disputed"`        // Whether or not the charge has been disputed by the customer
	Livemode       bool   `json:"livemode"`
}

type CreateChargeReq struct {
	Amount   int64  // A positive integer in cents representing how much to charge the card. The minimum amount is 50 cents.
	Currency string // 3-letter ISO code for currency. Currently, only 'usd' is supported.
	Customer string // (optional) Either customer or card is required, but not both The ID of an existing customer that will be charged in this request.
	Desc     string // An arbitrary string which you can attach to a charge object. It is displayed when in the web interface alongside the charge. It's often a good idea to use an email address as a description for tracking later.
	// TODO need to handle either a hash OR a token ...
	Card *Card // (optional) Either card or customer is required, but not both A card to be charged. The card can either be a token, like the ones returned by Stripe.js, or a hash containing a user's credit card details, with the options described below
}

type ListChargesReq struct {
	Count    int    // (optional, default is 10) A limit on the number of charges to be returned. Count can range between 1 and 100 charges.
	Offset   int    // (optional, default is 0) An offset into your charge array. The API will return the requested number of charges starting at that offset.
	Customer string // Only return charges for the customer specified by this customer ID.
}

type ChargeClient struct{}

// To charge a credit card, you create a new charge object.
//
// Returns a charge object if the charge succeeded. An error will be returned
// if something goes wrong. A common source of error is an invalid or expired
// card, or a valid card with insufficient available balance.
//
// see https://stripe.com/docs/api?lang=ruby#create_charge
func (self *ChargeClient) Create(charge *CreateChargeReq) (*Charge, error) {
	return nil, nil
}

// Retrieves the details of a charge that has previously been created.
//
// see https://stripe.com/docs/api?lang=ruby#retrieve_charge
func (self *ChargeClient) Retrieve(id string) (*Charge, error) {
	return nil, nil
}

// Refunds a charge that has previously been created but not yet refunded,
// for the full amount.
//
// see https://stripe.com/docs/api?lang=ruby#refund_charge
func (self *ChargeClient) Refund(id string) (*Charge, error) {
	return nil, nil
}

// Refunds a charge that has previously been created but not yet refunded,
// for the specified amount.
//
// see https://stripe.com/docs/api?lang=ruby#refund_charge
func (self *ChargeClient) RefundAmount(id string, amt int64) (*Charge, error) {
	return nil, nil
}

// Returns a list of charges you've previously created. The charges are returned
// in sorted order, with the most recent charges appearing first.
//
// see https://stripe.com/docs/api?lang=ruby#list_charges
func (self *ChargeClient) List(opts *ListChargesReq) ([]*Charge, error) {
	return nil, nil
}
