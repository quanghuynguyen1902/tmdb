package models

type Movie struct {
	ID                  uint    `gorm:"primaryKey"`
	Title               string  `json:"title"`
	ReleaseDate         string  `json:"release_date"`
	Budget              int     `json:"budget"`
	Revenue             int     `json:"revenue"`
	Popularity          float32 `json:"popularity"`
	Runtime             uint    `json:"runtime"`
	Rating              float32 `json:"rating"`
	OriginalLanguage    uint
	Language            Language `gorm:"foreignKey:OriginalLanguage;references:ID"`
	BelongsToCollection uint
	Collection          Collection `gorm:"foreignKey:BelongsToCollection;references:ID"`
	Overview            string     `json:"overview"`
}

type MovieFilter struct {
	Genre    string
	Company  string
	Language string
}
