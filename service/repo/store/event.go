package store

import (
	"event_schedule/data/entity"
	"fmt"
	"strings"
)

func (s *storeImpl) GetEvents(filters []string, from string, to string) ([]entity.Event, error) {
	events := []entity.Event{}

	var f = make([]string, len(filters))
	for i := range filters {
		f[i] = fmt.Sprintf(`'%s'`, filters[i])
	}

	rawQuery := fmt.Sprintf(`SELECT "events"."id","events"."title","events"."description","events"."location","events"."type" FROM "events" LEFT JOIN (SELECT distinct on ("event_id") * FROM "event_sessions") as es ON events.id = es.event_id WHERE type IN (%s)`, strings.Join(f, ","))

	switch {
	case from != "" && to != "":
		rawQuery += fmt.Sprintf(" AND es.start_time >= %s AND es.start_time <= %s", from, to)
	case from != "":
		rawQuery += fmt.Sprintf(" AND es.start_time >= %s", from)
	case to != "":
		rawQuery += fmt.Sprintf(" AND es.end_time <= %s", to)
	}

	fmt.Println(rawQuery)

	return events, s.DB.Raw(rawQuery).Preload("EventSessions").Find(&events).Error
}

func (s *storeImpl) GetEvent(id string) (*entity.Event, error) {
	event := entity.Event{}
	return &event, s.DB.Preload("EventSessions").First(&event, "id=?", id).Error
}
