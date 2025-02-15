/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// InvoiceEventPayment represents a InvoiceEventPayment struct.
// A nested data structure detailing the method of payment
type InvoiceEventPayment struct {
    Type                *string          `json:"type,omitempty"`
    MaskedAccountNumber *string          `json:"masked_account_number,omitempty"`
    MaskedRoutingNumber *string          `json:"masked_routing_number,omitempty"`
    CardBrand           *string          `json:"card_brand,omitempty"`
    CardExpiration      *string          `json:"card_expiration,omitempty"`
    LastFour            Optional[string] `json:"last_four"`
    MaskedCardNumber    *string          `json:"masked_card_number,omitempty"`
    Details             Optional[string] `json:"details"`
    Kind                *string          `json:"kind,omitempty"`
    Memo                Optional[string] `json:"memo"`
    Email               *string          `json:"email,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventPayment.
// It customizes the JSON marshaling process for InvoiceEventPayment objects.
func (i *InvoiceEventPayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventPayment object to a map representation for JSON marshaling.
func (i *InvoiceEventPayment) toMap() map[string]any {
    structMap := make(map[string]any)
    if i.Type != nil {
        structMap["type"] = *i.Type
    } else {
        structMap["type"] = "Invoice Event Payment"
    }
    if i.MaskedAccountNumber != nil {
        structMap["masked_account_number"] = i.MaskedAccountNumber
    }
    if i.MaskedRoutingNumber != nil {
        structMap["masked_routing_number"] = i.MaskedRoutingNumber
    }
    if i.CardBrand != nil {
        structMap["card_brand"] = i.CardBrand
    }
    if i.CardExpiration != nil {
        structMap["card_expiration"] = i.CardExpiration
    }
    if i.LastFour.IsValueSet() {
        structMap["last_four"] = i.LastFour.Value()
    }
    if i.MaskedCardNumber != nil {
        structMap["masked_card_number"] = i.MaskedCardNumber
    }
    if i.Details.IsValueSet() {
        structMap["details"] = i.Details.Value()
    }
    if i.Kind != nil {
        structMap["kind"] = i.Kind
    }
    if i.Memo.IsValueSet() {
        structMap["memo"] = i.Memo.Value()
    }
    if i.Email != nil {
        structMap["email"] = i.Email
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventPayment.
// It customizes the JSON unmarshaling process for InvoiceEventPayment objects.
func (i *InvoiceEventPayment) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Type                *string          `json:"type,omitempty"`
        MaskedAccountNumber *string          `json:"masked_account_number,omitempty"`
        MaskedRoutingNumber *string          `json:"masked_routing_number,omitempty"`
        CardBrand           *string          `json:"card_brand,omitempty"`
        CardExpiration      *string          `json:"card_expiration,omitempty"`
        LastFour            Optional[string] `json:"last_four"`
        MaskedCardNumber    *string          `json:"masked_card_number,omitempty"`
        Details             Optional[string] `json:"details"`
        Kind                *string          `json:"kind,omitempty"`
        Memo                Optional[string] `json:"memo"`
        Email               *string          `json:"email,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    i.Type = temp.Type
    i.MaskedAccountNumber = temp.MaskedAccountNumber
    i.MaskedRoutingNumber = temp.MaskedRoutingNumber
    i.CardBrand = temp.CardBrand
    i.CardExpiration = temp.CardExpiration
    i.LastFour = temp.LastFour
    i.MaskedCardNumber = temp.MaskedCardNumber
    i.Details = temp.Details
    i.Kind = temp.Kind
    i.Memo = temp.Memo
    i.Email = temp.Email
    return nil
}
