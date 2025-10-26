package api

import (
	"log"
	"strconv"

	"github.com/LiminalFish/kempt/internal/taskdb"
	"github.com/LiminalFish/kempt/pkg/models"
	"github.com/gin-gonic/gin"
)

func post_tasks_ID_title(c *gin.Context) {
	permLevel := models.USER

	type TitleChange struct {
		Title string `json:"title"`
	}

	var changeTitle TitleChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetTask := taskdb.GetTask(taskIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeTitle); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskdb.AssignTaskTitle(targetTask.TaskID, changeTitle.Title)
}

func post_tasks_ID_description(c *gin.Context) {
	permLevel := models.USER

	type DescriptionChange struct {
		Description string `json:"description"`
	}

	var changeDescription DescriptionChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetTask := taskdb.GetTask(taskIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeDescription); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskdb.AssignTaskTitle(targetTask.TaskID, changeDescription.Description)
}

func post_tasks_ID_duedate(c *gin.Context) {
	permLevel := models.USER

	type DueDateChange struct {
		DueDate int `json:"duedate"`
	}

	var changeDueDate DueDateChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetTask := taskdb.GetTask(taskIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeDueDate); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskdb.AssignTaskDuedate(targetTask.TaskID, changeDueDate.DueDate)
}

func post_tasks_ID_timeframe(c *gin.Context) {
	permLevel := models.USER

	type TimeframeChange struct {
		Timeframe string `json:"timeframe"`
	}

	var changeTimeframe TimeframeChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetTask := taskdb.GetTask(taskIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeTimeframe); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskdb.AssignTaskTimeframe(targetTask.TaskID, changeTimeframe.Timeframe)
}

func post_tasks_ID_points(c *gin.Context) {
	permLevel := models.USER

	type PointsChange struct {
		Points int `json:"points"`
	}

	var changePoints PointsChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetTask := taskdb.GetTask(taskIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changePoints); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskdb.AssignTaskPoints(targetTask.TaskID, changePoints.Points)
}

func post_tasks_ID_triggers(c *gin.Context) {
	permLevel := models.USER

	type TriggersChange struct {
		Triggers int `json:"description"`
	}

	var changeTriggers TriggersChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetTask := taskdb.GetTask(taskIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeTriggers); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskdb.AssignTaskTriggers(targetTask.TaskID, changeTriggers.Triggers)
}

func post_tasks_ID_hidden(c *gin.Context) {
	permLevel := models.ADMIN

	type HiddenStatusChange struct {
		HiddenStatus bool `json:"hidden"`
	}

	var changeHiddenStatus HiddenStatusChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetTask := taskdb.GetTask(taskIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeHiddenStatus); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	taskdb.AssignTaskHiddenStatus(targetTask.TaskID, changeHiddenStatus.HiddenStatus)
}
