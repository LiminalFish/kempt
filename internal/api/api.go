package api

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"jit": "latooogen"})
	})

	router.Run()
}
