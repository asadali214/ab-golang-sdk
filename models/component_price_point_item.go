package models

import (
    "encoding/json"
)

// ComponentPricePointItem represents a ComponentPricePointItem struct.
type ComponentPricePointItem struct {
    Name          *string        `json:"name,omitempty"`
    Handle        *string        `json:"handle,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme *PricingScheme `json:"pricing_scheme,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval      *int           `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit  *IntervalUnit  `json:"interval_unit,omitempty"`
    Prices        []Price        `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointItem.
// It customizes the JSON marshaling process for ComponentPricePointItem objects.
func (c *ComponentPricePointItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointItem object to a map representation for JSON marshaling.
func (c *ComponentPricePointItem) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Handle != nil {
        structMap["handle"] = c.Handle
    }
    if c.PricingScheme != nil {
        structMap["pricing_scheme"] = c.PricingScheme
    }
    if c.Interval != nil {
        structMap["interval"] = c.Interval
    }
    if c.IntervalUnit != nil {
        structMap["interval_unit"] = c.IntervalUnit
    }
    if c.Prices != nil {
        structMap["prices"] = c.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointItem.
// It customizes the JSON unmarshaling process for ComponentPricePointItem objects.
func (c *ComponentPricePointItem) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name          *string        `json:"name,omitempty"`
        Handle        *string        `json:"handle,omitempty"`
        PricingScheme *PricingScheme `json:"pricing_scheme,omitempty"`
        Interval      *int           `json:"interval,omitempty"`
        IntervalUnit  *IntervalUnit  `json:"interval_unit,omitempty"`
        Prices        []Price        `json:"prices,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Name = temp.Name
    c.Handle = temp.Handle
    c.PricingScheme = temp.PricingScheme
    c.Interval = temp.Interval
    c.IntervalUnit = temp.IntervalUnit
    c.Prices = temp.Prices
    return nil
}
