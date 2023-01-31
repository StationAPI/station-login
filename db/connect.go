package db

import (
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (gorm.DB, error) {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		return gorm.DB{}, errors.New("the dsn environment variable was not found")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return gorm.DB{}, err
	}

	return *db, nil
}
