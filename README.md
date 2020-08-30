# go-epcc-client
A simple API client built in Go to interact with EPCC endpoints.

This project was created as a learning activity.

# Setup
## Create the following environment variables
* GO_EPCC_CLIENT_ID 
* GO_EPCC_CLIENT_SECRET

Both these values can be found on the home page of your store dashboard.
Add them to .zshrc or .bashrc
`export GO_EPCC_CLIENT_ID=XXX`
`export GO_EPCC_CLIENT_SECRET=YYY`

# Timeouts
The core Go team did not set any timeouts on the standard net/http client. When creating a new client, a clientTimeout must be set. This is how long the client will wait for a response before timing out.

An exponential back-off for retrying has also been implemented using gopkg.in/retry.v1. This will retry for a maximum of the limitTimeout when responses are received with status codes 429 (too many requests), 500, (internal server error) 503 (service unavailable) or 504 (Gateway Timeout)are received. Note: limitTimeout is deliberately set to 10ms in unit tests to reduce the number of retries.


# Usage
Create a new API client and authenticate
```go
limitTimeout := 10 * time.Millisecond
clientTimeout := 10 * time.Second
client := epcc.NewClient(nil, limitTimeout, clientTimeout)
client.Authenticate()
```

Make a request to get all currencies
```go
currencies, err := epcc.Currencies.GetAll(client)
```