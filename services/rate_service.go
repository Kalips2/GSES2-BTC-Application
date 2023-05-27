package services

import (
	"btcApplication/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetCurrentRate() (float64, error) {
	resp, err := http.Get(utils.CryptoCompareApiURL + "?fsym=BTC&tsyms=UAH")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data map[string]float64
	json.Unmarshal(body, &data)

	rate, ok := data["UAH"]
	if !ok {
		return 0, errors.New("Failed to get UAH rate")
	}

	return rate, nil
}
