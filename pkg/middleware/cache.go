package middleware

import "github.com/gin-gonic/gin"

func verifyCache() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		c.Next()
	}
}
