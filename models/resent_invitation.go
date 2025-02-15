package models

import (
    "encoding/json"
)

// ResentInvitation represents a ResentInvitation struct.
type ResentInvitation struct {
    LastSentAt         *string `json:"last_sent_at,omitempty"`
    LastAcceptedAt     *string `json:"last_accepted_at,omitempty"`
    SendInviteLinkText *string `json:"send_invite_link_text,omitempty"`
    UninvitedCount     *int    `json:"uninvited_count,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ResentInvitation.
// It customizes the JSON marshaling process for ResentInvitation objects.
func (r *ResentInvitation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResentInvitation object to a map representation for JSON marshaling.
func (r *ResentInvitation) toMap() map[string]any {
    structMap := make(map[string]any)
    if r.LastSentAt != nil {
        structMap["last_sent_at"] = r.LastSentAt
    }
    if r.LastAcceptedAt != nil {
        structMap["last_accepted_at"] = r.LastAcceptedAt
    }
    if r.SendInviteLinkText != nil {
        structMap["send_invite_link_text"] = r.SendInviteLinkText
    }
    if r.UninvitedCount != nil {
        structMap["uninvited_count"] = r.UninvitedCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResentInvitation.
// It customizes the JSON unmarshaling process for ResentInvitation objects.
func (r *ResentInvitation) UnmarshalJSON(input []byte) error {
    temp := &struct {
        LastSentAt         *string `json:"last_sent_at,omitempty"`
        LastAcceptedAt     *string `json:"last_accepted_at,omitempty"`
        SendInviteLinkText *string `json:"send_invite_link_text,omitempty"`
        UninvitedCount     *int    `json:"uninvited_count,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.LastSentAt = temp.LastSentAt
    r.LastAcceptedAt = temp.LastAcceptedAt
    r.SendInviteLinkText = temp.SendInviteLinkText
    r.UninvitedCount = temp.UninvitedCount
    return nil
}
