package models

import (
    "encoding/json"
)

// PaginatedMetadata represents a PaginatedMetadata struct.
type PaginatedMetadata struct {
    TotalCount  *int       `json:"total_count,omitempty"`
    CurrentPage *int       `json:"current_page,omitempty"`
    TotalPages  *int       `json:"total_pages,omitempty"`
    PerPage     *int       `json:"per_page,omitempty"`
    Metadata    []Metadata `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PaginatedMetadata.
// It customizes the JSON marshaling process for PaginatedMetadata objects.
func (p *PaginatedMetadata) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaginatedMetadata object to a map representation for JSON marshaling.
func (p *PaginatedMetadata) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.TotalCount != nil {
        structMap["total_count"] = p.TotalCount
    }
    if p.CurrentPage != nil {
        structMap["current_page"] = p.CurrentPage
    }
    if p.TotalPages != nil {
        structMap["total_pages"] = p.TotalPages
    }
    if p.PerPage != nil {
        structMap["per_page"] = p.PerPage
    }
    if p.Metadata != nil {
        structMap["metadata"] = p.Metadata
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaginatedMetadata.
// It customizes the JSON unmarshaling process for PaginatedMetadata objects.
func (p *PaginatedMetadata) UnmarshalJSON(input []byte) error {
    temp := &struct {
        TotalCount  *int       `json:"total_count,omitempty"`
        CurrentPage *int       `json:"current_page,omitempty"`
        TotalPages  *int       `json:"total_pages,omitempty"`
        PerPage     *int       `json:"per_page,omitempty"`
        Metadata    []Metadata `json:"metadata,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.TotalCount = temp.TotalCount
    p.CurrentPage = temp.CurrentPage
    p.TotalPages = temp.TotalPages
    p.PerPage = temp.PerPage
    p.Metadata = temp.Metadata
    return nil
}
