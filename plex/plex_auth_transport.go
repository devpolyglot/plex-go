package plex

import (
	"net/http"
)

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Basic Authentication with the provided username and password.
type PlexAuthTransport struct {
	XPlexToken string // X-Plex-Token

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *PlexAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req) // per RoundTrip contract
	req.Header.Set("X-Plex-Product", XPlexProduct)
	req.Header.Set("X-Plex-Version", XPlexVersion)
	req.Header.Set("X-Plex-Client-Identifier", XPlexClientIdentifier)

	if len(t.XPlexToken) > 0 {
		req.Header.Set("X-Plex-Token", t.XPlexToken)
	}

	return t.transport().RoundTrip(req)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *PlexAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *PlexAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request. The clone is a
// shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
