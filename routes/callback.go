package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/stationapi/station-login/db"
	"github.com/stationapi/station-login/session"
	"gorm.io/gorm"
)

type GithubResponse struct {
	AccessToken string `json:"access_token"`
}

type GithubUser struct {
	Id int `json:"id"`
}

func Callback(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
	code := r.URL.Query().Get("code")

	token, err := getCode(code)

	sid := uuid.NewString()

	cookie := http.Cookie{
		Name:   "station",
		Value:  sid,
		Domain: os.Getenv("COOKIE_DOMAIN"),
		Path: "/",
	}

	http.SetCookie(w, &cookie)

	if err != nil {
		return err
	}

	user, err := getUser(token)

	if err != nil {
		return err
	}

	putErr := session.PutSession(sid, token, user.Id)

	if putErr != nil {
		return putErr
	}

	createUser(user.Id, db)

	http.Redirect(w, r, "https://statn.space/dashboard", 302)

	return nil
}

func createUser(githubId int, gormDB gorm.DB) string {
	user := db.User{
		GithubId: githubId,
	}

	exists := db.ValidateUser(githubId, gormDB)

	if exists {
		return "You have been logged in successfully."
	}

	apiKey := db.CreateUser(user, gormDB)

	fmt.Println(apiKey)

	return apiKey
}

func getUser(token string) (GithubUser, error) {
	client := http.Client{}

	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)

	if err != nil {
		return GithubUser{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	res, err := client.Do(req)

	if err != nil {
		return GithubUser{}, err
	}

	var user GithubUser

	parseErr := ProcessBody(res.Body, &user)

	if parseErr != nil {
		return GithubUser{}, parseErr
	}

	return user, nil
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
