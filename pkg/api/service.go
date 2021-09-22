package api

import (
	"github.com/tmdb/pkg/db"
	"github.com/tmdb/pkg/repository"
)

type Service struct {
	storage repository.Storage
	cache   *db.RedisCache
}

func NewService(storage repository.Storage, cache *db.RedisCache) *Service {
	return &Service{
		storage: storage,
		cache:   cache,
	}
}

func (s *Service) CreateMovieService() MovieService {
	return NewMovieService(s.storage, s.cache)
}
