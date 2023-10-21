package main

import (
	"golang/config/db"
	"golang/product"
	"golang/server"
	"gorm.io/gorm"
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

	defer dbManager.Close()

	server.Init()
}
