package repo

import (
	"event_schedule/config"
	"event_schedule/data/entity"
	"event_schedule/service/repo/store"
)

type repoImpl struct {
	Config *config.Config
	Store  store.Store
}

func NewRepo(cfg *config.Config) Repo {
	s := store.NewStore(cfg)

	return &repoImpl{
		Config: cfg,
		Store:  s,
	}
}

func (r *repoImpl) GetEvents() ([]entity.Event, error) {
	return r.Store.GetEvents()
}

func (r *repoImpl) GetEvent(id string) (*entity.Event, error) {
	return r.Store.GetEvent(id)
}

func (r *repoImpl) SubscriptionAmount(eventSessionID string) (*int64, error) {
	return r.Store.SubscriptionAmount(eventSessionID)
}

func (r *repoImpl) Subscribe(subscription *entity.Subscriptions) (*entity.Subscriptions, error) {
	return r.Store.Subscribe(subscription)
}

func (r *repoImpl) GetSubscription(id string) (*entity.Subscriptions, error) {
	return r.Store.GetSubscription(id)
}

func (r *repoImpl) UpdateSubscription(id string, eventSessionID string) (*entity.Subscriptions, error) {
	return r.Store.UpdateSubscription(id, eventSessionID)
}

func (r *repoImpl) DeleteSubscription(id string) error {
	return r.Store.DeleteSubscription(id)
}

func (r *repoImpl) UpdateSubscriptionCalendarID(id string, calendarID string) (*entity.Subscriptions, error) {
	return r.Store.UpdateSubscriptionCalendarID(id, calendarID)
}
