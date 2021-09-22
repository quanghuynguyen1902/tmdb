package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tmdb/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

// add this below APIStatus method
func (s *Server) GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination := utils.GeneratePaginationFromRequest(c)
		c.Header("Content-Type", "application/json")

		movies, err := s.service.CreateMovieService().Movies(&pagination)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   movies,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetDetailMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		movie, err := s.service.CreateMovieService().DetailMovie(uint(id))

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   movie,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetTopRatedMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination := utils.GeneratePaginationFromRequest(c)
		c.Header("Content-Type", "application/json")

		movies, err := s.service.CreateMovieService().TopRatedMovie(&pagination)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   movies,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetFilterMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination := utils.GeneratePaginationFromRequest(c)
		filter := utils.GenerateFilterFromRequest(c)
		c.Header("Content-Type", "application/json")

		movies, err := s.service.CreateMovieService().FilterMovie(&pagination, &filter)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]interface{}{
			"status": "success",
			"data":   movies,
		}

		c.JSON(http.StatusOK, response)
	}
}
