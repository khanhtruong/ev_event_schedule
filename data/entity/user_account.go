package entity

type UserAccount struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	IsVIP bool   `json:"is_vip"`
}
