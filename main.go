package main

import (
	"log"

	"github.com/stationapi/station-login/db"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)

func main() {
  db, err := db.Connect() 

  if err != nil {
    log.Fatal(err)
  }

  r := chi.NewRouter()

  r.Use(middleware.Logger)
  r.Use(middleware.RealIP)
}
