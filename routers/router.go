package routers

import (
	"practice/middleware"
	"practice/modules/user"

	_ "practice/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// authenticate
	router.POST("/auth/register", user.SignUp())
	router.POST("/auth/login", user.SignIn())
	router.GET("/auth/active", user.ActiveAccount())

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.GET("/profile", user.GetProfile())
		apiV1.PUT("/profile", user.UpdateUser())
	}

	return router
}
