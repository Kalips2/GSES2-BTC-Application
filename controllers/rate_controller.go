package controllers

import (
	"btcApplication/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentRateHandlerController(c *gin.Context) {
	rate, err := services.GetCurrentRate()
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to get current rate. Try again!")
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("%.5f", rate))
}
