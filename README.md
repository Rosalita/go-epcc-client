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

Make a request to create a new currency
```go
newCurrency := epcc.Currency{
	Type: "currency",
	Code: "INR",
	ExchangeRate: 142.15,
	Format: "₹{price}",
	DecimalPoint: ".",
	ThousandSeparator: ",",
	DecimalPlaces: 2,
	Default: false,
	Enabled: true,
}

result, err := epcc.Currencies.Create(client, &newCurrency)
```

Make a request to delete a currency
```go
err := epcc.Currencies.Delete(client, "8240bc0f-6e59-474e-a2fa-6813a0f1b713")
```

# To configure a custom client
```go
clientOptions := epcc.ClientOptions{
	BaseURL: "https://dev.moltin.com/",
	ClientTimeout: 10 * time.Second,
	RetryLimitTimeout: 20 * time.Second,
}

customClient := epcc.NewClient(clientOptions)
customClient.Authenticate()
```

## Client Options
* BaseURL - This is the baseURL that requests will be made to.
* ClientTimeout - This is how long the client will wait for a response before timing out.
* RetryLimitTimout - Requests will be retried for a maximum of the retryLimitTimeout when responses are received with status codes 429 (too many requests), 500, (internal server error) 503 (service unavailable) or 504 (Gateway Timeout)are received. 
