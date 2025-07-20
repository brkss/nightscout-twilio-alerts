package nightscout

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/brkss/nightscout-twillio-alerts/src/twilio"
	config "github.com/brkss/nightscout-twillio-alerts/src/utils"
)

type NightscoutEntry struct {
	Date       int64   `json:"date"`       // Timestamp in ms
	Sgv        float64 `json:"sgv"`        // Blood sugar value in mg/dL
	DateString string  `json:"dateString"` // Optional: ISO string
}

type NightscoutService struct {
	config        config.Config
	twilioService twilio.TwilioService
	// add tewillio service
}

func NewNightscoutService(config config.Config, twilioService twilio.TwilioService) *NightscoutService {
	return &NightscoutService{
		config,
		twilioService,
	}
}

func (ns *NightscoutService) fetchLatestEntry() (*NightscoutEntry, error) {
	nightscoutEP := fmt.Sprintf("%s/api/v1/entries.json?count=1", ns.config.NightscoutURL)
	response, err := http.Get(nightscoutEP)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var entries []NightscoutEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("no entry returned")
	}

	return &entries[0], nil
}

func (ns *NightscoutService) NightscoutBloodSugarCheckRoutine() {
	for {
		entry, err := ns.fetchLatestEntry()
		if err != nil {
			fmt.Println("Error fetching data:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		timestamp := time.UnixMilli(entry.Date)
		fmt.Printf("Fetched at: %v | SGV: %v\n", timestamp.Format(time.RFC3339), entry.Sgv)

		if entry.Sgv < 70 {
			// call twillio service !
			fmt.Println("⚠️ Alert: Blood sugar is low (< 80 mg/dL)")

			err := ns.twilioService.CallUrgentLow()
			if err != nil {
				fmt.Println("============= 🚨 ERROR ACCURED WHILE CALLING =============")
			}

		}

		nextFetch := timestamp.Add(5 * time.Minute)
		sleepDuration := time.Until(nextFetch)

		if sleepDuration < 0 {
			sleepDuration = 10 * time.Second // In case of clock skew
		}

		fmt.Printf("Sleeping until %s...\n", nextFetch.Format(time.RFC3339))
		time.Sleep(sleepDuration)
	}
}
