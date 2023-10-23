package main

import (
	"database/sql"
	"golang/config/db"
	"golang/domain/product"
	"golang/server"
	"gorm.io/gorm"
	"log"
)

func migrateAll(db *gorm.DB) {
	err := db.AutoMigrate(product.Product{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func main() {
	dbInit := db.Database{}.Init()
	migrateAll(dbInit)

	dbManager, _ := dbInit.DB()

	defer func(dbManager *sql.DB) {
		err := dbManager.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(dbManager)

	server.Init()
}
