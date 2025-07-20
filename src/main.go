package main

import (
	"log"

	"github.com/brkss/nightscout-twillio-alerts/src/nightscout"
	"github.com/brkss/nightscout-twillio-alerts/src/twilio"
	config "github.com/brkss/nightscout-twillio-alerts/src/utils"
	/*
		"os"

		twilio "github.com/twilio/twilio-go"
		openapi "github.com/twilio/twilio-go/rest/api/v2010"
	*/)

func main() {

	config, err := config.NewConfig()
	if err != nil {
		log.Panic("Invalid Config File ! : ", err)
	}

	twilioService := twilio.NewTwilioService(*config)
	nightscoutService := nightscout.NewNightscoutService(*config, *twilioService)

	// run nightscout service loop
	nightscoutService.NightscoutBloodSugarCheckRoutine()

}
