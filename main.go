package main

import (
	"btcApplication/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/rate", controllers.GetCurrentRateHandlerController)
	router.POST("/subscribe", controllers.SubscribeEmailController)
	router.POST("/sendEmails", controllers.SendToEmailsController)

	router.Run(":8080")

}
