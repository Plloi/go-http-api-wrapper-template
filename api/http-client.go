package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	// "github.com/dgrijalva/jwt-go" // When JSON Web Tokens are needed.
)

const (
	userAgent = "go-api-client"
)

// Client : The APi Client
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	userAgent  string
}

// APIError : An error from the API with a (hopefuly) useful message.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d - %s", e.StatusCode, e.Message)
}

// NewClient Returns API Client optionally takes an HttpClient as input
func NewClient(httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
		httpClient.Timeout = time.Second * 10
	}

	// TODO: Make Configurable
	baseURL, _ := url.Parse("https://example.com/")
	c := &Client{httpClient: httpClient, baseURL: baseURL, userAgent: userAgent}
	return c
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)
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
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode <= 200 && resp.StatusCode >= 300 {
		return resp, &APIError{resp.StatusCode, "Halp!"}
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
