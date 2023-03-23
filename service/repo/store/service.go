package store

import (
	"event_schedule/data/entity"
)

type Store interface {
	// Event
	GetEvents(filters []string, from string, to string) ([]entity.Event, error)
	GetEvent(id string) (*entity.Event, error)

	// Subscriptions
	SubscriptionAmount(eventSessionID string) (*int64, error)
	Subscribe(subscription *entity.Subscriptions) (*entity.Subscriptions, error)
	GetSubscription(id string) (*entity.Subscriptions, error)
	UpdateSubscription(id string, eventSessionID string) (*entity.Subscriptions, error)
	DeleteSubscription(id string) error
	UpdateSubscriptionCalendarID(id string, calendarID string) (*entity.Subscriptions, error)
	GetSubscriptions(userID uint) ([]entity.Subscriptions, error)
}
