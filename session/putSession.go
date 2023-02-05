package session

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type SessionObject struct {
	Id string `json:"cookie"`
	Session map[string]interface{} `json:"session"`
}

func GetSessionCache() string {
	ip := os.Getenv("STATION_SESSION_CACHE_SERVICE_HOST")
	port := os.Getenv("STATION_SESSION_CACHE_SERVICE_PORT")

	url := fmt.Sprintf("http://%s:%s/put", ip, port)

	return url 
}

func PutSession(sid string, accessToken string) error {
  session := SessionObject{
    Id: sid, 
    Session: map[string]interface{}{
      "access_token": accessToken,
    },
  }

  stringifed, err := json.Marshal(session)

  if err != nil {
    return err
  }

  client := http.Client{}

  req, err := http.NewRequest(
    "POST",
    GetSessionCache(),
    bytes.NewReader(stringifed),
  )

  if err != nil {
    return err
  }

  res, err := client.Do(req)

  if err != nil {
    return err
  }

  if res.StatusCode > 300 {
    return errors.New("there was an error caching the session.") 
  }

	return nil
}
