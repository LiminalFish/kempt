package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartKemptAPI() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/login", login)

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

	router.GET("/tasks", tasks)

	router.GET("tasks/:id", get_tasks_ID)

	router.GET("tasks/:id/title", get_tasks_ID_title)

	router.GET("tasks/:id/description", get_tasks_ID_description)

	router.GET("tasks/:id/duedate", get_tasks_ID_duedate)

	router.GET("tasks/:id/timeframe", get_tasks_ID_timeframe)

	router.GET("tasks/:id/points", get_tasks_ID_points)

	router.GET("tasks/:id/triggers", get_tasks_ID_triggers)

	router.GET("tasks/:id/hidden", get_tasks_ID_hidden)

	router.POST("tasks/:id/title", post_tasks_ID_title)

	router.POST("tasks/:id/description", post_tasks_ID_description)

	router.POST("tasks/:id/duedate", post_tasks_ID_duedate)

	router.POST("tasks/:id/timeframe", post_tasks_ID_timeframe)

	router.POST("tasks/:id/points", post_tasks_ID_points)

	router.POST("tasks/:id/triggers", post_tasks_ID_triggers)

	router.POST("tasks/:id/hidden", post_tasks_ID_hidden)

	router.Run(":8081")
}
