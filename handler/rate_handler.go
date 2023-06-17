package handler

import (
	"btc-app/config"
	"btc-app/service"
	"fmt"
	"net/http"
)

func GetCurrentRateHandler(w http.ResponseWriter, r *http.Request, c *config.Config) {
	if rate, err := service.GetCurrentRate(c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "%.5f", rate)
	}
}
