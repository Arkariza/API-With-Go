package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase()  {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_item"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&User{}, &Item{})
	DB = database
}