package plex

import (
	"net/http/httptest"
)

var (
	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)
