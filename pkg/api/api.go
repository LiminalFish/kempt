package api

import (
	"github.com/gin-gonic/gin"
)

func start_kempt_API() {
	router := gin.Default()

	router.GET("/login", login)

	router.GET("/tasks", tasks)

	router.GET("tasks/add", tasks_add)

	router.GET("tasks/edit/:id", tasks_edit_ID)

	router.GET("tasks/delete/:id", tasks_delete_ID)

	router.GET("/users", users)

	router.GET("/users/:id", users_ID)

	router.GET("/users/teams", users_teams)

	router.GET("/users/teams/:id", users_teams_ID)

	router.Run("localhost:311911")
}
