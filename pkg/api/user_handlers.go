package api

import (
	"github.com/LiminalFish/kempt/internal/userdb"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// /users
func users(c *gin.Context) {
	userIDs, err := userdb.GetAllUsers()

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"usersIDs": userIDs,
	})
}

// /users/<USER_ID>
func users_ID(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	user := userdb.GetUser(userIDInt)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(400, err.Error())
	// 	return
	// }
	//
	c.JSON(200, user)
}

// /users/<USER_ID>/name
func users_ID_name(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	name, err := userdb.GetUserName(userIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"userID": userIDInt,
		"name":   name,
	})
}

// /users/<USER_ID>/points
func users_ID_points(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	points, err := userdb.GetUserPoints(userIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"userID": userIDInt,
		"points": points,
	})
}

// /users/<USER_ID>/type
func users_ID_type(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userType, err := userdb.GetUserType(userIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"userID": userIDInt,
		"type":   userType,
	})
}

// /users/<USER_ID>/rank
func users_ID_rank(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	rank, err := userdb.GetUserRank(userIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"userID": userIDInt,
		"rank":   rank,
	})
}

// /users/<USER_ID>/team
func users_ID_team(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	teamID, err := userdb.GetUserTeamid(userIDInt)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"userID": userIDInt,
		"teamID": teamID,
	})
}
