package controllers

import (
	"log"
	"net/http"
	"upgrade/internal/bot"
)

func NewsLetterPost(bot bot.Bot) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		err := r.ParseForm()
		if err != nil {
			http.Error(rw, "Request parsing error", http.StatusBadRequest)
			return
		}

		message := r.Form.Get("message")
		if message == "" {
			http.Error(rw, "'message' key required'", http.StatusBadRequest)
			return
		}

		users, err := bot.Database.GetAllUsers()
		if err != nil {
			log.Println(err.Error())
			http.Error(rw, "Dependency error", http.StatusFailedDependency)
			return
		}

		//Send message for all registered users
		for _, user := range users {
			bot.SendMessage(user.TelegramId, message)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("OK"))
	}
}
