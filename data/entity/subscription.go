package entity

import "github.com/google/uuid"

type Subscriptions struct {
	ID                   uuid.UUID          `gorm:"type:uuid;primary_key" json:"id"`
	UserID               uint               `json:"user_id"`
	EventSessionID       uint               `json:"event_session_id"`
	LocalEventCalendarID string             `json:"local_event_calendar_id"`
	Status               ConfirmationStatus `json:"status"`
	OnNotifyNewEvent     bool               `json:"on_notify_new_event"`

	EventSession EventSession `json:"event_session"`
}
