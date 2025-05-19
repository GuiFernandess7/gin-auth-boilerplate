package routes

import (
	"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"

	//_ "gin-auth-boilerplate/docs"
	"gin-auth-boilerplate/internal/auth/router"
)

func InitializeRoutes() *gin.Engine {
	r := gin.Default()
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	router.RegisterAuthRoutes(r)

	return r
}
