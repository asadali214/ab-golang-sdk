package models

import (
    "encoding/json"
)

// RevokedInvitation represents a RevokedInvitation struct.
type RevokedInvitation struct {
    LastSentAt     *string `json:"last_sent_at,omitempty"`
    LastAcceptedAt *string `json:"last_accepted_at,omitempty"`
    UninvitedCount *int    `json:"uninvited_count,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RevokedInvitation.
// It customizes the JSON marshaling process for RevokedInvitation objects.
func (r *RevokedInvitation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RevokedInvitation object to a map representation for JSON marshaling.
func (r *RevokedInvitation) toMap() map[string]any {
    structMap := make(map[string]any)
    if r.LastSentAt != nil {
        structMap["last_sent_at"] = r.LastSentAt
    }
    if r.LastAcceptedAt != nil {
        structMap["last_accepted_at"] = r.LastAcceptedAt
    }
    if r.UninvitedCount != nil {
        structMap["uninvited_count"] = r.UninvitedCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RevokedInvitation.
// It customizes the JSON unmarshaling process for RevokedInvitation objects.
func (r *RevokedInvitation) UnmarshalJSON(input []byte) error {
    temp := &struct {
        LastSentAt     *string `json:"last_sent_at,omitempty"`
        LastAcceptedAt *string `json:"last_accepted_at,omitempty"`
        UninvitedCount *int    `json:"uninvited_count,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.LastSentAt = temp.LastSentAt
    r.LastAcceptedAt = temp.LastAcceptedAt
    r.UninvitedCount = temp.UninvitedCount
    return nil
}
