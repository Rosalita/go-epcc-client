package epcc

// ProductData contains the data for a single products
type ProductData struct {
	Data Product `json:"data"`
}

// ProductsData contains the data for multiple products
type ProductsData struct {
	Data []Product `json:"data"`
}

// Product represents a product
type Product struct {
	ID            string                 `json:"id,omitempty"`
	Type          string                 `json:"type"`
	Name          string                 `json:"name"`
	Slug          string                 `json:"slug"`
	SKU           string                 `json:"sku"`
	Description   string                 `json:"description"`
	ManageStock   bool                   `json:"manage_stock"`
	Status        string                 `json:"status"`
	CommodityType string                 `json:"commodity_type"`
	Price         []ProductPrice         `json:"price"`
	Meta          ProductMeta            `json:"meta,omitempty"`
	Weight        *ProductWeight         `json:"weight,omitempty"`
	Relationships ProductRelationships   `json:"relationships,omitempty"`
	Dimensions    map[string]Measurement `json:"dimensions,omitempty"`
}

// Measurement represents a measurement
type Measurement struct {
	Measurement string  `json:"measurement"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

// ProductPrice is a price for a Products meta
type ProductPrice struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	IncludesTax bool   `json:"includes_tax"`
}

// ProductMeta contains extra data for a product
type ProductMeta struct {
	Timestamps      Timestamps             `json:"timestamps,omitempty"`
	Stock           ProductStock           `json:"stock"`
	Variations      []ProductVariation     `json:"variations,omitempty"`
	VariationMatrix ProductVariationMatrix `json:"variation_matrix"`
}

//ProductVariationMatrix is a map of variationID's to VariationOptions and child product IDs
type ProductVariationMatrix map[string]VariationOptionToChildProduct

// VariationOptionToChildProduct is a map of variationOptionIDs to child productIDs
type VariationOptionToChildProduct map[string]string

// ProductVariation is a variation object for a ProductMeta
type ProductVariation struct {
	ID      string                    `json:"id"`
	Name    string                    `json:"name"`
	Options []ProductVariationOptions `json:"options"`
}

// ProductVariationOptions is a options object for a Products ProductVariation
type ProductVariationOptions struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ProductStock is a stock object for a Products meta
type ProductStock struct {
	Level        int    `json:"level"`
	Availability string `json:"availability"`
}

// ProductWeight represents the weight of a product
type ProductWeight struct {
	Grams     int     `json:"g"`
	Kilograms float64 `json:"kg"`
	Pounds    float64 `json:"lb"`
	Ounces    float64 `json:"oz"`
}

// ProductRelationships represents the relationships that can exist for a product
type ProductRelationships struct {
	Files       RelationshipItems `json:"files,omitempty"`
	Categories  RelationshipItems `json:"categories,omitempty"`
	Collections RelationshipItems `json:"collections,omitempty"`
	Brands      RelationshipItems `json:"brands,omitempty"`
	Variations  RelationshipItems `json:"variations,omitempty"`
	MainImage   RelationshipItem  `json:"main_image,omitempty"`
	Parent      RelationshipItem  `json:"parent,omitempty"`
	Children    RelationshipItems `json:"children,omitempty"`
}

// RelationshipItems holds multiple items of data about a relationship
type RelationshipItems struct {
	Data []Relationship `json:"data,omitempty"`
}

// RelationshipItem holds a single item of data about a relationship
type RelationshipItem struct {
	Data Relationship `json:"data,omitempty"`
}

// Relationship is a single item of information about a relationship
type Relationship struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
