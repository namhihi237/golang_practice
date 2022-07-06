package middleware

import (
	"practice/models"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		var err error
		var code = errors.SUCCESS

		token := utils.GetBearerToken(c)
		if token == "" {
			code = errors.UNAUTHORIZED
		} else {
			data, err = utils.ParseToken(token)
			if err != nil {
				code = errors.INVALID_TOKEN
			}

			user := data.(*utils.Claims)

			if isAdmin {
				code = validateAdmin(user)
			} else {
				code = validateUser(user)
			}
		}

		if code != errors.SUCCESS {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  errors.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("user", data)
		c.Next()
	}
}

func validateUser(user *utils.Claims) int {
	me, err := models.GetUserById(user.Id)
	if err != nil {
		return errors.SERVER_ERROR
	}

	if me == nil || user.UserType != "user" {
		return errors.UNAUTHORIZED_ACCESS
	}

	return CheckUser(me)
}

func validateAdmin(admin *utils.Claims) int {
	me, err := models.GetAdminById(admin.Id)
	if err != nil {
		return errors.SERVER_ERROR
	}

	if me == nil || admin.UserType != "admin" {
		return errors.UNAUTHORIZED_ACCESS
	}

	return CheckAdmin(me)
}
