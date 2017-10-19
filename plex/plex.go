package plex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/lorac/plex-go/model"
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
	return c
}

func RequestPlexToken(username, password string, httpClient *http.Client) (string, error) {
	var signInResponse model.SignInResponse

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

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}
