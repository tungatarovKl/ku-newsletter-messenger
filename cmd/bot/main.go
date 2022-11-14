package main

import (
	"flag"
	"log"
	"upgrade/config"
	"upgrade/internal/bot"
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

	tgBot := bot.Bot{
		Bot:      bot.InitBot(cfg.BotToken),
		Database: database,
	}

	tgBot.Bot.Handle("/start", tgBot.StartHandler)
	tgBot.Bot.Start()
}
