package main

import (
	"dining-hall/src/food"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	foodList := food.GetFoodList();
	foodMap := food.GetFoodMap();

	println(food.GetFoodList())
	println(foodList)
	println(foodMap)


	// default path
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Dining-Hall is up!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping/kitchen", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":4005")
}
