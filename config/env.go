package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DbUrl         string `json:"db_url"`
	Port          string `json:"port"`
	Go_env        string `json:"go_env"`
	JWT_secret    string `json:"jwt_secret"`
	Email         string `json:"email"`
	EmailPass     string `json:"email_pass"`
	AdminPassword string `json:"admin_password"`
	AdminUsername string `json:"admin_username"`
	AdminEmail    string `json:"admin_email"`
}

func GetEnv() (Env, error) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	return Env{
		DbUrl:         os.Getenv("DB_URL"),
		Port:          os.Getenv("PORT"),
		Go_env:        os.Getenv("GO_ENV"),
		JWT_secret:    os.Getenv("JWT_SECRET"),
		Email:         os.Getenv("EMAIL"),
		EmailPass:     os.Getenv("EMAIL_PASSWORD"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
		AdminUsername: os.Getenv("ADMIN_USERNAME"),
		AdminEmail:    os.Getenv("ADMIN_EMAIL"),
	}, e
}
