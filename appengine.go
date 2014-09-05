// +build appengine

package stripe

import (
	"appengine"
	"appengine/urlfetch"
	"net/http"
	"net/url"
	"io"
)

func getHttpClient(r *http.Request) *http.Client {
	c := appengine.NewContext(r)
	return urlfetch.Client(c)
}

func createRequest(method string, endpoint *url.URL, reqBody io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint.String(), reqBody)
	if err == nil {
		req.Header.Set("Stripe-Version", apiVersion)
		req.SetBasicAuth(_key, "")
	}
	return req, err
}