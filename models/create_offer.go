package models

import (
    "encoding/json"
)

// CreateOffer represents a CreateOffer struct.
type CreateOffer struct {
    Name                string                 `json:"name"`
    Handle              string                 `json:"handle"`
    Description         *string                `json:"description,omitempty"`
    ProductId           int                    `json:"product_id"`
    ProductPricePointId *int                   `json:"product_price_point_id,omitempty"`
    Components          []CreateOfferComponent `json:"components,omitempty"`
    Coupons             []string               `json:"coupons,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOffer.
// It customizes the JSON marshaling process for CreateOffer objects.
func (c *CreateOffer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOffer object to a map representation for JSON marshaling.
func (c *CreateOffer) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["name"] = c.Name
    structMap["handle"] = c.Handle
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    structMap["product_id"] = c.ProductId
    if c.ProductPricePointId != nil {
        structMap["product_price_point_id"] = c.ProductPricePointId
    }
    if c.Components != nil {
        structMap["components"] = c.Components
    }
    if c.Coupons != nil {
        structMap["coupons"] = c.Coupons
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOffer.
// It customizes the JSON unmarshaling process for CreateOffer objects.
func (c *CreateOffer) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name                string                 `json:"name"`
        Handle              string                 `json:"handle"`
        Description         *string                `json:"description,omitempty"`
        ProductId           int                    `json:"product_id"`
        ProductPricePointId *int                   `json:"product_price_point_id,omitempty"`
        Components          []CreateOfferComponent `json:"components,omitempty"`
        Coupons             []string               `json:"coupons,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Name = temp.Name
    c.Handle = temp.Handle
    c.Description = temp.Description
    c.ProductId = temp.ProductId
    c.ProductPricePointId = temp.ProductPricePointId
    c.Components = temp.Components
    c.Coupons = temp.Coupons
    return nil
}
