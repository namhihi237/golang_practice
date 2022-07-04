package middleware

import (
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

func getBearerToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(bearerToken) > 7 && bearerToken[0:7] == "Bearer " {
		return bearerToken[7:]
	}
	return ""
}
