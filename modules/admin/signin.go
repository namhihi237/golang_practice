package admin

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type AdminResponse struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

// @Summary Admin Login
// @Produce  json
// @Param user_name body string true "user name"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/login [post]
func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		var adminLogin AdminLogin
		if err := c.ShouldBindJSON(&adminLogin); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}

		admin, err := models.GetAdminByUserName(adminLogin.UserName)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.INVALID_USER_NAME_PASSWORD, nil)
			return
		}

		match := utils.ComparePassword(admin.Password, adminLogin.Password)
		if !match {
			appG.Response(http.StatusInternalServerError, errors.INVALID_USER_NAME_PASSWORD, nil)
			return
		}

		data := map[string]interface{}{
			"id":        admin.Id,
			"user_name": admin.UserName,
			"email":     admin.Email,
		}

		token, err := utils.GenerateToken(data)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.GEN_TOKEN_ERROR, nil)
			return
		}

		adminResponse := AdminResponse{
			Id:       admin.Id,
			UserName: admin.UserName,
			Email:    admin.Email,
		}

		appG.Response(http.StatusOK, errors.SUCCESS, map[string]interface{}{
			"token": token,
			"admin": adminResponse,
		})
	}
}
