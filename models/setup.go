package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DB *gorm.DB


func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&Character{}, &Weapon{}, &GachaHistory{})

	if err != nil {
		panic(err)
	}
	DB = database
	return database
}
