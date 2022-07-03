package main

import (
	"practice/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.SetUp()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run(":8000") // listen and serve on  :8080
}
