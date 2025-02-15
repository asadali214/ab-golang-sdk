package models

import (
    "encoding/json"
)

// SaleRepSubscription represents a SaleRepSubscription struct.
type SaleRepSubscription struct {
    Id              *int             `json:"id,omitempty"`
    SiteName        *string          `json:"site_name,omitempty"`
    SubscriptionUrl *string          `json:"subscription_url,omitempty"`
    CustomerName    *string          `json:"customer_name,omitempty"`
    CreatedAt       *string          `json:"created_at,omitempty"`
    Mrr             *string          `json:"mrr,omitempty"`
    Usage           *string          `json:"usage,omitempty"`
    Recurring       *string          `json:"recurring,omitempty"`
    LastPayment     *string          `json:"last_payment,omitempty"`
    ChurnDate       Optional[string] `json:"churn_date"`
}

// MarshalJSON implements the json.Marshaler interface for SaleRepSubscription.
// It customizes the JSON marshaling process for SaleRepSubscription objects.
func (s *SaleRepSubscription) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SaleRepSubscription object to a map representation for JSON marshaling.
func (s *SaleRepSubscription) toMap() map[string]any {
    structMap := make(map[string]any)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.SiteName != nil {
        structMap["site_name"] = s.SiteName
    }
    if s.SubscriptionUrl != nil {
        structMap["subscription_url"] = s.SubscriptionUrl
    }
    if s.CustomerName != nil {
        structMap["customer_name"] = s.CustomerName
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt
    }
    if s.Mrr != nil {
        structMap["mrr"] = s.Mrr
    }
    if s.Usage != nil {
        structMap["usage"] = s.Usage
    }
    if s.Recurring != nil {
        structMap["recurring"] = s.Recurring
    }
    if s.LastPayment != nil {
        structMap["last_payment"] = s.LastPayment
    }
    if s.ChurnDate.IsValueSet() {
        structMap["churn_date"] = s.ChurnDate.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SaleRepSubscription.
// It customizes the JSON unmarshaling process for SaleRepSubscription objects.
func (s *SaleRepSubscription) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id              *int             `json:"id,omitempty"`
        SiteName        *string          `json:"site_name,omitempty"`
        SubscriptionUrl *string          `json:"subscription_url,omitempty"`
        CustomerName    *string          `json:"customer_name,omitempty"`
        CreatedAt       *string          `json:"created_at,omitempty"`
        Mrr             *string          `json:"mrr,omitempty"`
        Usage           *string          `json:"usage,omitempty"`
        Recurring       *string          `json:"recurring,omitempty"`
        LastPayment     *string          `json:"last_payment,omitempty"`
        ChurnDate       Optional[string] `json:"churn_date"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Id = temp.Id
    s.SiteName = temp.SiteName
    s.SubscriptionUrl = temp.SubscriptionUrl
    s.CustomerName = temp.CustomerName
    s.CreatedAt = temp.CreatedAt
    s.Mrr = temp.Mrr
    s.Usage = temp.Usage
    s.Recurring = temp.Recurring
    s.LastPayment = temp.LastPayment
    s.ChurnDate = temp.ChurnDate
    return nil
}
