package main

import (
	"flag"
	"log"
	"upgrade/cmd/bot/Bot"
	"upgrade/config"
	"upgrade/internal/models"
)

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

	tgBot := Bot.Bot{
		Bot:      Bot.InitBot(cfg.BotToken),
		Database: database,
	}

	tgBot.Bot.Handle("/start", tgBot.StartHandler)
	tgBot.Bot.Start()
}
