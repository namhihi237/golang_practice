package routers

import (
	"practice/middleware"
	"practice/modules/admin"
	"practice/modules/category"
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

	// middleware.JWT(false) is used to check user
	// middleware.JWT(true) is used to check admin
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/profile", middleware.JWT(false), user.GetProfile())
		apiV1.PUT("/profile", middleware.JWT(false), user.UpdateUser())

		//category
		apiV1.POST("/categories", middleware.JWT(true), category.AddCategory())
		apiV1.PUT("/categories/:id", middleware.JWT(true), category.UpdateCategory())
	}

	return router
}
