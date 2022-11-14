package main

import (
	"flag"
	"log"
	"net/http"
	"upgrade/config"
	"upgrade/internal/bot"
	"upgrade/internal/controllers"
	"upgrade/internal/models"
)

type Config struct {
	BotToken   string
	DbAddress  string `toml:"dbAddress"`
	DbName     string `toml:"dbName"`
	DbUsername string `toml:"dbUsername"`
	DbPassword string `toml:"dbPassword"`
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg, cfgErr := config.ReadConfig(*configPath)
	if cfgErr != nil {
		log.Fatal(cfgErr.Error())
	}

	database, dbErr := models.NewDatabase(cfg.DbAddress, cfg.DbName, cfg.DbUsername, cfg.DbPassword)
	if dbErr != nil {
		log.Fatal("Error connecting to database: ", dbErr)
	}

	tgBot := bot.Bot{
		Bot:      bot.InitBot(cfg.BotToken),
		Database: database,
	}

	http.HandleFunc("/newsletter", controllers.NewsLetterPost(tgBot))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
