package epcc_test

import (
	"encoding/json"
	"testing"

	"github.com/Rosalita/go-epcc-client"
	"github.com/stretchr/testify/assert"
)

func TestMultipleProductsDataUnmarshal(t *testing.T) {
	rawJSON := `{
		"data": [
			{
				"type": "product",
				"id": "31ee57cc-4302-43ee-9e8e-906af5d8139f",
				"name": "Origami Frog",
				"slug": "LARGE-GREEN-FROG",
				"sku": "LARGE-GREEN-FROG",
				"manage_stock": true,
				"description": "An Origami Frog folded from one sheet of paper",
				"price": [
					{
						"amount": 1,
						"currency": "USD",
						"includes_tax": true
					}
				],
				"status": "live",
				"commodity_type": "physical",
				"meta": {
					"timestamps": {
						"created_at": "2020-09-04T11:01:06+00:00",
						"updated_at": "2020-09-04T11:10:38+00:00"
					},
					"stock": {
						"level": 0,
						"availability": "out-stock"
					}
				},
				"weight": {
					"g": 5,
					"kg": 0.005,
					"lb": 0.01102,
					"oz": 0.17637
				},
				"relationships": {
					"categories": {
						"data": [
							{
								"type": "category",
								"id": "162fbded-fb6a-4f6b-9232-d3d089d6df24"
							}
						]
					},
					"brands": {
						"data": [
							{
								"type": "brand",
								"id": "6c1cb310-5962-4c6f-ab52-503e1831e094"
							}
						]
					},
					"parent": {
						"data": {
							"type": "product",
							"id": "78ee7c20-df84-435d-bb1d-531e3537c4dc"
						}
					}
				}
			},
			{
				"type": "product",
				"id": "e261c5dd-e8f9-46dd-bbcb-7fffc2b79814",
				"name": "Origami Frog",
				"slug": "SMALL-GREEN-FROG",
				"sku": "SMALL-GREEN-FROG",
				"manage_stock": true,
				"description": "An Origami Frog folded from one sheet of paper",
				"price": [
					{
						"amount": 1,
						"currency": "USD",
						"includes_tax": true
					}
				],
				"status": "live",
				"commodity_type": "physical",
				"meta": {
					"timestamps": {
						"created_at": "2020-09-04T11:01:06+00:00",
						"updated_at": "2020-09-04T11:10:29+00:00"
					},
					"stock": {
						"level": 0,
						"availability": "out-stock"
					}
				},
				"weight": {
					"g": 5,
					"kg": 0.005,
					"lb": 0.01102,
					"oz": 0.17637
				},
				"relationships": {
					"categories": {
						"data": [
							{
								"type": "category",
								"id": "162fbded-fb6a-4f6b-9232-d3d089d6df24"
							}
						]
					},
					"brands": {
						"data": [
							{
								"type": "brand",
								"id": "6c1cb310-5962-4c6f-ab52-503e1831e094"
							}
						]
					},
					"parent": {
						"data": {
							"type": "product",
							"id": "78ee7c20-df84-435d-bb1d-531e3537c4dc"
						}
					}
				}
			},
			{
				"type": "product",
				"id": "78ee7c20-df84-435d-bb1d-531e3537c4dc",
				"name": "Origami Frog",
				"slug": "{SIZE}-{COLOUR}-FROG",
				"sku": "{SIZE}-{COLOUR}-FROG",
				"manage_stock": true,
				"description": "An Origami Frog folded from one sheet of paper",
				"price": [
					{
						"amount": 1,
						"currency": "USD",
						"includes_tax": true
					}
				],
				"status": "live",
				"commodity_type": "physical",
				"meta": {
					"timestamps": {
						"created_at": "2020-08-28T09:47:45+00:00",
						"updated_at": "2020-09-04T11:10:08+00:00"
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
						},
						{
							"id": "3326da08-d5c8-45e6-8e26-2db560ed7c35",
							"name": "Size",
							"options": [
								{
									"id": "3f540af3-08e3-407c-be2c-8c7b1fce1fb7",
									"name": "Small",
									"description": "Not very big"
								},
								{
									"id": "690f7481-74e4-4c2b-965f-12de456fb1e1",
									"name": "Large",
									"description": "Bigger than small"
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
					"files": {
						"data": [
							{
								"type": "file",
								"id": "271073e3-7fbe-45bf-8dcb-071bd10350cd"
							}
						]
					},
					"categories": {
						"data": [
							{
								"type": "category",
								"id": "162fbded-fb6a-4f6b-9232-d3d089d6df24"
							}
						]
					},
					"collections": {
						"data": [
							{
								"type": "collection",
								"id": "5bf1bd12-746f-4a81-9bf4-4b8b690b02cc"
							}
						]
					},
					"brands": {
						"data": [
							{
								"type": "brand",
								"id": "6c1cb310-5962-4c6f-ab52-503e1831e094"
							}
						]
					},
					"variations": {
						"data": [
							{
								"type": "product-variation",
								"id": "10d2fd55-c059-41f4-b78f-1b915f3e7237"
							},
							{
								"type": "product-variation",
								"id": "3326da08-d5c8-45e6-8e26-2db560ed7c35"
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
				"total": 3
			},
			"page": {
				"limit": 100,
				"offset": 0,
				"current": 1,
				"total": 1
			}
		}
	}`

	expectedProductsData := epcc.ProductsData{
		Data: []epcc.Product{
			{
				ID:          "31ee57cc-4302-43ee-9e8e-906af5d8139f",
				Type:        "product",
				Name:        "Origami Frog",
				Slug:        "LARGE-GREEN-FROG",
				SKU:         "LARGE-GREEN-FROG",
				Description: "An Origami Frog folded from one sheet of paper",
				ManageStock: true, Status: "live",
				CommodityType: "physical",
				Price: []epcc.ProductPrice{
					{Amount: 1,
						Currency:    "USD",
						IncludesTax: true,
					},
				},
				Meta: epcc.ProductMeta{
					Timestamps: epcc.Timestamps{
						CreatedAt: "2020-09-04T11:01:06+00:00",
						UpdatedAt: "2020-09-04T11:10:38+00:00",
					},
					Stock: epcc.ProductStock{
						Level:        0,
						Availability: "out-stock",
					},
				},
				Weight: epcc.ProductWeight{
					Grams:     5,
					Kilograms: 0.005,
					Pounds:    0.01102,
					Ounces:    0.17637,
				},
				Relationships: epcc.ProductRelationships{
					Categories: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "category",
								ID:   "162fbded-fb6a-4f6b-9232-d3d089d6df24",
							},
						},
					},
					Brands: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "brand",
								ID:   "6c1cb310-5962-4c6f-ab52-503e1831e094",
							},
						},
					},
					Parent: epcc.RelationshipItem{
						Data: epcc.Relationship{
							Type: "product",
							ID:   "78ee7c20-df84-435d-bb1d-531e3537c4dc",
						},
					},
				},
			},
			{
				ID:            "e261c5dd-e8f9-46dd-bbcb-7fffc2b79814",
				Type:          "product",
				Name:          "Origami Frog",
				Slug:          "SMALL-GREEN-FROG",
				SKU:           "SMALL-GREEN-FROG",
				Description:   "An Origami Frog folded from one sheet of paper",
				ManageStock:   true,
				Status:        "live",
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
						CreatedAt: "2020-09-04T11:01:06+00:00",
						UpdatedAt: "2020-09-04T11:10:29+00:00",
					},
					Stock: epcc.ProductStock{
						Level:        0,
						Availability: "out-stock",
					},
				},
				Weight: epcc.ProductWeight{
					Grams:     5,
					Kilograms: 0.005,
					Pounds:    0.01102,
					Ounces:    0.17637,
				},
				Relationships: epcc.ProductRelationships{
					Categories: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "category",
								ID:   "162fbded-fb6a-4f6b-9232-d3d089d6df24",
							},
						},
					},
					Brands: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "brand",
								ID:   "6c1cb310-5962-4c6f-ab52-503e1831e094",
							},
						},
					},
					Parent: epcc.RelationshipItem{
						Data: epcc.Relationship{
							Type: "product",
							ID:   "78ee7c20-df84-435d-bb1d-531e3537c4dc",
						},
					},
				},
			},
			{
				ID:            "78ee7c20-df84-435d-bb1d-531e3537c4dc",
				Type:          "product",
				Name:          "Origami Frog",
				Slug:          "{SIZE}-{COLOUR}-FROG",
				SKU:           "{SIZE}-{COLOUR}-FROG",
				Description:   "An Origami Frog folded from one sheet of paper",
				ManageStock:   true,
				Status:        "live",
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
						UpdatedAt: "2020-09-04T11:10:08+00:00",
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
						{
							ID:   "3326da08-d5c8-45e6-8e26-2db560ed7c35",
							Name: "Size",
							Options: []epcc.ProductVariationOptions{
								{
									Name:        "Small",
									Description: "Not very big",
								},
								{
									Name:        "Large",
									Description: "Bigger than small",
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
					Files: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "file",
								ID:   "271073e3-7fbe-45bf-8dcb-071bd10350cd",
							},
						},
					},
					Categories: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "category",
								ID:   "162fbded-fb6a-4f6b-9232-d3d089d6df24",
							},
						},
					},
					Collections: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "collection",
								ID:   "5bf1bd12-746f-4a81-9bf4-4b8b690b02cc",
							},
						},
					},
					Brands: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "brand",
								ID:   "6c1cb310-5962-4c6f-ab52-503e1831e094",
							},
						},
					},
					Variations: epcc.RelationshipItems{
						Data: []epcc.Relationship{
							{
								Type: "product-variation",
								ID:   "10d2fd55-c059-41f4-b78f-1b915f3e7237",
							},
							epcc.Relationship{
								Type: "product-variation",
								ID:   "3326da08-d5c8-45e6-8e26-2db560ed7c35",
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

	var productsData epcc.ProductsData
	err := json.Unmarshal([]byte(rawJSON), &productsData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedProductsData, productsData)
}
