package api

import (
	"log"
	"strconv"

	"github.com/LiminalFish/kempt/internal/userdb"
	"github.com/LiminalFish/kempt/pkg/models"
	"github.com/gin-gonic/gin"
)

// /users
func get_users(c *gin.Context) {
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

	userIDs, err := userdb.GetAllUserIDS()

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
func get_users_ID(c *gin.Context) {
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

	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	user := userdb.GetUser(userIDInt)
	user.Token = "SECRET"

	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.JSON(400, err.Error())
	// 	return
	// }
	//
	c.JSON(200, user)
}

// /users/<USER_ID>/name
func get_users_ID_name(c *gin.Context) {
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
func get_users_ID_points(c *gin.Context) {
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
func get_users_ID_type(c *gin.Context) {
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
func get_users_ID_rank(c *gin.Context) {
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
func get_users_ID_team(c *gin.Context) {
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
