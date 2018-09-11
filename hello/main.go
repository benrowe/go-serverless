package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// import "encoding/json"

// Response
type Response events.APIGatewayProxyResponse

// Request
type Request events.APIGatewayProxyRequest

// MyEvent Request Structure
type MyEvent struct {
	Name string `json:"name"`
}

// HandleRequest for lambda
func HandleRequest(ctx context.Context, request Request) (Response, error) {
	fmt.Println("Request handling started....")

	fmt.Println(request.Body)

	var event MyEvent

	err := json.Unmarshal([]byte(request.Body), &event)
	if err != nil {
		fmt.Println("Unmarshalling Error: ", err)
		return Response{
			Body:       `{ "message": "Something went wrong." }`,
			StatusCode: 500,
		}, nil
	}
	fmt.Println(event)
	var buf bytes.Buffer
	resp := fmt.Sprintf("Hello, %s", event.Name)

	body, err := json.Marshal(map[string]interface{}{
		"message": resp,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)
	response := Response{
		Body:       buf.String(),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
