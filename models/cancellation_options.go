package models

import (
    "encoding/json"
)

// CancellationOptions represents a CancellationOptions struct.
type CancellationOptions struct {
    // For your internal use. An indication as to why the subscription is being canceled.
    CancellationMessage *string `json:"cancellation_message,omitempty"`
    // The reason code associated with the cancellation. See the list of reason codes associated with your site.
    ReasonCode          *string `json:"reason_code,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CancellationOptions.
// It customizes the JSON marshaling process for CancellationOptions objects.
func (c *CancellationOptions) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CancellationOptions object to a map representation for JSON marshaling.
func (c *CancellationOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.CancellationMessage != nil {
        structMap["cancellation_message"] = c.CancellationMessage
    }
    if c.ReasonCode != nil {
        structMap["reason_code"] = c.ReasonCode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancellationOptions.
// It customizes the JSON unmarshaling process for CancellationOptions objects.
func (c *CancellationOptions) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CancellationMessage *string `json:"cancellation_message,omitempty"`
        ReasonCode          *string `json:"reason_code,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.CancellationMessage = temp.CancellationMessage
    c.ReasonCode = temp.ReasonCode
    return nil
}
