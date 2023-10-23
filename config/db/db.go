package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func connectWithSQLiteDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("config/db/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func (d Database) Init() *gorm.DB {
	d.db = connectWithSQLiteDatabase()
	return d.db
}
