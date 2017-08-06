package meetup

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/net/context/ctxhttp"
)

var (
	apiEndpointBase                 = "https://api.meetup.com"
	apiEndpointGETGroupEvents       = "/%s/events"
	apiEndpointGETGroupEvent        = "/%s/events/%s"
	apiEndpointGETEvents            = "/find/events"
	apiEndpointGETRecommendedEvents = "/recommended/events"
	apiEndpointGETMyCalendar        = "/self/calendar"
	apiEndpointGETMyEvents          = "/self/events"
)

type Client struct {
	apiKey     string
	URLBase    *url.URL
	HttpClient *http.Client
}

func NewClient(key string) (*Client, error) {
	meetupURL, err := url.ParseRequestURI(apiEndpointBase)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiKey:     key,
		URLBase:    meetupURL,
		HttpClient: http.DefaultClient,
	}, nil
}

func (c *Client) url(endpoint string) string {
	u := *c.URLBase
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

func (c *Client) get(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.url(endpoint), nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req)
}

func (c *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.url(endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return c.do(ctx, req)
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		return ctxhttp.Do(ctx, c.HttpClient, req)
	}
	return c.HttpClient.Do(req)
}
