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

type UserRegistration struct {
	UserName string `json:"user_name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email=required"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var appG = app.Gin{C: ctx}
		var userRegistration UserRegistration
		if err := ctx.ShouldBind(&userRegistration); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}

		validate := validator.New()
		err := validate.Struct(userRegistration)
		if err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}
		userExist, err := models.CheckUserByUsernameOrEmail(userRegistration.UserName, userRegistration.Email)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		if userExist {
			appG.Response(http.StatusBadRequest, errors.USER_ALREADY_EXIST, nil)
			return
		}

		hashPassword, err := utils.HashPassword(userRegistration.Password)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.HASH_PASSWORD_ERROR, nil)
			return
		}

		err = models.CreateUser(userRegistration.UserName, hashPassword, userRegistration.Email)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, err)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, nil)
	}
}
