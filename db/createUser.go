package db

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	GithubId int
	ApiKey   string
}

func CreateUser(user User, db gorm.DB) string {
	apiKey := fmt.Sprintf("station_%s", uuid.NewString())

	user.ApiKey = apiKey

	db.Create(user)

	return apiKey
}
