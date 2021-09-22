package models

type Country struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `json:"code"`
	Name string `json:"name"`
}
