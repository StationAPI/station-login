package main

import (
	"fmt"
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

	r.Get("/login/github", func(w http.ResponseWriter, r *http.Request) {
		err := routes.Login(w, r, db)

		if err != nil {
			fmt.Println(err)

			http.Error(w, "There was an error authenticating you. Please try again later.", http.StatusInternalServerError)
		}
	})

	r.Get("/login/github/callback", func(w http.ResponseWriter, r *http.Request) {
		err := routes.Callback(w, r, db)

		if err != nil {
			fmt.Println(err)

			http.Error(w, "There was an error authenticating you. Please try again later.", http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":3000", r)
}
