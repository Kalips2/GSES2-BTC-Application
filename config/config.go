package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port                 int
	CryptoApiURL         string
	CurrencyFrom         string
	CurrencyTo           string
	EmailStoragePath     string
	EmailServiceHost     string
	EmailServicePort     string
	EmailServiceSubject  string
	EmailServiceFrom     string
	EmailServicePassword string
}

type varToField struct {
	varName string
	field   interface{}
}

func (c *Config) InitConfigFromEnv() {
	requiredEnvVars := initRequiredVars(c)

	for _, envVar := range requiredEnvVars {
		value := os.Getenv(envVar.varName)
		switch field := envVar.field.(type) {
		case *int:
			*field, _ = strconv.Atoi(value)
		case *string:
			*field = value
		}
	}
}

func initRequiredVars(c *Config) []varToField {
	return []varToField{
		{"PORT", &c.Port},
		{"CRYPTO_API_URL", &c.CryptoApiURL},
		{"CURRENCY_FROM", &c.CurrencyFrom},
		{"CURRENCY_TO", &c.CurrencyTo},
		{"EMAIL_STORAGE_PATH", &c.EmailStoragePath},
		{"EMAIL_SERVICE_HOST", &c.EmailServiceHost},
		{"EMAIL_SERVICE_PORT", &c.EmailServicePort},
		{"EMAIL_SERVICE_SUBJECT", &c.EmailServiceSubject},
		{"EMAIL_SERVICE_FROM", &c.EmailServiceFrom},
		{"EMAIL_SERVICE_PASSWORD", &c.EmailServicePassword},
	}
}
