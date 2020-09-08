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
				Weight: &epcc.ProductWeight{
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

func fakeHandleProductsGet(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/v2/products/parentProductID" && req.Method == "GET":
		responseJSON := `{
			"data": {
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
					],
					"variation_matrix": {
						"f223eac4-2665-45a4-bc3c-e589ce6adadf": {
							"3f540af3-08e3-407c-be2c-8c7b1fce1fb7": "e261c5dd-e8f9-46dd-bbcb-7fffc2b79814",
							"690f7481-74e4-4c2b-965f-12de456fb1e1": "31ee57cc-4302-43ee-9e8e-906af5d8139f"
						}
					}
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
					"children": {
						"data": [
							{
								"type": "product",
								"id": "e261c5dd-e8f9-46dd-bbcb-7fffc2b79814"
							},
							{
								"type": "product",
								"id": "31ee57cc-4302-43ee-9e8e-906af5d8139f"
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
		}`
		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestProductsGet(t *testing.T) {
	expectedParentProduct := epcc.ProductData{
		Data: epcc.Product{
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
				VariationMatrix: epcc.ProductVariationMatrix{
					"f223eac4-2665-45a4-bc3c-e589ce6adadf": epcc.VariationOptionToChildProduct{
						"3f540af3-08e3-407c-be2c-8c7b1fce1fb7": "e261c5dd-e8f9-46dd-bbcb-7fffc2b79814",
						"690f7481-74e4-4c2b-965f-12de456fb1e1": "31ee57cc-4302-43ee-9e8e-906af5d8139f",
					},
				},
			},
			Weight: &epcc.ProductWeight{
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
						{
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
				Children: epcc.RelationshipItems{
					Data: []epcc.Relationship{
						{
							Type: "product",
							ID:   "e261c5dd-e8f9-46dd-bbcb-7fffc2b79814",
						},
						{
							Type: "product",
							ID:   "31ee57cc-4302-43ee-9e8e-906af5d8139f",
						},
					},
				},
			},
		},
	}

	tests := []struct {
		productID   string
		productData epcc.ProductData
		err         error
	}{
		{"parentProductID", expectedParentProduct, nil},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleProductsGet))
	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		productData, err := epcc.Products.Get(client, test.productID)
		assert.Equal(t, test.productData, *productData)
		assert.Equal(t, test.err, err)
	}
}

func fakeHandleProductsCreate(rw http.ResponseWriter, req *http.Request) {
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(req.Body)
	if err != nil {
		rw.WriteHeader(500)
		return
	}

	switch {
	case req.URL.String() == "/v2/products" && req.Method == "POST" && strings.Contains(buffer.String(), `"name":"Origami Crane"`):
		responseJSON := `{
			"data": {
				"type": "product",
				"id": "f8764ca8-aa1e-44db-ab5e-897ee96b3e9b",
				"name": "Origami Crane",
				"slug": "slug-origami-crane",
				"sku": "sku-origami-crane",
				"manage_stock": false,
				"description": "The Origami Crane is considered lucky.",
				"price": [
					{
						"amount": 1,
						"currency": "USD",
						"includes_tax": false
					}
				],
				"status": "draft",
				"commodity_type": "physical",
				"relationships": {},
				"meta": {
					"stock": {
						"level": 0,
						"availability": "out-stock"
					}
				}
			}		
		}`
		rw.WriteHeader(201)
		rw.Write([]byte(responseJSON))
	case req.URL.String() == "/v2/products" && req.Method == "POST" && strings.Contains(buffer.String(), `"name":"Invalid product"`):
		responseJSON := `{
			{
				"errors": [
					{
						"title": "Failed Validation",
						"detail": "The data.weight.unit field is required when data.weight is present."
					},
					{
						"title": "Failed Validation",
						"detail": "The data.weight.value field is required when data.weight is present."
					}
				]
			}
		}`
		rw.WriteHeader(422)
		rw.Write([]byte(responseJSON))
	default:
		rw.WriteHeader(500)
	}
}

func TestProductsCreate(t *testing.T) {
	validNewProduct := epcc.Product{
		Type:          "product",
		Name:          "Origami Crane",
		Slug:          "slug-origami-crane",
		SKU:           "sku-origami-crane",
		Description:   "The Origami Crane is considered lucky.",
		ManageStock:   false,
		Status:        "draft",
		CommodityType: "physical",
		Price: []epcc.ProductPrice{
			{
				Amount:      1,
				Currency:    "USD",
				IncludesTax: false,
			},
		},
	}

	invalidProductWeight := epcc.Product{
		Type:          "product",
		Name:          "Invalid product",
		Slug:          "slug-origami-crane",
		SKU:           "sku-origami-crane",
		Description:   "The Origami Crane is considered lucky.",
		ManageStock:   false,
		Status:        "draft",
		CommodityType: "physical",
		Price: []epcc.ProductPrice{
			{
				Amount:      1,
				Currency:    "USD",
				IncludesTax: false,
			},
		},
		Weight: &epcc.ProductWeight{
			Grams: 0,
		},
	}

	expectedProductData := epcc.ProductData{
		Data: epcc.Product{
			ID:            "f8764ca8-aa1e-44db-ab5e-897ee96b3e9b",
			Type:          "product",
			Name:          "Origami Crane",
			Slug:          "slug-origami-crane",
			SKU:           "sku-origami-crane",
			Description:   "The Origami Crane is considered lucky.",
			ManageStock:   false,
			Status:        "draft",
			CommodityType: "physical",
			Price: []epcc.ProductPrice{
				{
					Amount:      1,
					Currency:    "USD",
					IncludesTax: false,
				},
			},
			Meta: epcc.ProductMeta{
				Stock: epcc.ProductStock{Level: 0, Availability: "out-stock"},
			},
		},
	}

	tests := []struct {
		product     epcc.Product
		productData *epcc.ProductData
		err         error
	}{
		{validNewProduct, &expectedProductData, nil},
		{invalidProductWeight, nil, errors.New("status code 422 is not ok")},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(fakeHandleProductsCreate))
	options := epcc.ClientOptions{
		BaseURL:           testServer.URL,
		ClientTimeout:     10 * time.Second,
		RetryLimitTimeout: 10 * time.Millisecond,
	}
	client := epcc.NewClient(options)

	for _, test := range tests {
		productData, err := epcc.Products.Create(client, &test.product)
		if productData != nil {
			assert.Equal(t, test.productData, productData)
		}
		assert.Equal(t, test.err, err)
	}
}
