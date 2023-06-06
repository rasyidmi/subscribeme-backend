package helper

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"log"

	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println("error initializing app: ", err)

	}

	App = app
}

func SendPushNotification(data map[string]string) error {

	fcmCLient, err := App.Messaging(context.Background())
	if err != nil {
		log.Println("error initializing app: ", err)
		return nil
	}

	messaging := &messaging.Message{
		Notification: &messaging.Notification{
			Title: data["title"],
			Body:  data["body"],
		},
		Token: data["token"],
	}

	fcmCLient.Send(context.Background(), messaging)

	if err != nil {
		return err
	}

	return nil

}
