package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	connectString := "root:12345@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local";
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	return db;
}