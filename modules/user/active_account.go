package user

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ActiveAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		token := c.Query("token")

		if token == "" {
			appG.Response(http.StatusBadRequest, errors.INVALID_TOKEN, nil)
			return
		}

		data, err := utils.ParseToken(token)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.INVALID_TOKEN, nil)
			return
		}

		user, err := models.GetUserByEmail(data.Email)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		models.ActiveAccount(user.Id)

		appG.Response(http.StatusOK, errors.SUCCESS, nil)
	}
}
