package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event events.SNSEvent) {
	for _, record := range event.Records {
		snsRecord := record.SNS
		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		var person Person
		error := json.Unmarshal([]byte(snsRecord.Message), &person)
		if error != nil {
			fmt.Println("======error=======")
			fmt.Println(error)
			continue
		}
		spew.Dump(person)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
