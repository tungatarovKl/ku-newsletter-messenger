package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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

		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Request parsing error", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(requestBody, &apiRequest)
		if err != nil {
			http.Error(rw, "Request parsing error", http.StatusBadRequest)
			return
		}

		//Check whether the proper token was sent
		if checked, err := bot.Database.ValidateToken(apiRequest.Token); checked == false {
			if err != nil {
				log.Println(err.Error())
				http.Error(rw, "Dependency error", http.StatusFailedDependency)
				return
			}
			http.Error(rw, "Access denied", http.StatusNetworkAuthenticationRequired)
			return
		}

		if apiRequest.Message == "" {
			http.Error(rw, "Message is empty", http.StatusBadRequest)
			return
		}

		users, err := bot.Database.GetAllUsers()
		if err != nil {
			log.Println(err.Error())
			http.Error(rw, "Dependency error", http.StatusFailedDependency)
			return
		}

		if len(users) == 0 {
			log.Println("Users table is empty!")
			http.Error(rw, "Recipient list is empty", http.StatusExpectationFailed)
			return
		}

		//Send message to all registered users
		var sendWG sync.WaitGroup
		sendWG.Add(len(users))
		sendErr := make(chan error, len(users)) //Channel to store errors

		for _, user := range users {

			//New routine
			go func(u models.User, errChan chan error) {

				_, sErr := bot.SendMessage(u.TelegramId, apiRequest.Message)

				if sErr != nil {
					errChan <- sErr //Send error to channel
				}
				sendWG.Done()

			}(user, sendErr)
		}

		sendWG.Wait()
		close(sendErr)

		//Count percent of faults
		sendErrPercent := len(sendErr) * 100 / len(users)

		//Send error response if more than 50% of faults
		if sendErrPercent > 50 {
			http.Error(rw, "Sent to "+strconv.Itoa(100-sendErrPercent)+"% of users", http.StatusFailedDependency)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Sent to " + strconv.Itoa(100-sendErrPercent) + "% of users"))
	}
}
