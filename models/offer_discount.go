package models

import (
    "encoding/json"
)

// OfferDiscount represents a OfferDiscount struct.
type OfferDiscount struct {
    CouponCode *string `json:"coupon_code,omitempty"`
    CouponId   *int    `json:"coupon_id,omitempty"`
    CouponName *string `json:"coupon_name,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for OfferDiscount.
// It customizes the JSON marshaling process for OfferDiscount objects.
func (o *OfferDiscount) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OfferDiscount object to a map representation for JSON marshaling.
func (o *OfferDiscount) toMap() map[string]any {
    structMap := make(map[string]any)
    if o.CouponCode != nil {
        structMap["coupon_code"] = o.CouponCode
    }
    if o.CouponId != nil {
        structMap["coupon_id"] = o.CouponId
    }
    if o.CouponName != nil {
        structMap["coupon_name"] = o.CouponName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OfferDiscount.
// It customizes the JSON unmarshaling process for OfferDiscount objects.
func (o *OfferDiscount) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CouponCode *string `json:"coupon_code,omitempty"`
        CouponId   *int    `json:"coupon_id,omitempty"`
        CouponName *string `json:"coupon_name,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    o.CouponCode = temp.CouponCode
    o.CouponId = temp.CouponId
    o.CouponName = temp.CouponName
    return nil
}
