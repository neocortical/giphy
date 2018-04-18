package giphy

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// DefaultAPIScheme is the default scheme that will be used.
var DefaultAPIScheme = "https"

// DefaultAPIHost is the default host that will be used.
var DefaultAPIHost = "api.giphy.com"

// DefaultAPIVersion is the default API version that will be used.
var DefaultAPIVersion = "v1"

var defaultHTTPClient = &http.Client{
	Timeout: time.Second * 30,
}

// Client defines the Giphy API interface.
type Client interface {
	ClientBuilder
	GIF(id string, options *Options) (GIF, error)
	Random(tags []string, options *Options) (Random, error)
	Search(q string, options *Options) (Search, error)
	Translate(s string, options *Options) (Translate, error)
	Trending(options *Options) (Trending, error)
}

// ClientBuilder defines methods for customizing a Client.
type ClientBuilder interface {
	WithHTTPClient(c *http.Client) Client
	WithBaseURL(scheme, domain, version string) Client
}

// client implements Client.
type client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewClient creates a new API client.
func NewClient(apiKey string) Client {
	return &client{
		httpClient: defaultHTTPClient,
		apiKey:     apiKey,
		baseURL:    fmt.Sprintf("%s://%s/%s", DefaultAPIScheme, DefaultAPIHost, DefaultAPIVersion),
	}
}

// WithHTTPClient allows customization of the net client to use.
func (c *client) WithHTTPClient(newHTTPClient *http.Client) Client {
	c.httpClient = newHTTPClient
	return c
}

// WithBaseURL allows customization of the scheme, domain, and API version
func (c *client) WithBaseURL(scheme, domain, version string) Client {
	c.baseURL = fmt.Sprintf("%s://%s/%s", scheme, domain, version)
	return c
}

// GIF returns a ID response from the Giphy API
func (c *client) GIF(id string, options *Options) (result GIF, err error) {
	var query = buildQuery(c.apiKey, options)

	err = c.get(fmt.Sprintf("/gifs/%s", id), query, &result)
	return
}

// Random returns a random response from the Giphy API
func (c *client) Random(tags []string, options *Options) (result Random, err error) {
	var query = buildQuery(c.apiKey, options)
	if len(tags) > 0 {
		query.Set("tag", strings.Join(tags, ","))
	}

	err = c.get("/gifs/random", query, &result)
	return
}

// Search returns a search response from the Giphy API
func (c *client) Search(q string, options *Options) (result Search, err error) {
	if len(q) == 0 {
		return result, errors.New("no query specified")
	}

	var query = buildQuery(c.apiKey, options)
	query.Set("q", q)

	err = c.get("/gifs/search", query, &result)
	return
}

// Translate returns a translate response from the Giphy API
func (c *client) Translate(s string, options *Options) (result Translate, err error) {
	if len(s) == 0 {
		return result, errors.New("no query specified")
	}

	var query = buildQuery(c.apiKey, options)
	query.Set("s", s)

	err = c.get("/gifs/translate", query, &result)
	return
}

// Trending returns a trending response from the Giphy API
func (c *client) Trending(options *Options) (result Trending, err error) {
	var query = buildQuery(c.apiKey, options)

	err = c.get("/gifs/trending", query, &result)
	return
}
