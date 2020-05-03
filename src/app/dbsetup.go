package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func setUpDb() (*gorm.DB)  {

	//db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=go_app password=root sslmode=disable")
	db, err := gorm.Open("postgres", "host=my-postgres user=postgres dbname=go_app password=root sslmode=disable")

	if (err != nil) {
		panic( err.Error())
	}

	//defer db.Close()
	db.AutoMigrate(&Product{})
	return db
}
