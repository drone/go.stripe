// +build !appengine

package stripe

import (
	"net/http"
	"net/url"
	"io"
)

func getHttpClient(r *http.Request) *http.Client {
	client := new(http.Client)
	if client == nil {
		client = &http.Client{}
	}
	return client
}

func createRequest(method string, endpoint *url.URL, reqBody io.Reader) (*http.Request, error) {
	endpoint.User = url.User(_key)
	req, err := http.NewRequest(method, endpoint.String(), reqBody)
	if err == nil {
		req.Header.Set("Stripe-Version", apiVersion)
	}
	return req, err
}