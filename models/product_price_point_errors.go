package models

import (
    "encoding/json"
)

// ProductPricePointErrors represents a ProductPricePointErrors struct.
type ProductPricePointErrors struct {
    PricePoint   *string  `json:"price_point,omitempty"`
    Interval     []string `json:"interval,omitempty"`
    IntervalUnit []string `json:"interval_unit,omitempty"`
    Name         []string `json:"name,omitempty"`
    Price        []string `json:"price,omitempty"`
    PriceInCents []string `json:"price_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePointErrors.
// It customizes the JSON marshaling process for ProductPricePointErrors objects.
func (p *ProductPricePointErrors) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePointErrors object to a map representation for JSON marshaling.
func (p *ProductPricePointErrors) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.PricePoint != nil {
        structMap["price_point"] = p.PricePoint
    }
    if p.Interval != nil {
        structMap["interval"] = p.Interval
    }
    if p.IntervalUnit != nil {
        structMap["interval_unit"] = p.IntervalUnit
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.Price != nil {
        structMap["price"] = p.Price
    }
    if p.PriceInCents != nil {
        structMap["price_in_cents"] = p.PriceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePointErrors.
// It customizes the JSON unmarshaling process for ProductPricePointErrors objects.
func (p *ProductPricePointErrors) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricePoint   *string  `json:"price_point,omitempty"`
        Interval     []string `json:"interval,omitempty"`
        IntervalUnit []string `json:"interval_unit,omitempty"`
        Name         []string `json:"name,omitempty"`
        Price        []string `json:"price,omitempty"`
        PriceInCents []string `json:"price_in_cents,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.PricePoint = temp.PricePoint
    p.Interval = temp.Interval
    p.IntervalUnit = temp.IntervalUnit
    p.Name = temp.Name
    p.Price = temp.Price
    p.PriceInCents = temp.PriceInCents
    return nil
}
