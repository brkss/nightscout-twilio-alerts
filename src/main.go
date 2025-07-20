package main

import (
	"fmt"
	"log"

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

	fmt.Println("account sid : ", config.AccountSID)

	/*
		accountSid := ""
		authToken := ""
		twilioNumber := ""   // Your Twilio number
		personalNumber := "" // Your personal phone number

		client := twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		})

		params := &openapi.CreateCallParams{}
		params.SetTo(personalNumber)
		params.SetFrom(twilioNumber)
		params.SetUrl("http://demo.twilio.com/docs/voice.xml") // TwiML instructions

		resp, err := client.Api.CreateCall(params)
		if err != nil {
			fmt.Println("❌ Error making call:", err.Error())
			os.Exit(1)
		}

		fmt.Println("✅ Call initiated! SID:", *resp.Sid)
	*/
}
