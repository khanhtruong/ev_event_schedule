package entity

type EventSession struct {
	ID         uint  `gorm:"primaryKey" json:"id"`
	EventID    uint  `json:"event_id"`
	SlotAmount int   `json:"slot_amount"`
	StartTime  int64 `json:"start_time"`
	EndTime    int64 `json:"end_time"`

	Event Event `json:"event"`
}

type EventSessionInDetail struct {
	EventSession

	Available *int64 `json:"available"`
}
