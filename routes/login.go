package routes

import (
	"fmt"
	"net/http"
	"os"

	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	clientId, err := ReadSecret("./config/gh/id")

	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&scope=read:user&redirect_uri=http://%s:%s/login/github/callback",
		clientId,
		os.Getenv("CALLBACK_URL_HOST"),
		os.Getenv("CALLBACK_URL_PORT"),
	)

	http.Redirect(w, r, url, http.StatusFound)

	return nil
}
