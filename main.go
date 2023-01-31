package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stationapi/station-login/db"
	"github.com/stationapi/station-login/routes"
)

func main() {
	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Get("/auth/github", func(w http.ResponseWriter, r *http.Request) {
		routes.Login(w, r, db)
	})

	r.Get("/auth/github/callback", func(w http.ResponseWriter, r *http.Request) {
		routes.Callback(w, r, db)
	})
}
