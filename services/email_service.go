package services

import (
	"btcApplication/repositories"
	"btcApplication/utils"
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
)

func SendToEmailsService() {
	emails := repositories.GetEmailsFromStorage()

	rate, _ := GetCurrentRate()

	dialer, message := setUpMessageToSend(rate)

	for _, email := range emails {
		message.SetHeader("To", email)
		if err := dialer.DialAndSend(message); err != nil {
			continue
		}
	}

}

func SubscribeEmailService(email string) error {
	return repositories.SaveEmailToStorage(email)
}

func setUpMessageToSend(rate float64) (*gomail.Dialer, *gomail.Message) {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, utils.AppEmailLogin, utils.AppEmailCode)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	message := gomail.NewMessage()
	message.SetHeader("From", utils.AppEmailLogin)
	message.SetHeader("Subject", "Поточний курс BTC до UAH")
	message.SetBody("text/plain", "Поточний курс BTC до UAH: "+fmt.Sprintf("%.5f", rate)+".")

	return dialer, message
}
