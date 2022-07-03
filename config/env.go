package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DbUrl  string `json:"db_url"`
	Port   string `json:"port"`
	Go_env string `json:"go_env"`
}

func GetEnv() (Env, error) {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	return Env{
		DbUrl: os.Getenv("DB_URL"),
		Port:  os.Getenv("PORT"),
	}, e
}
