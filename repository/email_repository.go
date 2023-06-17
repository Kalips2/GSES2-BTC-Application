package repository

import (
	"encoding/csv"
	"github.com/pkg/errors"
	"os"
)

var (
	failToSaveEmailMessage = "Failed to subscribe email"
)

func SaveEmailToStorage(email string, pathToStorage string) error {
	var err error
	var file *os.File
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	file, err = setUpConnectionWithStorage(pathToStorage)
	if err != nil {
		return errors.New(failToSaveEmailMessage + err.Error())
	}

	err = writeToStorage(email, file)
	if err != nil {
		return errors.New(failToSaveEmailMessage + err.Error())
	}
	return err
}

func GetEmailsFromStorage(pathToStorage string) ([]string, error) {
	var file *os.File
	var err error
	var emails []string

	file, err = setUpConnectionWithStorage(pathToStorage)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	emails, err = readFromStorage(file)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return emails, err
}

func writeToStorage(email string, storage *os.File) error {
	writer := csv.NewWriter(storage)
	defer writer.Flush()
	err := writer.Write([]string{email})

	if err != nil {
		return errors.New("Failed to write into storage")
	}
	return err
}

func readFromStorage(storage *os.File) ([]string, error) {
	var data []string
	reader := csv.NewReader(storage)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("Failed to read from storage")
	}

	for _, record := range records {
		data = append(data, record[0])
	}
	return data, err
}

func setUpConnectionWithStorage(pathToStorage string) (*os.File, error) {
	storage, err := os.OpenFile(pathToStorage, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, errors.New("Failed to set up connection with storage")
	}
	return storage, err
}

func CheckEmailIsExist(email string, pathToStorage string) (bool, error) {
	var err error
	emails, err := GetEmailsFromStorage(pathToStorage)
	if err != nil {
		return false, errors.Wrap(err, "Failed to check the existence of email")
	}

	for _, existingEmail := range emails {
		if existingEmail == email {
			return true, err
		}
	}
	return false, err
}
