package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"upgrade/internal/bot"
)

type ApiRequest struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func NewsLetterPost(bot bot.Bot) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		var apiRequest ApiRequest

		requestBody, readErr := io.ReadAll(r.Body)
		if readErr != nil {
			http.Error(rw, "Request parsing error", http.StatusBadRequest)
			return
		}

		parseErr := json.Unmarshal(requestBody, &apiRequest)
		if parseErr != nil {
			http.Error(rw, "Request parsing error", http.StatusBadRequest)
			return
		}

		if apiRequest.Message == "" || apiRequest.Token == "" {
			http.Error(rw, "Some required fields are absent", http.StatusBadRequest)
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
			bot.SendMessage(user.TelegramId, apiRequest.Message)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("OK"))
	}
}
