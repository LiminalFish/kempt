package api

import "github.com/gin-gonic/gin"

func startServer() {
	router := gin.Default()

	router.Run()
}
