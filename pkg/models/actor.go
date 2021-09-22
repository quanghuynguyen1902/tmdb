package models

type Actor struct {
	ID       uint `gorm:"primaryKey""`
	PersonId uint
	Person   Person `gorm:"foreignKey:PersonId;references:ID"`
	MovieId  uint
	Movie    Movie `gorm:"foreignKey:MovieId;references:ID"`
	OrderId  uint  `json:"order_id"`
}
