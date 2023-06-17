package service

import (
	"btc-app/config"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

var (
	failToGetRateMessage = "Failed get current rate"
)

func GetCurrentRate(c *config.Config) (float64, error) {
	var resp *http.Response
	var err error
	var rate float64

	url := c.CryptoApiURL + "?fsym=" + c.CurrencyFrom + "&tsyms=" + c.CurrencyTo
	if resp, err = http.Get(url); err != nil {
		return 0, errors.Wrap(err, failToGetRateMessage)
	}

	if rate, err = getRateFromHttpResponse(resp); err != nil {
		return 0, errors.Wrap(err, failToGetRateMessage)
	}
	return rate, err
}

func getRateFromHttpResponse(resp *http.Response) (float64, error) {
	var err error
	var data map[string]float64
	var rate float64

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New("Failed to read body of response")
	}
	err = json.Unmarshal(body, &data)
	rate = data["UAH"]
	return rate, err
}
