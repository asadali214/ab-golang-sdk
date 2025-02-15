package models

import (
    "encoding/json"
)

// ListMetafieldsResponse represents a ListMetafieldsResponse struct.
type ListMetafieldsResponse struct {
    TotalCount  *int        `json:"total_count,omitempty"`
    CurrentPage *int        `json:"current_page,omitempty"`
    TotalPages  *int        `json:"total_pages,omitempty"`
    PerPage     *int        `json:"per_page,omitempty"`
    Metafields  []Metafield `json:"metafields,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListMetafieldsResponse.
// It customizes the JSON marshaling process for ListMetafieldsResponse objects.
func (l *ListMetafieldsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListMetafieldsResponse object to a map representation for JSON marshaling.
func (l *ListMetafieldsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if l.TotalCount != nil {
        structMap["total_count"] = l.TotalCount
    }
    if l.CurrentPage != nil {
        structMap["current_page"] = l.CurrentPage
    }
    if l.TotalPages != nil {
        structMap["total_pages"] = l.TotalPages
    }
    if l.PerPage != nil {
        structMap["per_page"] = l.PerPage
    }
    if l.Metafields != nil {
        structMap["metafields"] = l.Metafields
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMetafieldsResponse.
// It customizes the JSON unmarshaling process for ListMetafieldsResponse objects.
func (l *ListMetafieldsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        TotalCount  *int        `json:"total_count,omitempty"`
        CurrentPage *int        `json:"current_page,omitempty"`
        TotalPages  *int        `json:"total_pages,omitempty"`
        PerPage     *int        `json:"per_page,omitempty"`
        Metafields  []Metafield `json:"metafields,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.TotalCount = temp.TotalCount
    l.CurrentPage = temp.CurrentPage
    l.TotalPages = temp.TotalPages
    l.PerPage = temp.PerPage
    l.Metafields = temp.Metafields
    return nil
}
