package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Up!",
		})
	})

	app.GET("/getTime", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now().UTC(),
		})
	})

	app.Run(":8050")
}
