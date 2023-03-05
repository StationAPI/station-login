package routes

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	clientId, err := ReadSecret("./config/gh/id")

	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&scope=read:user&redirect_uri=http://195.154.71.170:3000/login/github/callback",
		clientId,
	)

	http.Redirect(w, r, url, http.StatusFound)

	return nil
}
