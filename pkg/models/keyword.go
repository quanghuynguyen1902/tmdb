package models

type Keyword struct {
	ID      uint   `gorm:"primaryKey"`
	Keyword string `json:"keyword"`
}
