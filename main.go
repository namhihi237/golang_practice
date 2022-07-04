package main

import (
	"practice/models"

	"practice/routers"
)

func main() {
	models.SetUp()

	initRouter := routers.InitRouter()

	initRouter.Run(":8000")

}
