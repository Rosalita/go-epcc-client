package epcc

import (
	"encoding/json"
	"fmt"
)

// Products is used to access the Products endpoints.
var Products products

type products struct{}

// GetAll fetches all products
func (products) GetAll(client *Client) (*ProductsData, error) {
	path := fmt.Sprintf("/v2/products")

	body, err := client.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var products ProductsData
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, err
	}

	return &products, nil
}
