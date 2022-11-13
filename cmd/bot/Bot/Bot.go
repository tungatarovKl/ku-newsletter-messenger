package Bot

import (
	"gopkg.in/telebot.v3"
	"log"
	"time"
	"upgrade/Model/Repository"
)

type Bot struct {
	Database *Repository.Database
	Bot      *telebot.Bot
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}
	return b
}

func (bot *Bot) StartHandler(ctx telebot.Context) error {
	err := bot.Database.NewUser(ctx.Sender().ID, ctx.Sender().FirstName, ctx.Sender().LastName, ctx.Chat().ID)
	if err != nil {
		log.Println("Error creating user:", err.Error())
	}
	return ctx.Send("Привет, " + ctx.Sender().FirstName)
}
