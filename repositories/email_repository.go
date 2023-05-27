package repositories

import (
	"btcApplication/utils"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"os"
)

func SaveEmailToStorage(email string) error {

	if checkEmailIsExist(email) {
		return errors.New("Email is already exist")
	}

	file, _ := os.OpenFile(utils.SubscriptionFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{encodeEmail(email)})

	return nil
}

func GetEmailsFromStorage() []string {
	file, _ := os.Open(utils.SubscriptionFilePath)
	defer file.Close()

	reader := csv.NewReader(file)
	emails, _ := reader.ReadAll()

	var emailList []string
	for _, encodedEmail := range emails {
		decodedEmail := decodeEmail(encodedEmail[0])
		emailList = append(emailList, decodedEmail)
	}

	return emailList
}

func checkEmailIsExist(email string) bool {
	emails := GetEmailsFromStorage()

	for _, existingEmail := range emails {
		if existingEmail == email {
			return true
		}
	}
	return false
}

func encodeEmail(email string) string {
	encodedEmail := base64.StdEncoding.EncodeToString([]byte(email))
	return encodedEmail
}
func decodeEmail(encodedEmail string) string {
	decodedEmail, _ := base64.StdEncoding.DecodeString(encodedEmail)

	return string(decodedEmail)
}
