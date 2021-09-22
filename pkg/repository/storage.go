package repository

import (
	"github.com/elastic/go-elasticsearch"
	"github.com/tmdb/pkg/models"
	"gorm.io/gorm"
)

type Storage interface {
	// movie storage
	GetMovies(pagination *models.Pagination) ([]models.Movie, error)
	GetTopRatedMovie(pagination *models.Pagination) ([]models.Movie, error)
	GetDetailMovie(id uint) (models.Movie, error)
	GetFilterMovie(pagination *models.Pagination, filter *models.MovieFilter) ([]models.Movie, error)
}

type storage struct {
	db *gorm.DB
	ESClient *elasticsearch.Client
}

func NewStorage(db *gorm.DB) Storage {
	return &storage{
		db: db,
	}
}
