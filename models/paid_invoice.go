package models

import (
    "encoding/json"
)

// PaidInvoice represents a PaidInvoice struct.
type PaidInvoice struct {
    // The uid of the paid invoice
    InvoiceId  *string        `json:"invoice_id,omitempty"`
    // The current status of the invoice. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more.
    Status     *InvoiceStatus `json:"status,omitempty"`
    // The remaining due amount on the invoice
    DueAmount  *string        `json:"due_amount,omitempty"`
    // The total amount paid on this invoice (including any prior payments)
    PaidAmount *string        `json:"paid_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PaidInvoice.
// It customizes the JSON marshaling process for PaidInvoice objects.
func (p *PaidInvoice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaidInvoice object to a map representation for JSON marshaling.
func (p *PaidInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.InvoiceId != nil {
        structMap["invoice_id"] = p.InvoiceId
    }
    if p.Status != nil {
        structMap["status"] = p.Status
    }
    if p.DueAmount != nil {
        structMap["due_amount"] = p.DueAmount
    }
    if p.PaidAmount != nil {
        structMap["paid_amount"] = p.PaidAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaidInvoice.
// It customizes the JSON unmarshaling process for PaidInvoice objects.
func (p *PaidInvoice) UnmarshalJSON(input []byte) error {
    temp := &struct {
        InvoiceId  *string        `json:"invoice_id,omitempty"`
        Status     *InvoiceStatus `json:"status,omitempty"`
        DueAmount  *string        `json:"due_amount,omitempty"`
        PaidAmount *string        `json:"paid_amount,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.InvoiceId = temp.InvoiceId
    p.Status = temp.Status
    p.DueAmount = temp.DueAmount
    p.PaidAmount = temp.PaidAmount
    return nil
}
