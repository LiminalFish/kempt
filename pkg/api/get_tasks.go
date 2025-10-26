package api

import (
	"log"
	"strconv"

	"github.com/LiminalFish/kempt/internal/taskdb"
	"github.com/LiminalFish/kempt/pkg/models"
	"github.com/gin-gonic/gin"
)

func tasks(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskIDs, err := taskdb.GetAllTaskIDS()

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskIDs": taskIDs,
	})
}

func get_tasks_ID(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	task := taskdb.GetTask(taskIDInt)

	c.JSON(200, task)
}

func get_tasks_ID_title(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	title, err := taskdb.GetTaskTitle(taskIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskID": taskIDInt,
		"title":  title,
	})
}

func get_tasks_ID_description(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	description, err := taskdb.GetTaskDescription(taskIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskID":      taskIDInt,
		"description": description,
	})
}

func get_tasks_ID_duedate(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	duedate, err := taskdb.GetTaskDuedate(taskIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskID":  taskIDInt,
		"duedate": duedate,
	})
}

func get_tasks_ID_timeframe(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	timeframe, err := taskdb.GetTaskTimeframe(taskIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskID":    taskIDInt,
		"timeframe": timeframe,
	})
}

func get_tasks_ID_points(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	points, err := taskdb.GetTaskPoints(taskIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskID": taskIDInt,
		"points": points,
	})
}

func get_tasks_ID_triggers(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	triggers, err := taskdb.GetTaskTriggers(taskIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskID":   taskIDInt,
		"triggers": triggers,
	})
}

func get_tasks_ID_hidden(c *gin.Context) {
	permLevel := models.USER

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	taskID := c.Param("id")

	taskIDInt, err := strconv.Atoi(taskID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	hidden, err := taskdb.GetTaskHiddenStatus(taskIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"taskID": taskIDInt,
		"hidden": hidden,
	})
}
