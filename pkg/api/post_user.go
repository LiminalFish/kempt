package api

import (
	"log"
	"strconv"

	"github.com/LiminalFish/kempt/internal/userdb"
	"github.com/LiminalFish/kempt/pkg/models"
	"github.com/gin-gonic/gin"
)

// /users/<USER_ID>/name
func post_users_ID_name(c *gin.Context) {
	permLevel := models.MOD

	type NameChange struct {
		Name string `json:"name"`
	}

	var changeName NameChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetUser := userdb.GetUser(userIDInt)

	if userPerm > int(permLevel) || token != targetUser.Token {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeName); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userdb.AssignUserName(targetUser.UserID, changeName.Name)
}

// /users/<USER_ID>/points
func post_users_ID_points(c *gin.Context) {
	permLevel := models.MOD

	type PointChange struct {
		Newpoint int `json:"newpoint"`
	}

	var changePoint PointChange

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetUser := userdb.GetUser(userIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changePoint); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userdb.AssignUserPoints(targetUser.UserID, changePoint.Newpoint)
}

// /users/<USER_ID>/type
func post_users_ID_type(c *gin.Context) {
	permLevel := models.ADMIN

	type PointPerm struct {
		NewPerm int `json:"newperm"`
	}

	var changePerm PointPerm

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetUser := userdb.GetUser(userIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changePerm); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userdb.AssignUserType(targetUser.UserID, changePerm.NewPerm)
}

// /users/<USER_ID>/rank
func post_users_ID_rank(c *gin.Context) {

	permLevel := models.MOD

	type ChangeRank struct {
		NewRank int `json:"newrank"`
	}

	var changeRank ChangeRank

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetUser := userdb.GetUser(userIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeRank); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userdb.AssignUserRank(targetUser.UserID, changeRank.NewRank)
}

// /users/<USER_ID>/team
func post_users_ID_team(c *gin.Context) {

	permLevel := models.MOD

	type ChangeTeam struct {
		NewTeam int `json:"newteam"`
	}

	var changeTeam ChangeTeam

	token := c.GetHeader("auth")

	userPerm, err := auth(token)

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userIDInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	targetUser := userdb.GetUser(userIDInt)

	if userPerm > int(permLevel) {
		log.Println("Not high enough permissions")
		c.JSON(400, "Not high enough permissions")
		return
	}

	if err := c.BindJSON(&changeTeam); err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	userdb.AssignUserTeamid(targetUser.UserID, changeTeam.NewTeam)
}
