package store

import (
	"event_schedule/data/entity"
)

func (s *storeImpl) SubscriptionAmount(eventSessionID string) (*int64, error) {
	var result int64 = 0
	return &result, s.DB.Table("subscriptions").Count(&result).Where("event_session_id=?", eventSessionID).Error
}

func (s *storeImpl) Subscribe(subscription *entity.Subscriptions) (*entity.Subscriptions, error) {
	return subscription, s.DB.Create(subscription).Error
}

func (s *storeImpl) GetSubscription(id string) (*entity.Subscriptions, error) {
	subscription := entity.Subscriptions{}
	return &subscription, s.DB.Preload("EventSession.Event").First(&subscription, "id=?", id).Error
}

func (s *storeImpl) UpdateSubscription(id string, eventSessionID string) (*entity.Subscriptions, error) {
	subscription := entity.Subscriptions{}
	return &subscription, s.DB.Model(&subscription).Where("id=?", id).Update("event_session_id", eventSessionID).Error
}

func (s *storeImpl) DeleteSubscription(id string) error {
	return s.DB.Model(&entity.Subscriptions{}).Where("id=?", id).Update("status", entity.Declined).Error
}

func (s *storeImpl) UpdateSubscriptionCalendarID(id string, calendarID string) (*entity.Subscriptions, error) {
	subscription := entity.Subscriptions{}
	return &subscription, s.DB.Model(&subscription).Where("id=?", id).Update("local_event_calendar_id", calendarID).Error
}
