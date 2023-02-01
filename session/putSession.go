package session

import (
	"fmt"
	"os"
)

type SessionObject struct {
	Cookie string `json:"cookie"`
	Session map[string]interface{} `json:"session"`
}

func GetSessionCache() string {
	ip := os.Getenv("FC_SESSION_CACHE_SERVICE_HOST")
	port := os.Getenv("FC_SESSION_CACHE_SERVICE_PORT")

	url := fmt.Sprintf("http://%s:%s", ip, port)

	return url 
}

func PutSession(accessToken string) error {
	return nil
}
