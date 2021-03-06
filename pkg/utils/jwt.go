package utils

import (
	"practice/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

// ref: https://pkg.go.dev/github.com/golang-jwt/jwt/v4@v4.4.2?utm_source=gopls#NewWithClaims
func GenerateToken(data interface{}) (*string, error) {
	var err error
	var token string
	var env config.Env

	env, err = config.GetEnv()
	if err != nil {
		return nil, err
	}

	claims := Claims{
		Id:       data.(map[string]interface{})["id"].(int64),
		UserName: data.(map[string]interface{})["user_name"].(string),
		Email:    data.(map[string]interface{})["email"].(string),
		UserType: data.(map[string]interface{})["user_type"].(string),
	}

	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: expireTime(),
	}

	tokeClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokeClaim.SignedString([]byte(env.JWT_secret))

	return &token, err
}

func ParseToken(token string) (*Claims, error) {
	var err error
	var env config.Env

	env, err = config.GetEnv()
	if err != nil {
		return nil, err
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.JWT_secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func expireTime() *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
}

func GetBearerToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(bearerToken) > 7 && bearerToken[0:7] == "Bearer " {
		return bearerToken[7:]
	}
	return ""
}
