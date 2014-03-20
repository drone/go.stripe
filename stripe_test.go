package stripe

import (
	"net/url"
	"testing"
)

func init() {
	// In order to execute Unit Test, you must set your Stripe API Key as
	// environment variable, STRIPE_API_KEY=xxxx
	if err := SetKeyEnv(); err != nil {
		panic(err)
	}
}

func TestClientKey(t *testing.T) {
	// Provide an invalid default key and defer reset to normal
	key := _key
	_key = ""
	defer func() { _key = key }()

	client := &Client{Key: key}

	values := url.Values{}
	var i interface{}
	err := client.query("GET", "/v1/account", values, &i)

	if err != nil {
		t.Errorf("Expected success, got Error %s", err.Error())
	}
}

func TestClientKeyError(t *testing.T) {
	// Provide an invalid client key and a valid default key
	client := &Client{Key: "123"}

	values := url.Values{}
	var i interface{}
	err := client.query("GET", "/v1/account", values, &i)

	if err == nil {
		t.Errorf("Expected error, got success!")
	}
}
