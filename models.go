package epcc

// CurrencyData contains the data for a single currency
type CurrencyData struct {
	Data Currency `json:"data"`
}

// CurrenciesData contains the data for multiple currencies
type CurrenciesData struct {
	Data []Currency `json:"data"`
}

// Currency represents a currency
type Currency struct {
	ID                string       `json:"id,omitempty"`
	Type              string       `json:"type"`
	Code              string       `json:"code,omitempty"`
	ExchangeRate      float64      `json:"exchange_rate,omitempty"`
	Format            string       `json:"format,omitempty"`
	DecimalPoint      string       `json:"decimal_point,omitempty"`
	ThousandSeparator string       `json:"thousand_separator,omitempty"`
	DecimalPlaces     int64        `json:"decimal_places,omitempty"`
	Default           bool         `json:"default,omitempty"`
	Enabled           bool         `json:"enabled,omitempty"`
	Links             Links        `json:"links,omitempty"`
	Meta              CurrencyMeta `json:"meta,omitempty"`
}

// CurrencyMeta contains extra data for a currency
type CurrencyMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

// Links contains link information
type Links struct {
	Self string `json:"self"`
}

// Timestamps contains timestamp information
type Timestamps struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
