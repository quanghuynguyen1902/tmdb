package repository

import (
	"github.com/tmdb/pkg/models"
)

func (s *storage) GetMovies(pagination *models.Pagination) ([]models.Movie, error) {
	var movies []models.Movie
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := s.db.Limit(pagination.Limit).Offset(offset).Order("id " + pagination.Sort)
	err := queryBuilder.Preload("Collection").Preload("Language").Find(&movies).Error

	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *storage) GetDetailMovie(id uint) (models.Movie, error) {
	var movie models.Movie

	err := s.db.Preload("Collection").Preload("Language").First(&movie, "movies.id", id).Error

	return movie, err
}

func (s *storage) GetTopRatedMovie(pagination *models.Pagination) ([]models.Movie, error) {
	var movies []models.Movie
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := s.db.Limit(pagination.Limit).Offset(offset).Order("rating " + pagination.Sort + " nulls last")
	err := queryBuilder.Preload("Collection").Preload("Language").Find(&movies).Error

	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *storage) GetFilterMovie(pagination *models.Pagination, filter *models.MovieFilter) ([]models.Movie, error) {
	var movies []models.Movie
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := s.db.Table("movies").Select("movies.* ").Joins("JOIN movies_genres on movies.id = movies_genres.movie_id").Joins("JOIN genres on movies_genres.genre_id = genres.id").Where("genres.name = ?", filter.Genre)
	if filter.Company != "" {
		queryBuilder = queryBuilder.Joins("JOIN movies_production_companies on movies.id = movies_production_companies.movie_id").Joins("JOIN production_companies on movies_production_companies.production_company_id = production_companies.id").Where("production_companies.name = ?", filter.Company)
	}
	queryBuilder = queryBuilder.Limit(pagination.Limit).Offset(offset).Order("id " + pagination.Sort)
	err := queryBuilder.Preload("Collection").Preload("Language").Find(&movies).Error

	if err != nil {
		return nil, err
	}
	return movies, nil
}
