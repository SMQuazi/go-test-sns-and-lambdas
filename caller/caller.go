package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/davecgh/go-spew/spew"
)

type SnsEvent struct {
	Name string `json:"name"`
}

func main() {

	cfg, error := config.LoadDefaultConfig(context.TODO())
	if error != nil {
		fmt.Println("===error===")
		fmt.Print(error)
		return
	}

	snsEvent := SnsEvent{Name: "John"}
	snsMessage, err := json.Marshal(snsEvent)
	if err != nil {
		fmt.Print("Failed parsing json!")
		return
	}
	snsString := string(snsMessage)

	client := sns.NewFromConfig(cfg)
	topicArn := "arn:aws:sns:us-east-2:682986148056:GoLambdaTestSNS"

	publishInput := sns.PublishInput{
		Message:  &snsString,
		TopicArn: &topicArn,
	}

	spew.Dump(publishInput)

	response, error := client.Publish(
		context.TODO(),
		&publishInput,
	)
	if error != nil {
		fmt.Println("===error===")
		fmt.Print(error)
		return
	}

	fmt.Println("===response===")
	spew.Dump(response.MessageId)
}
