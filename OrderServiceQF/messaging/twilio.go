package messaging

import (
	"OrderServiceQF/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	// "github.com/kevinburke/twilio-go"
)

func TwilioServeSms(orderProto *models.KafkaMsg) error {
	err := godotenv.Load("./secret/.env")
	if err != nil {
		log.Println("ENVFILE", err)
	}

	body := fmt.Sprintf("QuickFixx- %s is interested in your service, connect with him ", orderProto.Contact)

	accountSID := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(body)

	// Send SMS message
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println("Error sending SMS message:", err)
		return err
	}

	log.Println("SMS message sent successfully. SID:", resp.Sid)

	return nil
}
