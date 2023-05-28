package handlers

import (
	"btcApplication/services"
	"fmt"
	"net/http"
)

func SendToEmailsHandler(w http.ResponseWriter, r *http.Request) {
	if err := services.SendToEmailsService(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Emails have been sent.")
	}
}

func SubscribeEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	if err := services.SubscribeEmailService(email); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Emails have been added.")
	}

}
