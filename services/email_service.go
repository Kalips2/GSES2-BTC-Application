package services

import (
	"btcApplication/repositories"
	"btcApplication/utils"
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
)

func SendToEmailsService() error {
	var emails []string
	var err error

	emails, err = repositories.GetEmailsFromStorage()
	rate, err := GetCurrentRate()
	dialer, message := setUpMessageToSend(rate)

	for _, email := range emails {
		message.SetHeader("To", email)
		if err := dialer.DialAndSend(message); err != nil {
			continue
		}
	}

	return err
}

func SubscribeEmailService(email string) error {
	if err := utils.ValidateEmail(email); err != nil {
		return err
	}
	return repositories.SaveEmailToStorage(email)
}

func setUpMessageToSend(rate float64) (*gomail.Dialer, *gomail.Message) {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, utils.AppEmailLogin, utils.AppEmailCode)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	message := gomail.NewMessage()
	message.SetHeader("From", utils.AppEmailLogin)
	message.SetHeader("Subject", "Поточний курс "+utils.FromCurrency+" до "+utils.ToCurrency)
	message.SetBody("text/plain", "Поточний курс "+utils.FromCurrency+"до "+utils.ToCurrency+": "+fmt.Sprintf("%.5f", rate)+".")

	return dialer, message
}
