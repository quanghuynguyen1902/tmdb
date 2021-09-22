package models

type MoviesGenre struct {
	ID      uint `gorm:"primaryKey"`
	MovieId uint
	Movie   Movie `gorm:"foreignKey:MovieId;references:ID"`
	GenreId uint
	Genre   Genre `gorm:"foreignKey:GenreId;references:ID"`
}
