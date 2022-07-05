package user

import (
	"net/http"
	"practice/middleware"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	UserName string `json:"user_name" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserResponse struct {
	Id       int64      `json:"id"`
	UserName string     `json:"user_name"`
	FullName string     `json:"fullName"`
	Email    string     `json:"email"`
	Phone    string     `json:"phone"`
	Address  string     `json:"address"`
	Gender   string     `json:"gender"`
	Birthday *time.Time `json:"birthday"`
}

// @Summary Login
// @Produce  json
// @Param user_name body string true "user name"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/login [post]
func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appG = app.Gin{C: c}
		var validate = validator.New()

		var userLogin UserLogin
		if err := c.ShouldBind(&userLogin); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}

		err := validate.Struct(userLogin)
		if err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}

		user, err := models.GetUserByUsername(userLogin.UserName)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.INVALID_USER_NAME_PASSWORD, nil)
			return
		}

		match := utils.ComparePassword(user.Password, userLogin.Password)
		if !match {
			appG.Response(http.StatusBadRequest, errors.INVALID_USER_NAME_PASSWORD, nil)
			return
		}

		code := middleware.CheckUser(user)
		if code != errors.SUCCESS {
			appG.Response(http.StatusUnauthorized, code, nil)
			return
		}

		data := map[string]interface{}{
			"id":        user.Id,
			"user_name": user.UserName,
			"email":     user.Email,
			"user_type": user.UserType.Name,
		}

		token, err := utils.GenerateToken(data)
		if err != nil || token == nil {
			appG.Response(http.StatusInternalServerError, errors.GEN_TOKEN_ERROR, nil)
			return
		}

		userResponse := &UserResponse{
			Id:       user.Id,
			UserName: user.UserName,
			FullName: user.FullName,
			Email:    user.Email,
			Phone:    user.Phone,
			Address:  user.Address,
			Gender:   user.Gender,
			Birthday: user.Birthday,
		}

		appG.Response(http.StatusOK, errors.SUCCESS, map[string]interface{}{
			"token": token,
			"user":  userResponse,
		})

	}
}
