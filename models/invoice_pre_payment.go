package models

import (
    "encoding/json"
)

// InvoicePrePayment represents a InvoicePrePayment struct.
type InvoicePrePayment struct {
    // The subscription id for the prepayment account
    SubscriptionId       *int   `json:"subscription_id,omitempty"`
    // The amount in cents of the prepayment that was created as a result of this payment.
    AmountInCents        *int64 `json:"amount_in_cents,omitempty"`
    // The total balance of the prepayment account for this subscription including any prior prepayments
    EndingBalanceInCents *int64 `json:"ending_balance_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePrePayment.
// It customizes the JSON marshaling process for InvoicePrePayment objects.
func (i *InvoicePrePayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePrePayment object to a map representation for JSON marshaling.
func (i *InvoicePrePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    if i.SubscriptionId != nil {
        structMap["subscription_id"] = i.SubscriptionId
    }
    if i.AmountInCents != nil {
        structMap["amount_in_cents"] = i.AmountInCents
    }
    if i.EndingBalanceInCents != nil {
        structMap["ending_balance_in_cents"] = i.EndingBalanceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePrePayment.
// It customizes the JSON unmarshaling process for InvoicePrePayment objects.
func (i *InvoicePrePayment) UnmarshalJSON(input []byte) error {
    temp := &struct {
        SubscriptionId       *int   `json:"subscription_id,omitempty"`
        AmountInCents        *int64 `json:"amount_in_cents,omitempty"`
        EndingBalanceInCents *int64 `json:"ending_balance_in_cents,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    i.SubscriptionId = temp.SubscriptionId
    i.AmountInCents = temp.AmountInCents
    i.EndingBalanceInCents = temp.EndingBalanceInCents
    return nil
}
