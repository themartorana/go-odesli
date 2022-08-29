package odesli

import (
	"net/http"
)

type RoundTripperFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TransportWithAPIToken(t *http.Transport, apiToken string) http.RoundTripper {
	return RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		q := req.URL.Query()
		q.Set("key", apiToken)
		req.URL.RawQuery = q.Encode()
		return t.RoundTrip(req)
	})
}
