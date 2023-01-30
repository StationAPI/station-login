package main

import (
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		log.Fatal("The DSN environment variable was not found.")

		return
	}
}
