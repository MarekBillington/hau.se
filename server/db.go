package main

import (
	"log"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	// dbURL := "postgres://postgres:postgres@127.0.0.1:5432/postgres"
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/London"

	fmt.Println(dsn);

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&House{})

	return db
}
