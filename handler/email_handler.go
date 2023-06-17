package handler

import (
	"btc-app/config"
	"btc-app/service"
	"fmt"
	"net/http"
)

func SendToEmailsHandler(w http.ResponseWriter, r *http.Request, c *config.Config) {
	if err := service.SendRateToEmails(c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Emails have been sent.")
	}
}

func SubscribeEmailHandler(w http.ResponseWriter, r *http.Request, c *config.Config) {
	email := r.FormValue("email")

	if err := service.SubscribeEmail(email, c); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Emails have been added.")
	}

}
