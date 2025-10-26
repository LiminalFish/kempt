package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve HTML pages
	router.StaticFile("/", "./web/index.html")
	router.StaticFile("/completed", "./web/completed.html")
	router.StaticFile("/scoreboard", "./web/scoreboard.html")
	router.StaticFile("/settings", "./web/settings.html")

	// FIXED: Serve task.html for /task/:id paths (e.g. /task/1)
	router.GET("/task/:id", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./web/task.html")
	})

	router.Static("/assets", "./web/assets/")

	// Hardcoded example tasks
	tasks := []map[string]string{
		{"id": "1", "title": "Task 1", "description": "Do something important", "status": "Incomplete", "due_date": "2025-11-01"},
		{"id": "2", "title": "Task 2", "description": "Do another thing", "status": "Complete", "due_date": "2025-11-05"},
	}

	// Get all tasks
	router.GET("/api/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})

	// Get one task by ID
	router.GET("/api/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, t := range tasks {
			if t["id"] == id {
				c.JSON(http.StatusOK, t)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	// Edit task (mock)
	router.POST("/api/tasks/edit/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updated map[string]string
		if err := c.BindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		for i, t := range tasks {
			if t["id"] == id {
				for k, v := range updated {
					tasks[i][k] = v
				}
				c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	})

	router.Run(":8080")
}
