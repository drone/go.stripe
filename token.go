package stripe

import (
	"net/url"
)

// see https://stripe.com/docs/api#token_object
type Token struct {
	Id       int    `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Card     *Card  `json:"card"`
	Created  int64  `json:"created"`
	Used     bool   `json:"used"`
	Livemode bool   `json:"livemode"`
}

type TokenClient struct{}

// TokenParams is a data structure that represents the required input parameters
// for Creating and Credit Card Tokens in the system.
type TokenParams struct {
	Currency string
	Card     *CardParams
}

// Creates a single use token that wraps the details of a credit card.
// This token can be used in place of a credit card hash with any API method.
// These tokens can only be used once: by creating a new charge object, or
// attaching them to a customer.
//
// see https://stripe.com/docs/api#create_token
func (self *TokenClient) Create(params *TokenParams) (*Token, error) {
	token := Token{}
	values := url.Values{ "currency" : {params.Currency}}
	appendCardParamsToValues(params.Card, &values)

	err := query("POST", "/v1/tokens", values, &token)
	return &token, err
}

// Retrieves the card token with the given Id.
//
// see https://stripe.com/docs/api#retrieve_token
func (self *TokenClient) Retrieve(id string) (*Token, error) {
	token := Token{}
	path := "/v1/tokens/" + url.QueryEscape(id)
	err := query("GET", path, nil, &token)
	return &token, err
}
