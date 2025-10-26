package api

import (
	"log"

	"github.com/LiminalFish/kempt/internal/userdb"
)

func auth(token string) (int, error) {

	user, err := userdb.GetUserFromToken(token)

	if err != nil {
		log.Println(err.Error())
		return -1, err
	}

	return user.Type, nil
}
