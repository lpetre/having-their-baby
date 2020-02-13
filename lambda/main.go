package main

import (
	"fmt"
	"net/url"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"github.com/BTBurke/twiml"
	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/api/option"
)

type TwilioSMSEvent struct {
	From string `json:"From"`
	Body string `json:"Body"`
}

func updateBaby(ctx context.Context, from string, status string) (int, error) {
	opt := option.WithCredentialsFile("./credentials.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return 0, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return 0, err
	}

	defer client.Close()

	ownerRef, err := client.Collection("numbers").Doc(from).Get(ctx)
	if err != nil {
		return 0, err
	}
	ownerData := ownerRef.Data()
	babies, foundBabies := ownerData["babies"]
	if !foundBabies {
		return 0, nil
	}

	batch := client.Batch()

	babiesMap := babies.([]interface{})
	for _, baby := range babiesMap {
		babyRef := client.Collection("babies").Doc(baby.(string))
		batch.Set(babyRef, map[string]interface{}{
			"status": status,
		})
	}

	_, err = batch.Commit(ctx)
	return len(babiesMap), err
}

func HandleRequest(ctx context.Context, sms TwilioSMSEvent) (string, error) {
	from, err := url.QueryUnescape(sms.From)
	if err != nil {
		return "", err
	}

	status, err := url.QueryUnescape(sms.Body)
	if err != nil {
		return "", err
	}

	count, err := updateBaby(ctx, from, status)
	if err != nil {
		return "", err
	}

	res := twiml.NewResponse()
	res.Add(&twiml.Sms{
		Text: fmt.Sprintf("Updated %d sites", count),
	})
	return res.String()
}

func main() {
	lambda.Start(HandleRequest)
}
