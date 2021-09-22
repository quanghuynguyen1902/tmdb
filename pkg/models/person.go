package models

type Person struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
