// +build !appengine

package stripe

import (
	"net/http"
)

func getHttpClient(r *http.Request) *http.Client {
	client := new(http.Client)
	if client == nil {
		client = &http.Client{}
	}
	return client
}