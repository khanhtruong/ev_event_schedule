package entity

type EventDetail struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	EventTypeID uint   `json:"event_type_id"`

	EventType EventType `json:"event_type"`
}
