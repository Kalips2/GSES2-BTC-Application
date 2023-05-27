package controllers

import (
	"btcApplication/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendToEmailsController(context *gin.Context) {
	services.SendToEmailsService()
	context.String(http.StatusOK, "E-mailʼи відправлено")
}

func SubscribeEmailController(context *gin.Context) {
	email := context.Query("email")

	if err := services.SubscribeEmailService(email); err != nil {
		context.String(http.StatusConflict, err.Error())
	} else {
		context.String(http.StatusOK, "E-mail додано")
	}

}
