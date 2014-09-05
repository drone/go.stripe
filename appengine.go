// +build appengine

package stripe

import (
	"appengine"
	"appengine/urlfetch"
	"net/http"
)

func getHttpClient(r *http.Request) *http.Client {
	c := appengine.NewContext(r)
	return urlfetch.Client(c)
}