package models

type Director struct {
	ID         uint `gorm:"primaryKey"`
	MovieId    uint
	Movie      Movie `gorm:"foreignKey:MovieId;references:ID"`
	DirectorId uint
	Person     Person `gorm:"foreignKey:DirectorId;references:ID"`
}
