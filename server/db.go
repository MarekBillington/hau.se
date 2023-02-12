package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(isDev bool) *gorm.DB {

	// @todo move to .env file
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/London"
	if (isDev) {
		dsn = "postgres://postgres:postgres@127.0.0.1:5432/postgres"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
