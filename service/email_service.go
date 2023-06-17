package service

import (
	"btc-app/config"
	"btc-app/repository"
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
	"regexp"
)

var (
	emailRegex                  = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	ErrEmailIsAlreadySubscribed = errors.New("Email is already subscribed!")
	failToSendRateMessage       = "Failed to send the rate to emails"
	failToSubscribeMessage      = "Failed to subscribe email"
)

func SendRateToEmails(c *config.Config) error {
	var emails []string
	var err error

	emails, err = repository.GetEmailsFromStorage(c.EmailStoragePath)
	if err != nil {
		return errors.Wrap(err, failToSendRateMessage)
	}

	rate, err := GetCurrentRate(c)
	if err != nil {
		return errors.Wrap(err, failToSendRateMessage)
	}

	dialer, message := setUpMessageToSend(rate, c)
	err = sendMessageToEmails(message, emails, dialer)
	if err != nil {
		return errors.Wrap(err, failToSendRateMessage)
	}
	return err
}

func SubscribeEmail(email string, c *config.Config) error {
	var err error

	err = validateEmail(email)
	if err != nil {
		return errors.Wrap(err, failToSubscribeMessage)
	}

	exist, err := repository.CheckEmailIsExist(email, c.EmailStoragePath)
	if exist {
		err = ErrEmailIsAlreadySubscribed
	}
	if err != nil {
		return errors.Wrap(err, failToSubscribeMessage)
	}

	err = repository.SaveEmailToStorage(email, c.EmailStoragePath)
	if err != nil {
		return errors.Wrap(err, failToSubscribeMessage)
	}
	return err
}

func validateEmail(email string) error {
	var err error
	if !emailRegex.MatchString(email) {
		err = errors.New("Email doesn't match regex: " + emailRegex.String())
	}
	return err
}

func setUpMessageToSend(rate float64, c *config.Config) (*gomail.Dialer, *gomail.Message) {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, c.EmailServiceFrom, c.EmailServicePassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	message := gomail.NewMessage()
	message.SetHeader("From", c.EmailServiceFrom)
	message.SetHeader("Subject", c.EmailServiceSubject)
	message.SetBody("text/plain", "Поточний курс "+c.CurrencyFrom+" до "+c.CurrencyTo+": "+fmt.Sprintf("%.5f", rate)+".")

	return dialer, message
}

func sendMessageToEmails(message *gomail.Message, emails []string, dialer *gomail.Dialer) error {
	var err error
	for _, email := range emails {
		message.SetHeader("To", email)
		if err = dialer.DialAndSend(message); err != nil {
			continue
		}
	}
	return err
}
