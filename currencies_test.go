package epcc_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Rosalita/go-epcc-client"
	"github.com/stretchr/testify/assert"
)

func fakeHandleCurrenciesGet(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/v2/currencies/validCurrencyID" && req.Method == "GET":
		responseJSON := `{
			"data":{
				"id":"3563bde2-fb72-4721-8584-504058f63780",
				"type":"currency",
				"code":"GBP",
				"exchange_rate":1.33,
				"format":"£{price}",
				"decimal_point":".",
				"thousand_separator":",",
				"decimal_places":2,
				"default":true,
				"enabled":true,
				"links":{
					"self":"https://api.moltin.com/currencies/9ea32a77-982b-41fd-873c-62570bbfc1e0"
				},
				"meta":{
					"timestamps":{
						"created_at":"2020-08-28T09:56:36.852Z",
						"updated_at":"2020-08-28T09:56:36.852Z"
					}
				}
			}
		}`

		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestCurrenciesGet(t *testing.T) {
	expectedCurrencyData := epcc.CurrencyData{
		Data: epcc.Currency{
			ID:                "3563bde2-fb72-4721-8584-504058f63780",
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
	}

	tests := []struct {
		currencyID   string
		currencyData epcc.CurrencyData
		err          error
	}{
		{"validCurrencyID", expectedCurrencyData, nil},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleCurrenciesGet))
	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		currencyData, err := epcc.Currencies.Get(client, test.currencyID)
		assert.Equal(t, test.currencyData, *currencyData)
		assert.Equal(t, test.err, err)
	}
}

func fakeHandleCurrenciesGetAll(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/v2/currencies" && req.Method == "GET":
		responseJSON := `{
			"data":[{
				"id":"9ea32a77-982b-41fd-873c-62570bbfc1e0",
				"type":"currency",
				"code":"GBP",
				"exchange_rate":1.33,
				"format":"£{price}",
				"decimal_point":".",
				"thousand_separator":",",
				"decimal_places":2,
				"default":true,
				"enabled":true,
				"links":{
					"self":"https://api.moltin.com/currencies/9ea32a77-982b-41fd-873c-62570bbfc1e0"
				},
				"meta":{
					"timestamps":{
						"created_at":"2020-08-28T09:56:36.852Z",
						"updated_at":"2020-08-28T09:56:36.852Z"
					}
				}
			},{
				"id":"9cc80fcd-115c-47ba-9ff0-07072b378ded",
				"type":"currency",
				"code":"USD",
				"exchange_rate":1,
				"format":"${price}",
				"decimal_point":".",
				"thousand_separator":",",
				"decimal_places":2,
				"default":false,
				"enabled":true,
				"links":{
					"self":"https://api.moltin.com/currencies/9cc80fcd-115c-47ba-9ff0-07072b378ded"
				},
				"meta":{
					"timestamps":{
						"created_at":"2020-08-28T09:45:11.822Z",
						"updated_at":"2020-08-28T09:45:11.822Z"
					}
				}
			}]
		}`

		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestCurrenciesGetAll(t *testing.T) {
	expectedCurrencies := epcc.CurrenciesData{
		Data: []epcc.Currency{
			{
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
			{
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
		responseJSON := `{
			"data": {
				"id":"f8f0689e-4767-4924-b112-be89f490e1f5",
				"type":"currency",
				"code":"INR",
				"exchange_rate":142.15,
				"format":"₹{price}",
				"decimal_point":".",
				"thousand_separator":",",
				"decimal_places":2,
				"default":false,
				"enabled":true,
				"links":{
					"self":"https://api.moltin.com/currencies/f8f0689e-4767-4924-b112-be89f490e1f5"
				},
				"meta":{
					"timestamps":{
						"created_at":"2020-09-01T15:48:10.050234331Z",
						"updated_at":"2020-09-01T15:48:10.050234395Z"
					}
				}
			}
		}`

		rw.WriteHeader(201)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/v2/currencies" && req.Method == "POST" && strings.Contains(buffer.String(), `"code":"EUR"`):
		responseJSON := `{
			"errors":[{
				"status":400,
				"title":"Currency already exists",
				"detail":"The specified currency code already exists for this store"
			}]
		}`
		rw.WriteHeader(400)
		rw.Write([]byte(responseJSON))

	default:
		rw.WriteHeader(500)
	}
}

func TestCurrenciesCreate(t *testing.T) {
	validNewCurrency := epcc.Currency{
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

	currencyAlreadyExists := epcc.Currency{
		Type:              "currency",
		Code:              "EUR",
		ExchangeRate:      1.13,
		Format:            "€{price}",
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
		currency     epcc.Currency
		currencyData *epcc.CurrencyData
		err          error
	}{
		{validNewCurrency, &expectedCurrencyData, nil},
		{currencyAlreadyExists, nil, errors.New("status code 400 is not ok")},
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
		currencyData, err := epcc.Currencies.Create(client, &test.currency)
		if currencyData != nil {
			assert.Equal(t, test.currencyData, currencyData)
		}
		assert.Equal(t, test.err, err)
	}
}

func fakeHandleCurrenciesDelete(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/v2/currencies/validCurrencyID" && req.Method == "DELETE":
		responseJSON := `{}`
		rw.WriteHeader(204)
		rw.Write([]byte(responseJSON))
	case req.URL.String() == "/v2/currencies/notFound" && req.Method == "DELETE":
		responseJSON := `{
			"errors":[{
				"status":404,
				"title":"Currency not found",
				"detail":"The requested currency could not be found"
			}]
		}`
		rw.WriteHeader(404)
		rw.Write([]byte(responseJSON))
	case req.URL.String() == "/v2/currencies/defaultCurrency" && req.Method == "DELETE":
		responseJSON := `{
			"errors":[{
				"status":400,
				"title":"Cannot delete default currency",
				"detail":"Make another currency default before removing"
			}]
		}`
		rw.WriteHeader(400)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestCurrenciesDelete(t *testing.T) {
	tests := []struct {
		currencyID string
		err        error
	}{
		{"validCurrencyID", nil},
		{"notFound", errors.New("status code 404 is not ok")},
		{"defaultCurrency", errors.New("status code 400 is not ok")},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleCurrenciesDelete))

	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		err := epcc.Currencies.Delete(client, test.currencyID)
		assert.Equal(t, test.err, err)
	}
}

func fakeHandleCurrenciesUpdate(rw http.ResponseWriter, req *http.Request) {
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(req.Body)
	if err != nil {
		rw.WriteHeader(500)
		return
	}

	switch {
	case req.URL.String() == "/v2/currencies/largeUpdate" && req.Method == "PUT" && strings.Contains(buffer.String(), `"exchange_rate":1.13`):
		responseJSON := `{
			"data": {
				"id":"3563bde2-fb72-4721-8584-504058f63780",
				"type":"currency",
				"code":"EUR",
				"exchange_rate":1.13,
				"format":"€{price}",
				"decimal_point":".",
				"thousand_separator":",",
				"decimal_places":2,
				"default":false,
				"enabled":true,
				"links":{
					"self":"https://api.moltin.com/currencies/3563bde2-fb72-4721-8584-504058f63780"
				},
				"meta":{
					"timestamps":{
						"created_at":"2020-09-01T15:39:18.273Z",
						"updated_at":"2020-09-02T14:38:03.120282079Z"
					}
				}
			}
		}`
		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))
	case req.URL.String() == "/v2/currencies/smallUpdate" && req.Method == "PUT" && strings.Contains(buffer.String(), `"exchange_rate":1.14`):
		responseJSON := `{
			"data": {
				"id":"3563bde2-fb72-4721-8584-504058f63780",
				"type":"currency",
				"code":"EUR",
				"exchange_rate":1.14,
				"format":"€{price}",
				"decimal_point":".",
				"thousand_separator":",",
				"decimal_places":2,
				"default":false,
				"enabled":true,
				"links":{
					"self":"https://api.moltin.com/currencies/3563bde2-fb72-4721-8584-504058f63780"
				},
				"meta":{
					"timestamps":{
						"created_at":"2020-09-01T15:39:18.273Z",
						"updated_at":"2020-09-02T14:38:03.120282079Z"
					}
				}
			}
		}`
		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))

	default:
		rw.WriteHeader(500)
	}
}

