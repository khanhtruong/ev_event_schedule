package entity

type ConfirmationStatus string

const (
	Pending   ConfirmationStatus = "pending"
	Confirmed ConfirmationStatus = "confirmed"
	Declined  ConfirmationStatus = "declined"
)
