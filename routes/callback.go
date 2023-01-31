package routes

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Callback(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	code := r.URL.Query().Get("code")

	client := http.Client{}

	clientId, err := ReadSecret("./config/gh/id")

	if err != nil {
		return err
	}

	clientSecret, err := ReadSecret("./config/gh/secret")

	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		clientId,
		clientSecret,
		code,
	)

	req, err := http.NewRequest(
		"POST",
		url,
		nil,
	)

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
