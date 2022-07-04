package user

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	UserName string `json:"user_name" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserResponse struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}

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

		if !isActive(user) {
			appG.Response(http.StatusBadRequest, errors.INACTIVE_USER, nil)
			return
		}

		if isBlocked(user) {
			appG.Response(http.StatusBadRequest, errors.USER_BLOCKED, nil)
			return
		}

		data := map[string]interface{}{
			"id":        user.Id,
			"user_name": user.UserName,
			"email":     user.Email,
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
		}

		appG.Response(http.StatusOK, errors.SUCCESS, map[string]interface{}{
			"token": token,
			"user":  userResponse,
		})

	}
}

func isActive(user *models.User) bool {
	return user.IsActive
}

func isBlocked(user *models.User) bool {
	return user.IsBlocked
}
