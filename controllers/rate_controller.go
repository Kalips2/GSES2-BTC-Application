package controllers

import (
	"btcApplication/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentRateHandlerController(c *gin.Context) {
	if rate, err := services.GetCurrentRate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, fmt.Sprintf("%.5f", rate))
	}
}
