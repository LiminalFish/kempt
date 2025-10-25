package main

// import "fmt"
import (
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("This currently does nothing")
	router := gin.Default()

	router.StaticFile("/", "./web/index.html")

	// Define a GET route for the root path "/"
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello from Gin!") // Respond with a string
	// })

	// Another example route
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "pong"}) // Respond with JSON
	// })

	// Run the server on port 8080
	router.Run(":8080")
}
