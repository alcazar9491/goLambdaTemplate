package main

import (
	// "fmt"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type inputEvent struct {
	ApiToken      string `json:"token"`
	TokenAddress  string `json:"tokenAddress"`
}


func main() {
	lambda.Start(Handler)
}

func Handler( event events.APIGatewayProxyRequest ) ( events.APIGatewayProxyResponse, error ) {

	var reqBody inputEvent

	response := events.APIGatewayProxyResponse {
		Headers: map[string]string{
			"Content-Type":		"application/json",
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST",
		},
	}

	err := json.Unmarshal( []byte(event.Body), &reqBody )

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		return response, err
	}

	


	response.StatusCode = http.StatusOK

	return response, nil
}