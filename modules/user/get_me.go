package user

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int64  `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	UserName string `gorm:"not null; size:255; unique;" json:"user_name"`
	Email    string `gorm:"not null; size:255; unique;" json:"email"`
	FullName string `gorm:"not null; size:255;" json:"full_name"`
	Phone    string `gorm:"size:20;" json:"phone"`
	Address  string `gorm:"size:255;" json:"address"`
	Gender   string `gorm:"size:10;" json:"gender"`
}

func GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		userLogin := c.MustGet("user").(*utils.Claims)

		user, err := models.GetUserById(userLogin.Id)

		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		userResponse := &User{
			Id:       user.Id,
			UserName: user.UserName,
			Email:    user.Email,
			FullName: user.FullName,
			Phone:    user.Phone,
			Address:  user.Address,
			Gender:   user.Gender,
		}

		appG.Response(http.StatusOK, errors.SUCCESS, *userResponse)
	}
}
