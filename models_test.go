package epcc_test

import (
	"encoding/json"
	"testing"

	"github.com/Rosalita/go-epcc-client"
	"github.com/stretchr/testify/assert"
)

func TestCurrenciesDataUnmarshalJSON(t *testing.T) {
	rawJSON := `{` +
		`"data":[{` +
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

	expectedCurrenciesData := epcc.CurrenciesData{
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

	var currenciesData epcc.CurrenciesData
	err := json.Unmarshal([]byte(rawJSON), &currenciesData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedCurrenciesData, currenciesData)
}
