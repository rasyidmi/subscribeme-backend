package helper

import (
	"context"
	"projects-subscribeme-backend/constant"

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

func SendPushNotification(data map[string]interface{}, notifType constant.NotificationEnum) error {

	fcmCLient, err := App.Messaging(context.Background())
	if err != nil {
		log.Println("error initializing app: ", err)
		return nil
	}

	messaging := &messaging.Message{
		Notification: &messaging.Notification{
			Title: data["title"].(string),
			Body:  data["body"].(string),
		},
		Token: data["token"].(string),
	}

	fcmCLient.Send(context.Background(), messaging)

	if err != nil {
		return err
	}

	return nil

}
