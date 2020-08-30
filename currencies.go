package epcc

import (
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
