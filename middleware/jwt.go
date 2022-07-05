package middleware

import (
	"fmt"
	"practice/models"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		var err error
		var code = errors.SUCCESS

		token := getBearerToken(c)
		if token == "" {
			code = errors.UNAUTHORIZED
		} else {
			data, err = utils.ParseToken(token)
			if err != nil {
				code = errors.INVALID_TOKEN
			}
		}

		user := data.(*utils.Claims)
		me, err := models.GetUserById(user.Id)
		if err != nil {
			code = errors.SERVER_ERROR
		}
		fmt.Println(me)
		code = CheckUser(me)
		fmt.Println(code)

		if code != errors.SUCCESS {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  errors.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func getBearerToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(bearerToken) > 7 && bearerToken[0:7] == "Bearer " {
		return bearerToken[7:]
	}
	return ""
}
