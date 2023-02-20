package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SnsEvent struct {
	Name string `json:"name"`
}

func main() {

	snsEvent := SnsEvent{Name: "Sayem"}
	snsMessage, err := json.Marshal(snsEvent)
	if err != nil {
		fmt.Print("Failed parsing json!")
		return
	}
	fmt.Println(string(snsMessage))

	cfg, error := config.LoadDefaultConfig(context.TODO())
	if error != nil {
		fmt.Println("===error===")
		fmt.Print(error)
		return
	}

	client := sns.NewFromConfig(cfg)
	topicArn := "arn:aws:sns:us-east-2:682986148056:GoLambdaTestSNS"
	response, error := client.Publish(
		context.TODO(),
		&sns.PublishInput{Message: aws.String(string(snsMessage)), TopicArn: &topicArn},
	)
	if error != nil {
		fmt.Println("===error===")
		fmt.Print(error)
		return
	}

	fmt.Println("===response===")
	fmt.Println(response)
}
