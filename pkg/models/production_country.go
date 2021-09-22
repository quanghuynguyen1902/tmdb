package models

type ProductionCountry struct {
	ID        uint `gorm:"primaryKey"`
	MovieId   uint
	Movie     Movie `gorm:"foreignKey:MovieId;references:ID"`
	CountryId uint
	Country   Country `gorm:"foreignKey:CountryId;references:ID"`
}
