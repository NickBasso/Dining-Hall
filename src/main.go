package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// default path
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Dining-Hall is up!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":4005")
}
