package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Currencies is used to access the Currencies endpoints.
var Currencies currencies

type currencies struct{}

// Get fetches a single currency
func (currencies) Get(client *Client, currencyID string) (*CurrencyData, error) {
	path := fmt.Sprintf("/v2/currencies/%s", currencyID)

	body, err := client.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var currencies CurrencyData
	if err := json.Unmarshal(body, &currencies); err != nil {
		return nil, err
	}

	return &currencies, nil
}

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
	if err := json.Unmarshal(body, &newCurrency); err != nil {
		return nil, err
	}

	return &newCurrency, nil
}

// Delete deletes a currency.
func (currencies) Delete(client *Client, currencyID string) error {
	path := fmt.Sprintf("/v2/currencies/%s", currencyID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a currency.
func (currencies) Update(client *Client, currencyID string, currency *Currency) (*CurrencyData, error) {

	currencyData := CurrencyData{
		Data: *currency,
	}

	jsonPayload, err := json.Marshal(currencyData)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/v2/currencies/%s", currencyID)

	body, err := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	var updatedCurrency CurrencyData
	if err := json.Unmarshal(body, &updatedCurrency); err != nil {
		return nil, err
	}

	return &updatedCurrency, nil
}
