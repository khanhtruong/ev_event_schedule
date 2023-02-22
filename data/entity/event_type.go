package entity

type EventType string

const (
	InPerson  EventType = "in_person"
	Virtual   EventType = "virtual"
	TestDrive EventType = "test_drive"
)
