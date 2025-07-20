package config

import (
	"os"

	"github.com/joho/godotenv"
)



type Config struct {
	AccountSID string;
	AuthToken string;
	TwilioNumber string;
	PersonalNumber string;
	NightscoutURL string;
}


func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		AccountSID: os.Getenv("ACCOUNT_SID"),
		AuthToken: os.Getenv("AUTH_TOKEN"),
		TwilioNumber: os.Getenv("TWILLIO_NUMBER"),
		PersonalNumber: os.Getenv("PERSONAL_NUMBER"),
		NightscoutURL: os.Getenv("NIGHTSCOUT_URL"),
	}, nil
}