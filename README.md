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
* RetryLimitTimout - Requests will be retried for a maximum of the retryLimitTimeout when responses are received with status codes 429 (too many requests), 500 (internal server error), 503 (service unavailable) or 504 (Gateway Timeout)are received. 

# Querying Endpoints

## Currencies 
Make a request to get a single currency.
```go
currency, err := epcc.Currencies.Get(client, "3563bde2-fb72-4721-8584-504058f63780")
```

Make a request to get all currencies.
```go
currencies, err := epcc.Currencies.GetAll(client)
```

Make a request to create a new currency.
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

Make a request to update a currency.
```go
update := epcc.Currency{
	Type:              "currency",
	ExchangeRate:      1.14,
}

result, err := epcc.Currencies.Update(client, "3563bde2-fb72-4721-8584-504058f63780", &update)
```

Make a request to delete a currency.
```go
err := epcc.Currencies.Delete(client, "8240bc0f-6e59-474e-a2fa-6813a0f1b713")
```

## Products
Make a request to get all products.
```go
products, err := epcc.Products.GetAll(client)
```

Make a request to get a single product by ID.
```go
product, err := epcc.Products.Get(client, "78ee7c20-df84-435d-bb1d-531e3537c4dc")
```

Make a request to create a product
```go
newProduct := epcc.Product{
	Type:          "product",
	Name:          "Origami Crane",
	Slug:          "slug-origami-crane",
	SKU:           "sku-origami-crane",
	Description:   "The Origami Crane is considered lucky.",
	ManageStock:   false,
	Status:        "draft",
	CommodityType: "physical",
	Price: []epcc.ProductPrice{
		{
			Amount:      1,
			Currency:    "USD",
			IncludesTax: false,
		},
	},
}

result, err := epcc.Products.Create(client, &newProduct)
```