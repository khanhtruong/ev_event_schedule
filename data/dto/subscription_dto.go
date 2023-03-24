package dto

type SubscribeDto struct {
	Data SubscribeDataDto `json:"data"`
}

type SubscribeDataDto struct {
	EventSessionID   uint `json:"event_session_id"`
	OnNotifyNewEvent bool `json:"on_notify_new_event"`
}

type SubscribeCalendarIDDto struct {
	Data SubscribeCalendarIDDataDto `json:"data"`
}

type SubscribeCalendarIDDataDto struct {
	LocalEventCalendarID string `json:"local_event_calendar_id"`
}
