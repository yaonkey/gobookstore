package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	db, err = gorm.Open("postgres", "user=gobookstore password=test dbname=gobookstoredb sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
