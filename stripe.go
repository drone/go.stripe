package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// enable logging to print the request and reponses to stdout
var _log bool

// the API Key used to authenticate all Stripe API requests
var _key string

// the default URL for all Stripe API requests
var _url string = "https://api.stripe.com"

const apiVersion = "2013-08-13"

// SetUrl will override the default Stripe API URL. This is primarily used
// for unit testing.
func SetUrl(url string) {
	_url = url
}

// SetKey will set the default Stripe API key used to authenticate all Stripe
// API requests.
func SetKey(key string) {
	_key = key
}

// Available APIs
var (
	Charges       = new(ChargeClient)
	Coupons       = new(CouponClient)
	Customers     = new(CustomerClient)
	Invoices      = new(InvoiceClient)
	InvoiceItems  = new(InvoiceItemClient)
	Plans         = new(PlanClient)
	Subscriptions = new(SubscriptionClient)
	Tokens        = new(TokenClient)
)

// SetKeyEnv retrieves the Stripe API key using the STRIPE_API_KEY environment
// variable.
func SetKeyEnv() (err error) {
	_key = os.Getenv("STRIPE_API_KEY")
	if _key == "" {
		err = errors.New("STRIPE_API_KEY not found in environment")
	}
	return
}

// query submits an http.Request and parses the JSON-encoded http.Response,
// storing the result in the value pointed to by v.
func query(method, path string, values url.Values, v interface{}) error {
	// parse the stripe URL
	endpoint, err := url.Parse(_url)
	if err != nil {
		return err
	}

	// set the endpoint for the specific API
	endpoint.Path = path
	endpoint.User = url.User(_key)

	// if this is an http GET, add the url.Values to the endpoint
	if method == "GET" {
		endpoint.RawQuery = values.Encode()
	}

	// else if this is not a GET, encode the url.Values in the body.
	var reqBody io.Reader
	if method != "GET" && values != nil {
		reqBody = strings.NewReader(values.Encode())
	}

	// Log request if logging enabled
	if _log {
		fmt.Println("REQUEST: ", method, endpoint.String())
		fmt.Println(values.Encode())
	}

	// create the request
	req, err := http.NewRequest(method, endpoint.String(), reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Stripe-Version", apiVersion)

	// submit the http request
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// read the body of the http message into a byte array
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	// Log response if logging enabled
	if _log {
		fmt.Println("RESPONSE: ", r.StatusCode)
		fmt.Println(string(body))
	}

	// is this an error?
	if r.StatusCode != 200 {
		error := Error{}
		json.Unmarshal(body, &error)
		return &error
	}

	//parse the JSON response into the response object
	return json.Unmarshal(body, v)
}

// Response to a Deletion request.
type DeleteResp struct {
	// ID of the Object that was deleted
	Id string `json:"id"`
	// Boolean value indicating object was successfully deleted.
	Deleted bool `json:"deleted"`
}
