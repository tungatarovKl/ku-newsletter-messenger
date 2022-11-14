package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
	DbAddress  string
	DbName     string
	DbUsername string
	DbPassword string
}

func NewDatabase(dbAddress, dbName, dbUsername, dbPassword string) (*Database, error) {
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbAddress + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &Database{
		Connection: db,
		DbAddress:  dbAddress,
		DbName:     dbName,
		DbUsername: dbUsername,
		DbPassword: dbPassword,
	}, err
}
