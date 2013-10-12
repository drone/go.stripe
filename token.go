package stripe

import (
	"net/url"
)

// Token represents a unique identifier for a credit card that can be safely
// stored without having to hold sensitive card information on your own servers.
//
// see https://stripe.com/docs/api#token_object
type Token struct {
	Id       string `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Created  int64  `json:"created"`
	Used     bool   `json:"used"`
	Livemode bool   `json:"livemode"`
	Type     string `json:"type"`
	Card     *Card  `json:"card"`
}

// TokenClient encapsulates operations for creating and querying tokens using
// the Stripe REST API.
type TokenClient struct{}

// TokenParams encapsulates options for creating a new Card Token.
type TokenParams struct {
	//Currency string REMOVED! no longer part of the API
	Card *CardParams
}

// Creates a single use token that wraps the details of a credit card.
// This token can be used in place of a credit card hash with any API method.
// These tokens can only be used once: by creating a new charge object, or
// attaching them to a customer.
//
// see https://stripe.com/docs/api#create_token
func (self *TokenClient) Create(params *TokenParams) (*Token, error) {
	token := Token{}
	values := url.Values{} // REMOVED "currency": {params.Currency}}
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
