package models

import (
    "encoding/json"
)

// PaymentForAllocation represents a PaymentForAllocation struct.
// Information for captured payment, if applicable
type PaymentForAllocation struct {
    Id            *int    `json:"id,omitempty"`
    AmountInCents *int64  `json:"amount_in_cents,omitempty"`
    Success       *bool   `json:"success,omitempty"`
    Memo          *string `json:"memo,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentForAllocation.
// It customizes the JSON marshaling process for PaymentForAllocation objects.
func (p *PaymentForAllocation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentForAllocation object to a map representation for JSON marshaling.
func (p *PaymentForAllocation) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.AmountInCents != nil {
        structMap["amount_in_cents"] = p.AmountInCents
    }
    if p.Success != nil {
        structMap["success"] = p.Success
    }
    if p.Memo != nil {
        structMap["memo"] = p.Memo
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentForAllocation.
// It customizes the JSON unmarshaling process for PaymentForAllocation objects.
func (p *PaymentForAllocation) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id            *int    `json:"id,omitempty"`
        AmountInCents *int64  `json:"amount_in_cents,omitempty"`
        Success       *bool   `json:"success,omitempty"`
        Memo          *string `json:"memo,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Id = temp.Id
    p.AmountInCents = temp.AmountInCents
    p.Success = temp.Success
    p.Memo = temp.Memo
    return nil
}
