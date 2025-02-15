package models

import (
    "encoding/json"
)

// ReplayWebhooksResponse represents a ReplayWebhooksResponse struct.
type ReplayWebhooksResponse struct {
    Status *string `json:"status,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ReplayWebhooksResponse.
// It customizes the JSON marshaling process for ReplayWebhooksResponse objects.
func (r *ReplayWebhooksResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReplayWebhooksResponse object to a map representation for JSON marshaling.
func (r *ReplayWebhooksResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReplayWebhooksResponse.
// It customizes the JSON unmarshaling process for ReplayWebhooksResponse objects.
func (r *ReplayWebhooksResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Status *string `json:"status,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.Status = temp.Status
    return nil
}
