package db

import "gorm.io/gorm"

func ValidateUser(githubId string, db gorm.DB) bool {
	var user User

	db.First(&user, githubId)

	return user.GithubId != ""
}
