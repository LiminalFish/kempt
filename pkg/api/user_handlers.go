package api

import (
	"github.com/LiminalFish/kempt/internal/userdb"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func users(c *gin.Context) {

}

func users_ID(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	username, err := userdb.GetUserName(userIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"name":   username,
		"userID": userID,
	})
}

func users_teams(c *gin.Context) {

}

func users_teams_ID(c *gin.Context) {

}
