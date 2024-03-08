package messaging

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/joho/godotenv"

	"google.golang.org/api/option"
)

func SendMessage(token string, spName string, date string) (string, error) {

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

	// parsedDate, err := time.Parse("2006-01-02", date)
	// if err != nil {
	// 	return "", fmt.Errorf("error parsing date: %v", err)
	// }

	// // Format the date in a human-readable format
	// formattedDate := parsedDate.Format("Monday, January 2, 2006")

	body := fmt.Sprintf("%s will contact you shortly", spName)

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Quickfixx",
			Body:  body,
		},
		Token: os.Getenv("DEVICE_TOKEN"), //just a work around
	}

	fmt.Println(os.Getenv("DEVICE_TOKEN"))

	response, err := fcmApp.Send(context.Background(), message)
	if err != nil {
		log.Println("SEND MESSAGE", err)
		return "", err
	}

	return response, nil
}
