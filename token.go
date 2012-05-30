package stripe

// see https://stripe.com/docs/api#token_object
type Token struct {
	Id       int    "id"
	Amount   int64  "amount"
	Currency string "currency"
	Card     *Card  "card"
	Created  int64  "created"
	Used     bool   "used"
	Livemode bool   "livemode"
}


type TokenClient struct { }

// Creates a single use token that wraps the details of a credit card.
// This token can be used in place of a credit card hash with any API method.
// These tokens can only be used once: by creating a new charge object, or
// attaching them to a customer.
//
// see https://stripe.com/docs/api?lang=ruby#create_token
func (self *TokenClient) Create(card *Card, currency string) (*Token, error) {
	return nil, nil
}

// Retrieves the card token with the given Id.
//
// see https://stripe.com/docs/api?lang=java#retrieve_token
func (self *TokenClient) Retrieve(id string) (*Token, error) {
	return nil, nil
}
