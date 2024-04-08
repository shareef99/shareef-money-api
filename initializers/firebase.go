package initializers

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
)

var FirebaseApp *firebase.App

func InitializeFirebaseAppDefault() *firebase.App {
	FirebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return FirebaseApp
}
