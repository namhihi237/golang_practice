package routers

import (
	"practice/middleware"
	"practice/modules/admin"
	"practice/modules/category"
	"practice/modules/product"
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
	apiV1.Use(middleware.JWT(false))
	{
		apiV1.GET("/profile", user.GetProfile())
		apiV1.PUT("/profile", user.UpdateUser())

		//category

	}

	adminV1 := router.Group("/admin/api/v1")
	adminV1.Use(middleware.JWT(true))
	{
		// category
		adminV1.POST("/categories", category.AddCategory())
		adminV1.PUT("/categories/:id", category.UpdateCategory())
		adminV1.GET("/categories", category.AdminGetListCategories())
		adminV1.GET("/categories/:id", category.AdminGetCategory())
		adminV1.DELETE("/categories/:id", category.DeleteCategory())

		// product
		adminV1.POST("/products", product.AdminAddProduct())
	}

	return router
}
