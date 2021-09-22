package models

type SpokenLanguage struct {
	ID         uint `gorm:"primaryKey"`
	MovieId    uint
	Movie      Movie `gorm:"foreignKey:MovieId;references:ID"`
	LanguageId uint
	Language   Language `gorm:"foreignKey:LanguageId;references:ID"`
}
