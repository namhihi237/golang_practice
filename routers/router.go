package routers

import (
	"practice/modules/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// authenticate
	router.POST("/api/v1/auth/register", user.SignUp())
	router.POST("/api/v1/auth/login", user.SignIn())

	return router
}
