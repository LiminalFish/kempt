package main

// import "fmt"
import (
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("This currently does nothing")
	router := gin.Default()

	router.StaticFile("/", "./web/index.html")
	router.StaticFile("/scoreboard", "./web/scoreboard.html")
	router.StaticFile("/settings", "./web/settings.html")
	router.Static("/assets", "./web/assets/")

	// get database tasks
	router.GET("/api/tasks", func(c *gin.Context) {
    tasks := []map[string]string{
        {"title": "Task 1", "description": "Do something important"},
        {"title": "Task 2", "description": "Do another thing"},
    }
    c.JSON(200, tasks)
})

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
