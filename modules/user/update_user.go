package user

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserUpdate struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

// @Summary Update profile
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param full_name body string false "full name"
// @Param phone body string false "phone number"
// @Param birthday body string false "birthday"
// @Param gender body string false "gender"
// @Param email body string false "email"
// @Param address body string false "address"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /profile [put]
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		userLogin := c.MustGet("user").(*utils.Claims)

		var userUpdate UserUpdate
		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}

		user, err := models.GetUserById(userLogin.Id)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		if userUpdate.FullName != "" {
			user.FullName = userUpdate.FullName
		}

		if userUpdate.Phone != "" {
			user.Phone = userUpdate.Phone
		}

		if userUpdate.Birthday != "" {
			birthday, err := utils.ConvertStringToTime(userUpdate.Birthday)
			if err != nil {
				appG.Response(http.StatusInternalServerError, errors.INVALID_PARAMS, nil)
				return
			}
			user.Birthday = birthday
		}

		if userUpdate.Address != "" {
			user.Address = userUpdate.Address
		}

		if userUpdate.Email != "" {
			userExist, err := models.GetUserByEmail(userUpdate.Email)
			if err != nil {
				appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
				return
			}
			if userExist != nil {
				appG.Response(http.StatusInternalServerError, errors.ERROR_EXIST_EMAIL, nil)
				return
			}
			user.Email = userUpdate.Email
		}

		if userUpdate.Address != "" {
			user.Address = userUpdate.Address
		}

		if userUpdate.Gender != "" {
			user.Gender = userUpdate.Gender
		}

		err = models.UpdateUser(user.Id, user)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, nil)

	}
}
