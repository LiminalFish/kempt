package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// takes user:pass and returns encrypted hash
func login(c *gin.Context) {
	pass := &LoginHeader{}

	err := c.ShouldBindHeader(pass)

	if err != nil {
		c.JSON(400, err.Error())
		log.Println(err.Error())
	}

	c.JSON(200, "my nuts")
}
