package db

import "gorm.io/gorm"

func ValidateUser(githubId int, db gorm.DB) bool {
	var user User

	db.First(&user, githubId)

	return user.GithubId != 0 
}
