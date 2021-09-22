package models

type MoviesProductionCompany struct {
	ID                  uint `gorm:"primaryKey"`
	MovieId             uint
	Movie               Movie `gorm:"foreignKey:MovieId;references:ID"`
	ProductionCompanyId uint
	ProductionCompany   ProductionCompany `gorm:"foreignKey:ProductionCompanyId;references:ID"`
}
