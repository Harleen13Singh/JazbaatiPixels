package vendorgateway

import (
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type Client struct {
	client  *fasthttp.Client
	apiKey  string
	model   string
	header  map[string]string
	pool    map[string]string
	timeout time.Duration
}

// Set the response timeout in minutes.
func (c *Client) SetTimeOut(times int) {
	if times != 0 {
		c.timeout = time.Duration(times) * time.Minute
	}
}

func NewClient(c *Config) (*Client, error) {
	client := &Client{
		apiKey: c.ApiKey,
		model:  c.DefaultModel,
		header: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
			"x-api-key":    c.ApiKey,
		},
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader: true,
			Dial:                     fasthttpproxy.FasthttpProxyHTTPDialer(),
		},
	}
	if client.model == "" {
		client.model = "claude-instant-1"
	}
	return client, nil
}

// Send is responsible to make http request with given vendor
func (c *Client) Send(req *VendorRequest) (*VendorResponse, error) {
	if req.Message == nil {
		return nil, ErrEmptyMessage
	}
	if req.SessionId == "" {
		return nil, ErrEmptySessionId
	}
	apiReq, apiRes := c.Acquire(req.Url)
	// release the acquired connection to back to the pool
	defer c.Release(apiReq, apiRes)
	// todo(Harleen Singh): revisit the logic
	return nil, nil
}

// Acquire returns an empty fasthttp instance from request pool.
//
// The returned fasthttp instance may be passed to Release when it is no longer needed.
// This allows Request recycling, reduces GC pressure and usually improves performance.
func (c *Client) Acquire(url string) (*fasthttp.Request, *fasthttp.Response) {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	return req, resp
}

// Release returns req and resp acquired via Acquire to request pool
func (c *Client) Release(req *fasthttp.Request, res *fasthttp.Response) {
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
