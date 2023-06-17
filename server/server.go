package server

import (
	"btc-app/config"
	"btc-app/handler"
	"github.com/go-chi/chi"
	"net/http"
)

type Server struct {
	Config config.Config
	Router *chi.Mux
}

func NewServer(conf config.Config) *Server {
	return &Server{
		Config: conf,
		Router: chi.NewRouter(),
	}
}

func (s *Server) InitHandlers() {
	s.Router.Get("/rate", func(w http.ResponseWriter, r *http.Request) {
		handler.GetCurrentRateHandler(w, r, &s.Config)
	})

	s.Router.Post("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		handler.SubscribeEmailHandler(w, r, &s.Config)
	})

	s.Router.Post("/sendEmails", func(w http.ResponseWriter, r *http.Request) {
		handler.SendToEmailsHandler(w, r, &s.Config)
	})
	_ = http.ListenAndServe(":8080", s.Router)
}
