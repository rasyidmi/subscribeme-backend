package helper

import (
	firebase "firebase.google.com/go"
	"github.com/appleboy/go-fcm"

	"log"
)

var App *firebase.App

// func InitFirebase() {
// 	opt := option.WithCredentialsFile("./serviceAccountKey.json")
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Println("error initializing app: ", err)

// 	}

// 	App = app
// }

// func SendPushNotification(data map[string]string) error {

// 	fcmCLient, err := App.Messaging(context.Background())
// 	if err != nil {
// 		log.Println("error initializing app: ", err)
// 		return nil
// 	}

// 	messaging := &messaging.Message{
// 		Notification: &messaging.Notification{
// 			Title: data["title"],
// 			Body:  data["body"],
// 		},
// 		Token: data["token"],
// 	}

// 	fcmCLient.Send(context.Background(), messaging)

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

func SendPushNotification(data map[string]string) error {
	msg := &fcm.Message{
		To: data["token"],
		Notification: &fcm.Notification{
			Title: data["title"],
			Body:  data["body"],
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("AAAAk61QKQU:APA91bHT9wIxx12PTIJ0AVw2kaliPeJl2IJG0EwNgu5N0vFT3OT9t5i6rnUDkM2RaLhhv0v5x4DXemISitWviY5TA2e6R_4pqgVChKIFeU6CRA-hNMfllFJaeI9MS4eviFYhUbOVDEHP")
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(client)

	// Send the message and receive the response without retries.
	resp, err := client.Send(msg)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(resp)

	return nil
}
