package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	data "../book"
	"log"
)

func InitDatabase()  {

	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect database")
		panic(err.Error())
	}

	db.AutoMigrate(&data.Books{})
	log.Println("splite database connect")

}