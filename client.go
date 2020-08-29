
package epcc

import (
	"net/http"
	"time"

	"gopkg.in/retry.v1"
)

// Client is the type used to interface with EPCC API.
type Client struct {
	BaseURL       string
	HTTPClient    *http.Client
	RetryStrategy retry.Strategy
}

// NewClient creates a new instance of a Client.
// baseURL is optional and can be used to override default baseURL
func NewClient(baseURL *string, limitTimeout time.Duration, clientTimeout time.Duration) *Client {

	// If baseURL is not provided, fall back to using the configured default
	if baseURL == nil{
		baseURL = &Cfg.BaseURL
	}

	exp := retry.Exponential{
		Initial: 10 * time.Millisecond,
		Factor:  1.5,
		Jitter:  true,
	}
	strategy := retry.LimitTime(limitTimeout, exp)

	return &Client{
		BaseURL: *baseURL,
		HTTPClient: &http.Client{
			Timeout: clientTimeout,
		},
		RetryStrategy: strategy,
	}
}
