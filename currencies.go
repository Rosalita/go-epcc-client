package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Currencies is used to access the Currencies endpoints.
var Currencies currencies

type currencies struct{}

// GetAll fetches all currencies
func (currencies) GetAll(client *Client) (*CurrenciesData, error) {

	path := fmt.Sprintf("/v2/currencies")

	body, err := client.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var currencies CurrenciesData
	if err := json.Unmarshal(body, &currencies); err != nil {
		return nil, err
	}

	return &currencies, nil
}

// Create creates a currency
func (currencies) Create(client *Client, currency *Currency) (*CurrencyData, error) {

	currencyData := CurrencyData{
		Data: *currency,
	}

	jsonPayload, err := json.Marshal(currencyData)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/v2/currencies")

	body, err := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))

	if err != nil {
		return nil, err
	}

	var newCurrency CurrencyData
	err = json.Unmarshal(body, &newCurrency)
	if err != nil {
		return nil, err
	}

	return &newCurrency, nil
}
