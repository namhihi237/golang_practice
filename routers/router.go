package routers

import (
	"practice/middleware"
	"practice/modules/admin"
	"practice/modules/user"

	_ "practice/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// authentication user
	router.POST("/auth/register", user.SignUp())
	router.POST("/auth/login", user.SignIn())
	router.GET("/auth/active", user.ActiveAccount())

	// login admin
	router.POST("/admin/login", admin.SignIn())

	apiV1 := router.Group("/api/v1")
	apiV1.Use(middleware.JWT(false))
	{
		apiV1.GET("/profile", user.GetProfile())
		apiV1.PUT("/profile", user.UpdateUser())
	}

	return router
}
