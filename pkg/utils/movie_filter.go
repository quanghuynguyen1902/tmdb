package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/tmdb/pkg/models"
)

//GeneratePaginationFromRequest ..
func GenerateFilterFromRequest(c *gin.Context) models.MovieFilter {
	// Initializing default
	genre := ""
	language := ""
	company := ""
	//	var mode string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "genre":
			genre = queryValue
			break
		case "language":
			language = queryValue
			break
		case "company":
			company = queryValue
			break

		}
	}
	return models.MovieFilter{
		Genre:    genre,
		Language: language,
		Company:  company,
	}

}
