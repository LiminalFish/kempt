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

	router.GET("/users", get_users)

	router.GET("/users/:id", get_users_ID)

	router.GET("/users/:id/name", get_users_ID_name)

	router.GET("/users/:id/points", get_users_ID_points)

	router.GET("/users/:id/type", get_users_ID_type)

	router.GET("/users/:id/rank", get_users_ID_rank)

	router.GET("/users/:id/team", get_users_ID_team)

	router.POST("/users/:id/name", post_users_ID_name)

	router.POST("/users/:id/points", post_users_ID_points)

	router.POST("/users/:id/type", post_users_ID_type)

	router.POST("/users/:id/rank", post_users_ID_rank)

	router.POST("/users/:id/team", post_users_ID_team)

	router.Run(":8081")
}
