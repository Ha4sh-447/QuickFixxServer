package messaging

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/joho/godotenv"

	"google.golang.org/api/option"
)

func SendMessage(token string) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	opt := option.WithCredentialsFile(os.Getenv("ADMIN_SDK_KEY"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println("error initializing app: %v", err)
	}

	fcmApp, err := app.Messaging(context.Background())
	if err != nil {
		log.Println("ERROR MESSAGING", err)
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Hello from go server",
			Body:  "THIS IS YOUR BODY",
		},
		Token: os.Getenv("DEVICE_TOKEN"), //just a work around
	}

	response, err := fcmApp.Send(context.Background(), message)
	if err != nil {
		log.Println("SEND MESSAGE", err)
		return "", err
	}

	return response, nil
}
