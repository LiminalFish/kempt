package frontend

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Static("/assets", "./web/assets")

	router.StaticFile("/", "./web/index.html")

	router.Run("8080")
}
