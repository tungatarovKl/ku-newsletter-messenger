package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"upgrade/internal/bot"
	"upgrade/internal/models"
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

		if bot.Database.CheckToken(apiRequest.Token) == nil {
			http.Error(rw, "Invalid token entered", http.StatusBadRequest)
			return
		}

		if apiRequest.Message == "" {
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
		var sendWG sync.WaitGroup
		sendWG.Add(len(users))
		for _, user := range users {
			//New routine
			go func(u models.User) {
				bot.SendMessage(u.TelegramId, apiRequest.Message)
				sendWG.Done()
			}(user)
		}
		//Waiting for all routines to complete
		sendWG.Wait()

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("OK"))
	}
}
