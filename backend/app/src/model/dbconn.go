package model

import (
  	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
)


func Init_DB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	dbconn = db
}