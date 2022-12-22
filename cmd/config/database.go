package config

import (
	firebase "firebase.google.com/go"

	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

// firestore
func ConnectFirestore() (*firestore.Client, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("/home/syam/go/src/unary_grpc/grpc-golang-kotakode-firebase-adminsdk-j36qt-7f15aab435.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	//client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	return client, nil
}
