package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Detail struct {
	EventType  string
	GeofenceId string
	DeviceId   string
	Position   []float64
}

type LocationEvent struct {
	Version    string    `json:"version"`
	ID         string    `json:"id"`
	DetailType string    `json:"detail-type"`
	Source     string    `json:"source"`
	AccountID  string    `json:"account"`
	Time       time.Time `json:"time"`
	Region     string    `json:"region"`
	Resources  []string  `json:"resources"`
	Detail     Detail    `json:"detail"`
}

func Handler(context context.Context, request events.SQSEvent) {

	event, err := json.Marshal(request.Records)
	if err != nil {
		fmt.Print("Error marshalling")
	}
	fmt.Print("Location Event " + string(event))

}
func main() {
	lambda.Start(Handler)
}
