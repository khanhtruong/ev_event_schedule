package store

import "event_schedule/data/entity"

func (s *storeImpl) GetEvents() ([]entity.Event, error) {
	events := []entity.Event{}
	return events, s.DB.Find(&events).Error
}

func (s *storeImpl) GetEvent(id string) (*entity.Event, error) {
	event := entity.Event{}
	return &event, s.DB.Preload("EventSessions").First(&event, "id=?", id).Error
}
