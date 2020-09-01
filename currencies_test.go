package epcc_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Rosalita/go-epcc-client"
	"github.com/stretchr/testify/assert"
)

func fakeHandleCurrenciesGetAll(rw http.ResponseWriter, req *http.Request) {

	switch {
	case req.URL.String() == "/v2/currencies" && req.Method == "GET":

		responseJSON := `{` +
			`"data":[` +
			`{` +
			`"id":"9ea32a77-982b-41fd-873c-62570bbfc1e0",` +
			`"type":"currency",` +
			`"code":"GBP",` +
			`"exchange_rate":1.33,` +
			`"format":"£{price}",` +
			`"decimal_point":".",` +
			`"thousand_separator":",",` +
			`"decimal_places":2,` +
			`"default":true,` +
			`"enabled":true,` +
			`"links":{` +
			`"self":"https://api.moltin.com/currencies/9ea32a77-982b-41fd-873c-62570bbfc1e0"` +
			`},` +
			`"meta":{` +
			`"timestamps":{` +
			`"created_at":"2020-08-28T09:56:36.852Z",` +
			`"updated_at":"2020-08-28T09:56:36.852Z"` +
			`}}},{` +
			`"id":"9cc80fcd-115c-47ba-9ff0-07072b378ded",` +
			`"type":"currency",` +
			`"code":"USD",` +
			`"exchange_rate":1,` +
			`"format":"${price}",` +
			`"decimal_point":".",` +
			`"thousand_separator":",",` +
			`"decimal_places":2,` +
			`"default":false,` +
			`"enabled":true,` +
			`"links":{` +
			`"self":"https://api.moltin.com/currencies/9cc80fcd-115c-47ba-9ff0-07072b378ded"` +
			`},` +
			`"meta":{` +
			`"timestamps":{` +
			`"created_at":"2020-08-28T09:45:11.822Z",` +
			`"updated_at":"2020-08-28T09:45:11.822Z"` +
			`}}}]}`

		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestCurrenciesGetAll(t *testing.T) {
	expectedCurrencies := epcc.CurrenciesData{
		Data: []epcc.Currency{
			epcc.Currency{
				ID:                "9ea32a77-982b-41fd-873c-62570bbfc1e0",
				Type:              "currency",
				Code:              "GBP",
				ExchangeRate:      1.33,
				Format:            "£{price}",
				DecimalPoint:      ".",
				ThousandSeparator: ",",
				DecimalPlaces:     2,
				Default:           true,
				Enabled:           true,
				Links: epcc.Links{
					Self: "https://api.moltin.com/currencies/9ea32a77-982b-41fd-873c-62570bbfc1e0",
				},
				Meta: epcc.CurrencyMeta{
					Timestamps: epcc.Timestamps{
						CreatedAt: "2020-08-28T09:56:36.852Z",
						UpdatedAt: "2020-08-28T09:56:36.852Z",
					},
				},
			},
			epcc.Currency{
				ID:                "9cc80fcd-115c-47ba-9ff0-07072b378ded",
				Type:              "currency",
				Code:              "USD",
				ExchangeRate:      1,
				Format:            "${price}",
				DecimalPoint:      ".",
				ThousandSeparator: ",",
				DecimalPlaces:     2,
				Default:           false,
				Enabled:           true,
				Links: epcc.Links{
					Self: "https://api.moltin.com/currencies/9cc80fcd-115c-47ba-9ff0-07072b378ded",
				},
				Meta: epcc.CurrencyMeta{
					Timestamps: epcc.Timestamps{
						CreatedAt: "2020-08-28T09:45:11.822Z",
						UpdatedAt: "2020-08-28T09:45:11.822Z",
					},
				},
			},
		},
	}

	tests := []struct {
		currenciesData epcc.CurrenciesData
		err            error
	}{
		{expectedCurrencies, nil},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleCurrenciesGetAll))
	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		currenciesData, err := epcc.Currencies.GetAll(client)
		assert.Equal(t, test.currenciesData, *currenciesData)
		assert.Equal(t, test.err, err)
	}
}

func fakeHandleCurrenciesCreate(rw http.ResponseWriter, req *http.Request) {

	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(req.Body)
	if err != nil {
		rw.WriteHeader(500)
		return
	}

	switch {
	case req.URL.String() == "/v2/currencies" && req.Method == "POST" && strings.Contains(buffer.String(), `"code":"INR"`):

		responseJSON := `{` +
			`"data":` +
			`{` +
			`"id":"f8f0689e-4767-4924-b112-be89f490e1f5",` +
			`"type":"currency",` +
			`"code":"INR",` +
			`"exchange_rate":142.15,` +
			`"format":"₹{price}",` +
			`"decimal_point":".",` +
			`"thousand_separator":",` +
			`",` +
			`"decimal_places":2,` +
			`"default":false,` +
			`"enabled":true,` +
			`"links":{` +
			`"self":"https://api.moltin.com/currencies/f8f0689e-4767-4924-b112-be89f490e1f5"` +
			`},` +
			`"meta":{` +
			`"timestamps":{` +
			`"created_at":"2020-09-01T15:48:10.050234331Z",` +
			`"updated_at":"2020-09-01T15:48:10.050234395Z"` +
			`}}}}`

		rw.WriteHeader(201)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestCurrenciesCreate(t *testing.T) {

	newCurrency := epcc.Currency{
		Type:              "currency",
		Code:              "INR",
		ExchangeRate:      142.15,
		Format:            "₹{price}",
		DecimalPoint:      ".",
		ThousandSeparator: ",",
		DecimalPlaces:     2,
		Default:           false,
		Enabled:           true,
	}

	expectedCurrencyData := epcc.CurrencyData{
		Data: epcc.Currency{
			ID:                "f8f0689e-4767-4924-b112-be89f490e1f5",
			Type:              "currency",
			Code:              "INR",
			ExchangeRate:      142.15,
			Format:            "₹{price}",
			DecimalPoint:      ".",
			ThousandSeparator: ",",
			DecimalPlaces:     2,
			Default:           false,
			Enabled:           true,
			Links: epcc.Links{
				Self: "https://api.moltin.com/currencies/f8f0689e-4767-4924-b112-be89f490e1f5",
			},
			Meta: epcc.CurrencyMeta{
				Timestamps: epcc.Timestamps{
					CreatedAt: "2020-09-01T15:48:10.050234331Z",
					UpdatedAt: "2020-09-01T15:48:10.050234395Z",
				},
			},
		},
	}

	tests := []struct {
		currency       epcc.Currency
		currenciesData epcc.CurrencyData
		err            error
	}{
		{newCurrency, expectedCurrencyData, nil},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleCurrenciesCreate))
	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		currenciesData, err := epcc.Currencies.Create(client, &newCurrency)
		assert.Equal(t, test.currenciesData, *currenciesData)
		assert.Equal(t, test.err, err)
	}
}
