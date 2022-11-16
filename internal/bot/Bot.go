package bot

import (
	"gopkg.in/telebot.v3"
	"log"
	"net/http"
	"time"
	"upgrade/internal/models"
)

type Bot struct {
	Database *models.Database
	Bot      *telebot.Bot
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		//Timeout to send message
		Client: &http.Client{Timeout: 2 * time.Second},
	}

	b, err := telebot.NewBot(pref)

	if err != nil {
		log.Fatalf("Bot initialization error %v", err)
	}
	return b
}

// User registration
func (bot *Bot) StartHandler(ctx telebot.Context) error {
	err := bot.Database.NewUser(ctx.Sender().ID, ctx.Sender().FirstName, ctx.Sender().LastName, ctx.Chat().ID)
	if err != nil {
		log.Println("Error creating user:", err.Error())
	}
	return ctx.Send("Привет, " + ctx.Sender().FirstName)
}

// Send message by chat ID
func (bot *Bot) SendMessage(chatId int64, message string) (*telebot.Message, error) {
	chat := telebot.Chat{ID: chatId}
	return bot.Bot.Send(&chat, message)
}
