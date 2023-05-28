package repositories

import (
	"btcApplication/utils"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"os"
)

func SaveEmailToStorage(email string) error {
	if err := checkEmailIsExist(email); err != nil {
		return err
	}

	var err error
	file, err := os.OpenFile(utils.SubscriptionFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	defer func(file *os.File) {
		err = file.Close()
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write([]string{encodeEmail(email)})
	return err
}

func GetEmailsFromStorage() ([]string, error) {
	var file *os.File
	var err error
	if file, err = os.Open(utils.SubscriptionFilePath); err != nil {
		return []string{}, err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	reader := csv.NewReader(file)
	emails, err := reader.ReadAll()

	var emailList []string
	for _, encodedEmail := range emails {
		decodedEmail := decodeEmail(encodedEmail[0])
		emailList = append(emailList, decodedEmail)
	}

	return emailList, err
}

func checkEmailIsExist(email string) error {
	emails, err := GetEmailsFromStorage()

	for _, existingEmail := range emails {
		if existingEmail == email {
			return errors.New("Email is already subscribed.")
		}
	}
	return err
}

func encodeEmail(email string) string {
	encodedEmail := base64.StdEncoding.EncodeToString([]byte(email))
	return encodedEmail
}
func decodeEmail(encodedEmail string) string {
	decodedEmail, _ := base64.StdEncoding.DecodeString(encodedEmail)

	return string(decodedEmail)
}
