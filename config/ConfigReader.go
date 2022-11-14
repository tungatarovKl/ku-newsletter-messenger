// File for config reading
package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	BotToken   string `toml:"botToken"`
	DbAddress  string `toml:"dbAddress"`
	DbName     string `toml:"dbName"`
	DbUsername string `toml:"dbUsername"`
	DbPassword string `toml:"dbPassword"`
}

func ReadConfig(configPath string) (*Config, error) {
	var cfg Config
	var emptyFields error = nil
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		log.Fatal("Error decoding config: ", configPath, err.Error())
	}
	if cfg.DbAddress == "" ||
		cfg.DbName == "" ||
		cfg.DbPassword == "" ||
		cfg.DbUsername == "" ||
		cfg.BotToken == "" {
		emptyFields = errors.New("Needed config strings are absent")
	}
	return &cfg, emptyFields
}
