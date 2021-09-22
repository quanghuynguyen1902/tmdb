package models

type ProductionCompany struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
