package entity

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Type        EventType `json:"type"`

	EventSessions []EventSession `json:"event_sessions,omitempty"`
}

type EventInDetail struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Type        EventType `json:"type"`

	EventSessions []EventSessionInDetail `json:"event_sessions"`
}
