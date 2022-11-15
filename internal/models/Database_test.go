package models

import (
	"flag"
	"testing"
	"upgrade/config"
)

var configPath = flag.String("config", "", "Path to config file")

func TestNewDatabase(t *testing.T) {
	cfg, _ := config.ReadConfig(*configPath)
	database, dbErr := NewDatabase(cfg.DbAddress, cfg.DbName, cfg.DbUsername, cfg.DbPassword)
	if dbErr != nil || database.Connection.Error != nil {
		t.Errorf("Error: %v\nDB Error: %v\n", dbErr, database.Connection.Error)
	}
}
