package models

type Language struct {
	ID      uint   `gorm:"primaryKey"`
	LangKey string `json:"lang_key"`
	Name    string `json:"name"`
}
