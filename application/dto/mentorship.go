package dto

import (
	"study-event-go/types"
	"time"
)

// Mentorship ...
type Mentorship struct {
	ID        types.MentorshipID `json:"id,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	DeletedAt *time.Time         `json:"deleted_at,omitempty"`
	Name      string             `json:"name,omitempty"`
}
