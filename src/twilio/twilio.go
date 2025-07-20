package twilio

import (
	"fmt"

	config "github.com/brkss/nightscout-twillio-alerts/src/utils"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioService struct {
	config config.Config
}

func NewTwilioService(config config.Config) *TwilioService {
	return &TwilioService{
		config,
	}
}

func (ts TwilioService) CallUrgentLow() error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: ts.config.AccountSID,
		Password: ts.config.AuthToken,
	})

	params := &openapi.CreateCallParams{}
	params.SetTo(ts.config.PersonalNumber)
	params.SetFrom(ts.config.TwilioNumber)
	params.SetUrl("http://demo.twilio.com/docs/voice.xml") // TwiML instructions

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		return err
	}

	fmt.Println("✅ Call initiated! SID:", *resp.Sid)

	return nil
}
