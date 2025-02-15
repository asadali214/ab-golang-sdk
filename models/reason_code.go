package models

import (
    "encoding/json"
    "log"
    "time"
)

// ReasonCode represents a ReasonCode struct.
type ReasonCode struct {
    Id          *int       `json:"id,omitempty"`
    SiteId      *int       `json:"site_id,omitempty"`
    Code        *string    `json:"code,omitempty"`
    Description *string    `json:"description,omitempty"`
    Position    *int       `json:"position,omitempty"`
    CreatedAt   *time.Time `json:"created_at,omitempty"`
    UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ReasonCode.
// It customizes the JSON marshaling process for ReasonCode objects.
func (r *ReasonCode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReasonCode object to a map representation for JSON marshaling.
func (r *ReasonCode) toMap() map[string]any {
    structMap := make(map[string]any)
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.Code != nil {
        structMap["code"] = r.Code
    }
    if r.Description != nil {
        structMap["description"] = r.Description
    }
    if r.Position != nil {
        structMap["position"] = r.Position
    }
    if r.CreatedAt != nil {
        structMap["created_at"] = r.CreatedAt.Format(time.RFC3339)
    }
    if r.UpdatedAt != nil {
        structMap["updated_at"] = r.UpdatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReasonCode.
// It customizes the JSON unmarshaling process for ReasonCode objects.
func (r *ReasonCode) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id          *int    `json:"id,omitempty"`
        SiteId      *int    `json:"site_id,omitempty"`
        Code        *string `json:"code,omitempty"`
        Description *string `json:"description,omitempty"`
        Position    *int    `json:"position,omitempty"`
        CreatedAt   *string `json:"created_at,omitempty"`
        UpdatedAt   *string `json:"updated_at,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.Id = temp.Id
    r.SiteId = temp.SiteId
    r.Code = temp.Code
    r.Description = temp.Description
    r.Position = temp.Position
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        r.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        r.UpdatedAt = &UpdatedAtVal
    }
    return nil
}
