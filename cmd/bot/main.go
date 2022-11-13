package main

import (
	"flag"
	"log"
	"upgrade/cmd/bot/Bot"
	"upgrade/config"
	"upgrade/internal/models"
)

func main() {
	println("BOT")
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()
	cfg := config.ReadConfig(*configPath)
	database := models.NewDatabase(cfg.DbAddress, cfg.DbName, cfg.DbUsername, cfg.DbPassword)
	if database.Connection.Error != nil {
		log.Fatal("Error connecting to database:", database.Connection.Error)
	}
	tgBot := Bot.Bot{
		Bot:      Bot.InitBot(cfg.BotToken),
		Database: database,
	}
	tgBot.Bot.Handle("/start", tgBot.StartHandler)
	tgBot.Bot.Start()
}
