package models

type Collection struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
