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


# Usage
Create a new API client with default options and authenticate
```go
	client := epcc.NewClient()
	client.Authenticate()
```

Make a request to get all currencies
```go
currencies, err := epcc.Currencies.GetAll(client)
```

To configure a custom client
```go
clientOptions := epcc.ClientOptions{
	BaseURL: "https://dev.moltin.com/",
	ClientTimeout: 10 * time.Second,
	RetryLimitTimeout: 20 * time.Second,
}

customClient := epcc.NewClient(clientOptions)
customClient.Authenticate()
```

# Client Options
* BaseURL - This is the baseURL that requests will be made to.
* ClientTimeout - This is how long the client will wait for a response before timing out.
* RetryLimitTimout - Requests will be retried for a maximum of the retryLimitTimeout when responses are received with status codes 429 (too many requests), 500, (internal server error) 503 (service unavailable) or 504 (Gateway Timeout)are received. 
