package models

type Genre struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
