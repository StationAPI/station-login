package db

import (
	"fmt"
	"crypto/sha256"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	GithubId int
	ApiKey   string
  Bumps int
}

func CreateUser(user User, db gorm.DB) string {
	apiKey := fmt.Sprintf("station_%s", uuid.NewString())

	hash := sha256.Sum256([]byte(apiKey))

	user.ApiKey =	string(hash[:]) 
  user.Bumps = 0

	db.Create(user)

	return apiKey
}
