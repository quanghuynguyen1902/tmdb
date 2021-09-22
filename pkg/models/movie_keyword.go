package models

type MoviesKeyword struct {
	ID        uint `gorm:"primaryKey"`
	MovieId   uint
	Movie     Movie `gorm:"foreignKey:MovieId;references:ID"`
	KeywordId uint
	Keyword   Keyword `gorm:"foreignKey:KeywordId;references:ID"`
}
