package api

import (
	"encoding/json"
	"github.com/tmdb/pkg/db"
	"github.com/tmdb/pkg/models"
)

type MovieService interface {
	Movies(pagination *models.Pagination) ([]models.Movie, error)
	DetailMovie(id uint) (models.Movie, error)
	TopRatedMovie(pagination *models.Pagination) ([]models.Movie, error)
	FilterMovie(pagination *models.Pagination, filter *models.MovieFilter) ([]models.Movie, error)
}

type MovieRepository interface {
	GetMovies(pagination *models.Pagination) ([]models.Movie, error)
	GetDetailMovie(id uint) (models.Movie, error)
	GetTopRatedMovie(pagination *models.Pagination) ([]models.Movie, error)
	GetFilterMovie(pagination *models.Pagination, filter *models.MovieFilter) ([]models.Movie, error)
}

type movieService struct {
	storage MovieRepository
	cache   *db.RedisCache
}

func NewMovieService(movieRepo MovieRepository, cache *db.RedisCache) MovieService {
	return &movieService{
		storage: movieRepo,
		cache:   cache,
	}
}

func (m *movieService) Movies(pagination *models.Pagination) ([]models.Movie, error) {
	movies, err := m.storage.GetMovies(pagination)
	return movies, err
}

func (m *movieService) DetailMovie(id uint) (models.Movie, error) {
	movieCache, err := m.cache.Get(string(id))
	if err != nil {
		movie, e := m.storage.GetDetailMovie(id)
		m.cache.Set(string(id), movie)
		return movie, e
	}
	data := models.Movie{}
	json.Unmarshal([]byte(movieCache), &data)

	return data, nil
}

func (m *movieService) TopRatedMovie(pagination *models.Pagination) ([]models.Movie, error) {
	movies, err := m.storage.GetTopRatedMovie(pagination)
	return movies, err
}

func (m *movieService) FilterMovie(pagination *models.Pagination, filter *models.MovieFilter) ([]models.Movie, error) {
	movies, err := m.storage.GetFilterMovie(pagination, filter)
	return movies, err
}
