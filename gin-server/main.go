package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.GET("/getTime", *gin.Context) {
		c.JSON(200, gin.H{
			"serverTime": time.Now().UY,
	})
}
