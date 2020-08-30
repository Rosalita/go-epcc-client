package epcc

// CurrenciesData contains the data for currencies
type CurrenciesData struct {
	Data []Currency `json:"data"`
}

// Currency contains the data for a currency
type Currency struct {
	ID                string       `json:"id,omitempty"`
	Type              string       `json:"type"`
	Code              string       `json:"code"`
	ExchangeRate      float64      `json:"exchange_rate"`
	Format            string       `json:"format"`
	DecimalPoint      string       `json:"decimal_point"`
	ThousandSeparator string       `json:"thousand_separator"`
	DecimalPlaces     int64        `json:"decimal_places"`
	Default           bool         `json:"default"`
	Enabled           bool         `json:"enabled"`
	Links             Links        `json:"links"`
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
