package handlers

import (
	"btcApplication/services"
	"fmt"
	"net/http"
)

func GetCurrentRateHandler(w http.ResponseWriter, r *http.Request) {
	if rate, err := services.GetCurrentRate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "%.5f", rate)
	}
}
