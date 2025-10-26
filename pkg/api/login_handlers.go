package api

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/LiminalFish/kempt/pkg/models"
	"github.com/gin-gonic/gin"
)

func HashString(input string) string {
	// 3. Create a new SHA-256 hash object.
	hasher := sha256.New()

	// 4. Write the input string (as a byte slice) into the hasher.
	// The Write method never returns an error.
	hasher.Write([]byte(input))

	// 5. Calculate the final hash.
	// Sum(nil) appends the current hash to a nil slice and returns it.
	hashBytes := hasher.Sum(nil)

	// 6. Convert the raw bytes of the hash into a human-readable
	// hex string (e.g., "a1b2c3...")
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

// takes user:pass and returns encrypted hash
func login(c *gin.Context) {
	pass := &models.LoginHeader{}

	err := c.ShouldBindHeader(pass)

	if err != nil {
		c.JSON(400, err.Error())
		log.Println(err.Error())
	}

	c.JSON(200, HashString(pass.Pass))
}
