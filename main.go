package main

import (
	"practice/models"
	"practice/routers"
	"practice/seed"
)

func main() {
	models.SetUp()
	seed.Run(models.GetDb())

	initRouter := routers.InitRouter()

	initRouter.Run(":8000")
}
