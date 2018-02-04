package plex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL        = "http://127.0.0.1:32400/"
	plexURL               = "https://plex.tv/"
	XPlexProduct          = "plex-go"
	XPlexVersion          = "0.0.1"
	XPlexClientIdentifier = "plex-go-client-id"
	signInURL             = plexURL + "users/sign_in.json"
)

type Client struct {
	client  *http.Client // HTTP client used to communicate with the API.
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Library *LibraryService
	Status  *StatusService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c
	c.Library = (*LibraryService)(&c.common)
	c.Status = (*StatusService)(&c.common)
	return c
}

func RequestPlexToken(username, password string, httpClient *http.Client) (string, error) {
	var signInResponse SignInResponse

	form := url.Values{}
	form.Set("user[login]", username)
	form.Set("user[password]", password)

	req, err := http.NewRequest("POST", signInURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	json.NewDecoder(resp.Body).Decode(&signInResponse)
	return signInResponse.User.AuthToken, err
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// XML decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if _, err := url.Parse(e.URL); err == nil {
				return nil, e
			}
		}
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}
	return resp, err
}