func TestCurrenciesUpdate(t *testing.T) {
	largeUpdate := epcc.Currency{
		Type:              "currency",
		Code:              "EUR",
		ExchangeRate:      1.13,
		Format:            "€{price}",
		DecimalPoint:      ".",
		ThousandSeparator: ",",
		DecimalPlaces:     2,
		Default:           false,
		Enabled:           true,
	}

	expectedLargeUpdateCurrencyData := epcc.CurrencyData{
		Data: epcc.Currency{
			ID:                "3563bde2-fb72-4721-8584-504058f63780",
			Type:              "currency",
			Code:              "EUR",
			ExchangeRate:      1.13,
			Format:            "€{price}",
			DecimalPoint:      ".",
			ThousandSeparator: ",",
			DecimalPlaces:     2,
			Default:           false,
			Enabled:           true,
			Links: epcc.Links{
				Self: "https://api.moltin.com/currencies/3563bde2-fb72-4721-8584-504058f63780",
			},
			Meta: epcc.CurrencyMeta{
				Timestamps: epcc.Timestamps{
					CreatedAt: "2020-09-01T15:39:18.273Z",
					UpdatedAt: "2020-09-02T14:38:03.120282079Z",
				},
			},
		},
	}

	smallUpdate := epcc.Currency{
		Type:         "currency",
		ExchangeRate: 1.14,
	}

	expectedSmallUpdateCurrencyData := epcc.CurrencyData{
		Data: epcc.Currency{
			ID:                "3563bde2-fb72-4721-8584-504058f63780",
			Type:              "currency",
			Code:              "EUR",
			ExchangeRate:      1.14,
			Format:            "€{price}",
			DecimalPoint:      ".",
			ThousandSeparator: ",",
			DecimalPlaces:     2,
			Default:           false,
			Enabled:           true,
			Links: epcc.Links{
				Self: "https://api.moltin.com/currencies/3563bde2-fb72-4721-8584-504058f63780",
			},
			Meta: epcc.CurrencyMeta{
				Timestamps: epcc.Timestamps{
					CreatedAt: "2020-09-01T15:39:18.273Z",
					UpdatedAt: "2020-09-02T14:38:03.120282079Z",
				},
			},
		},
	}

	tests := []struct {
		currencyID   string
		update       epcc.Currency
		currencyData *epcc.CurrencyData
		err          error
	}{
		{"largeUpdate", largeUpdate, &expectedLargeUpdateCurrencyData, nil},
		{"smallUpdate", smallUpdate, &expectedSmallUpdateCurrencyData, nil},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleCurrenciesUpdate))
	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		currencyData, err := epcc.Currencies.Update(client, test.currencyID, &test.update)
		if currencyData != nil {
			assert.Equal(t, test.currencyData, currencyData)
		}
		assert.Equal(t, test.err, err)
	}
}
