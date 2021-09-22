package http

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// group all routes under /v1/api
	v1 := router.Group("/v1/api")
	{
		//movie
		v1.GET("/movie", s.GetMovies())
		v1.GET("/movie/top-rated", s.GetTopRatedMovie())
		v1.GET("/movie/filter", s.GetFilterMovie())
		v1.GET("/movie/:id", s.GetDetailMovie())

	}

	v1 = router.Group("")
	{
		v1.GET("/", s.rootHandler())

		// Login route
		v1.GET("/login/github/", s.githubLoginHandler())

		// Github callback
		v1.GET("/login/github/callback", s.githubCallbackHandler())

	}

	return router
}
