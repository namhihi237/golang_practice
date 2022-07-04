package utils

import (
	"practice/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
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
	}

	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: expireTime(),
	}

	tokeClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokeClaim.SignedString([]byte(env.JWT_secret))

	return &token, err
}

func expireTime() *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
}
