package main

import (
	"flag"
	"gorm.io/gorm"
	"log"
	"upgrade/Model/Repository"
	"upgrade/config"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()
	cfg := config.ReadConfig(*configPath)
	database := Repository.NewDatabase(cfg.DbAddress, cfg.DbName, cfg.DbUsername, cfg.DbPassword)
	CreateUsersTable(database.Connection)
}

func CreateUsersTable(db *gorm.DB) {
	//Check if table already exists
	if db.Migrator().HasTable(&Repository.User{}) {
		log.Fatal("Table Users is already exists")
	}

	//Create new table
	err := db.Migrator().CreateTable(&Repository.User{})
	if err != nil {
		log.Fatal("Error creating table 'Users': ", err.Error())
	}
}
