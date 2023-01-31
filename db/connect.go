package db

import (
	"errors"
  "os"

	"gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func Connect() (gorm.DB, error) {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		return gorm.DB{}, errors.New("The DSN environment variable was not found.")
	}

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    DisableForeignKeyConstraintWhenMigrating: true,
  })

  if err != nil {
    return gorm.DB{}, err
  }

  return *db, nil
}
