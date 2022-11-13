package config

import (
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

func ReadConfig(configPath string) *Config {
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		log.Fatal("Error decoding config: ", configPath, err.Error())
	}
	return &cfg
}
