package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartKemptAPI() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/login", login)

	router.GET("/tasks", tasks)

	router.GET("tasks/add", tasks_add)

	router.GET("tasks/edit/:id", tasks_edit_ID)

	router.GET("tasks/delete/:id", tasks_delete_ID)

	router.GET("/users", users)

	router.GET("/users/:id", users_ID)

	router.GET("/users/:id/name", users_ID_name)

	router.GET("/users/:id/points", users_ID_points)

	router.GET("/users/:id/type", users_ID_type)

	router.GET("/users/:id/rank", users_ID_rank)

	router.GET("/users/:id/team", users_ID_team)

	router.Run(":8081")
}
