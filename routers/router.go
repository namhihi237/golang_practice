package routers

import (
	"practice/middleware"
	"practice/modules/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// authenticate
	router.POST("/auth/register", user.SignUp())
	router.POST("/auth/login", user.SignIn())

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.GET("/me", user.GetMe())
	}

	return router
}
