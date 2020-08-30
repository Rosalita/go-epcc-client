package epcc

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/retry.v1"
)

// Client is the type used to interface with EPCC API.
type Client struct {
	BaseURL       string
	HTTPClient    *http.Client
	RetryStrategy retry.Strategy
	accessToken   string
}

// NewClient creates a new instance of a Client.
// baseURL is optional and can be used to override the default baseURL.
func NewClient(baseURL *string, limitTimeout time.Duration, clientTimeout time.Duration) *Client {

	// If baseURL is not provided, fall back to using the configured default.
	if baseURL == nil {
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

//Authenticate attempts to generate an access token and save it on the client.
func (c *Client) Authenticate() error {

	token, err := Auth(*c)
	if err != nil {
		return err
	}

	c.accessToken = token
	return nil
}

// DoRequest makes a html request to the EPCC API and handles the response.
func (c *Client) DoRequest(method string, path string, payload io.Reader) (body []byte, err error) {

	reqURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	reqURL.Path = path

	req, err := http.NewRequest(method, reqURL.String(), payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	for r := retry.Start(c.RetryStrategy, nil); r.Next(); {
		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 429, 500, 503, 504:
			log.Printf("Response Status %d Retrying request", resp.StatusCode)
			continue

		case 200, 201:
			var buffer bytes.Buffer
			_, err := buffer.ReadFrom(resp.Body)
			if err != nil {
				return nil, err
			}
			return buffer.Bytes(), nil

		case 204:
			return nil, nil

		default:
			err = errors.New("status code not ok")
			return nil, err
		}
	}

	err = errors.New("retry timeout error")
	return nil, err
}
