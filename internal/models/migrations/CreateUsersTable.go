package main

import (
	"flag"
	"gorm.io/gorm"
	"log"
	"upgrade/config"
	"upgrade/internal/models"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()
	cfg := config.ReadConfig(*configPath)
	database := models.NewDatabase(cfg.DbAddress, cfg.DbName, cfg.DbUsername, cfg.DbPassword)
	CreateUsersTable(database.Connection)
}

func CreateUsersTable(db *gorm.DB) {
	//Check if table already exists
	if db.Migrator().HasTable(&models.User{}) {
		log.Println("Table already exists")
	}

	//Create new table
	err := db.Migrator().CreateTable(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Table %v created!")
}
