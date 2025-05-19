package routes

import (
	"gin-auth-boilerplate/internal/auth/router"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	router.RegisterAuthRoutes(r)

	return r
}
