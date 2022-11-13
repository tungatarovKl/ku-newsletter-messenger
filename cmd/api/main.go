package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	BotToken   string
	DbAddress  string `toml:"dbAddress"`
	DbName     string `toml:"dbName"`
	DbUsername string `toml:"dbUsername"`
	DbPassword string `toml:"dbPassword"`
}

func handleNewsLetter(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	if cfg.DbAddress == "" || cfg.DbName == "" || cfg.DbPassword == "" || cfg.DbUsername == "" {
		log.Fatalf("Отсутсвуют необходимые для работы строки конфига")
	}

	dsn := cfg.DbUsername + ":" + cfg.DbPassword + "@tcp(" + cfg.DbAddress + ")/" + cfg.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	http.HandleFunc("/newsletter", handleNewsLetter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
