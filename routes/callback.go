package routes

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type GithubResponse struct {
	AccessToken string `json:"access_token"`
}

type GithubUser struct {
	Id string `json:"id"`
}

func Callback(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	code := r.URL.Query().Get("code")

	token, err := getCode(code)

	cookie := http.Cookie{
		Name:  "station",
		Value: token,
	}

	r.AddCookie(&cookie)

	if err != nil {
		return err
	}

	return nil
}

func getUser(token string) (GithubUser, error) {
	return GithubUser{}, nil
}

func getCode(code string) (string, error) {
	client := http.Client{}

	clientId, err := ReadSecret("./config/gh/id")

	if err != nil {
		return "", err
	}

	clientSecret, err := ReadSecret("./config/gh/secret")

	if err != nil {
		return "", err
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
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	var githubResponse GithubResponse

	parseErr := ProcessBody(res.Body, &githubResponse)

	if parseErr != nil {
		return "", parseErr
	}

	return githubResponse.AccessToken, nil
}
