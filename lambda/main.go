package main

import (
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"github.com/BTBurke/twiml"
	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/api/option"
)

type TwilioSMSEvent struct {
	From string `json:"From"`
	Body string `json:"Body"`
}

func updateBaby(ctx context.Context, who string, status string) error {
	opt := option.WithCredentialsFile("./credentials.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}

	defer client.Close()

	babies := client.Collection("babies")
	curr := babies.Doc(who)
	_, err = curr.Update(ctx, []firestore.Update{{Path: "status", Value: status}})
	return err
}

func HandleRequest(ctx context.Context, sms TwilioSMSEvent) ([]byte, error) {
	err := updateBaby(ctx, "lukeandrebecca", sms.Body)
	if err != nil {
		return nil, err
	}

	res := twiml.NewResponse()
	return res.Encode()
}

func main() {
	lambda.Start(HandleRequest)
}
