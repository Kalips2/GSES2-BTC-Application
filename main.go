package main

import (
	"btcApplication/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/rate", handlers.GetCurrentRateHandler)
	r.Post("/subscribe", handlers.SubscribeEmailHandler)
	r.Post("/sendEmails", handlers.SendToEmailsHandler)

	http.ListenAndServe(":8080", r)

}
