package epcc_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Rosalita/go-epcc-client"
	"github.com/stretchr/testify/assert"
)

func fakeHandleProductsGetAll(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/v2/products" && req.Method == "GET":
		responseJSON := `{
			"data": [
				{
					"type": "product",
					"id": "78ee7c20-df84-435d-bb1d-531e3537c4dc",
					"name": "Origami Frog",
					"slug": "FRG",
					"sku": "FRG",
					"manage_stock": true,
					"description": "An Origami Frog folded from one sheet of paper",
					"price": [
						{
							"amount": 1,
							"currency": "USD",
							"includes_tax": true
						}
					],
					"status": "draft",
					"commodity_type": "physical",
					"meta": {
						"timestamps": {
							"created_at": "2020-08-28T09:47:45+00:00",
							"updated_at": "2020-09-03T15:09:27+00:00"
						},
						"stock": {
							"level": 100,
							"availability": "in-stock"
						},
						"variations": [
							{
								"id": "10d2fd55-c059-41f4-b78f-1b915f3e7237",
								"name": "colour",
								"options": [
									{
										"id": "f223eac4-2665-45a4-bc3c-e589ce6adadf",
										"name": "Green",
										"description": "This is a nice colour"
									}
								]
							}
						]
					},
					"weight": {
						"g": 5,
						"kg": 0.005,
						"lb": 0.01102,
						"oz": 0.17637
					},
					"relationships": {
						"variations": {
							"data": [
								{
									"type": "product-variation",
									"id": "10d2fd55-c059-41f4-b78f-1b915f3e7237"
								}
							]
						},
						"main_image": {
							"data": {
								"type": "main_image",
								"id": "32d80649-687e-4e17-a614-a3d612b07ced"
							}
						}
					}
				}
			],
			"links": {
				"current": "https://api.moltin.com/v2/products?page[limit]=100&page[offset]=0",
				"first": "https://api.moltin.com/v2/products?page[limit]=100&page[offset]=0",
				"last": null
			},
			"meta": {
				"results": {
					"total": 1
				},
				"page": {
					"limit": 100,
					"offset": 0,
					"current": 1,
					"total": 1
				}
			}		
		}`
		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestProductsGetAll(t *testing.T) {
	expectedProducts := epcc.ProductsData{
		Data: []epcc.Product{
			{
				ID:            "78ee7c20-df84-435d-bb1d-531e3537c4dc",
				Type:          "product",
				Name:          "Origami Frog",
				Slug:          "FRG",
				SKU:           "FRG",
				Description:   "An Origami Frog folded from one sheet of paper",
				ManageStock:   true,
				Status:        "draft",
				CommodityType: "physical",
				Price: []epcc.ProductPrice{
					{
						Amount:      1,
						Currency:    "USD",
						IncludesTax: true,
					},
				},
				Meta: epcc.ProductMeta{
					Timestamps: epcc.Timestamps{
						CreatedAt: "2020-08-28T09:47:45+00:00",
						UpdatedAt: "2020-09-03T15:09:27+00:00",
					},
					Stock: epcc.ProductStock{
						Level:        100,
						Availability: "in-stock",
					},
					Variations: []epcc.ProductVariation{
						{
							ID:   "10d2fd55-c059-41f4-b78f-1b915f3e7237",
							Name: "colour",
							Options: []epcc.ProductVariationOptions{
								epcc.ProductVariationOptions{
									Name:        "Green",
									Description: "This is a nice colour",
								},
							},
						},
					},
				},
				Weight: epcc.ProductWeight{
					Grams:     5,
					Kilograms: 0.005,
					Pounds:    0.01102,
					Ounces:    0.17637,
				},
				Relationships: epcc.ProductRelationships{
					Variations: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{Type: "product-variation",
								ID: "10d2fd55-c059-41f4-b78f-1b915f3e7237",
							},
						},
					},
					MainImage: epcc.RelationshipItem{
						Data: epcc.Relationship{
							Type: "main_image",
							ID:   "32d80649-687e-4e17-a614-a3d612b07ced",
						},
					},
				},
			},
		},
	}

	tests := []struct {
		productsData epcc.ProductsData
		err          error
	}{
		{expectedProducts, nil},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleProductsGetAll))
	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		productsData, err := epcc.Products.GetAll(client)
		assert.Equal(t, test.productsData, *productsData)
		assert.Equal(t, test.err, err)
	}
}
