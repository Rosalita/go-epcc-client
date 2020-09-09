package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"errors"
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

// Get fetches a single product
func (products) Get(client *Client, productID string) (*ProductData, error) {
	path := fmt.Sprintf("/v2/products/%s", productID)

	body, err := client.DoRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var product ProductData
	if err := json.Unmarshal(body, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

// Create creates a product
func (products) Create(client *Client, product *Product) (*ProductData, error) {

	productData := ProductData{
		Data: *product,
	}

	jsonPayload, err := json.Marshal(productData)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/v2/products")

	body, err := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	var newProduct ProductData
	if err := json.Unmarshal(body, &newProduct); err != nil {
		return nil, err
	}

	return &newProduct, nil
}

// Update updates a product.
func (products) Update(client *Client, product *Product) (*ProductData, error) {

	if product.ID == "" {
		return nil, errors.New("error productID is required")
	}

	productData := ProductData{
		Data: *product,
	}

	jsonPayload, err := json.Marshal(productData)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/v2/products/%s", product.ID)

	body, err := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	var updatedProduct ProductData
	if err := json.Unmarshal(body, &updatedProduct); err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}
