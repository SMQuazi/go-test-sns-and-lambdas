package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type SnsEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event SnsEvent) {
	fmt.Printf("Event: %s!", event)
	fmt.Printf("Name: %s!", event.Name)
}

func main() {
	lambda.Start(HandleRequest)
}
