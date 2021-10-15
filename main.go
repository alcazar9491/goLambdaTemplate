package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type inputEvent struct {
	TokenAddress  string `json:"tokenAddress"`
	UrlBase		  string `json:"urlBase"`
	Action        string `json:"action"`
	ApiToken        string `json:"apiToken"`
	
}

type Response struct {
	ErrorData string `json:"errorData"`
	Message string 	 `json:"message"`
	Result string 	 `json:"result"`
	Status string 	 `json:"status"`
}



func main() {
	lambda.Start(Handler)
}

func Handler( event events.APIGatewayProxyRequest ) ( events.APIGatewayProxyResponse, error ) {
	
	// init response object
	response := events.APIGatewayProxyResponse {
		Headers: map[string]string{
			"Content-Type":					"application/json",
			"Access-Control-Allow-Origin": 	"*",
			"Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST",
		},
	}

	//Get request body
	var reqBody inputEvent
	err := json.Unmarshal( []byte(event.Body), &reqBody )
	if err != nil {
		response.StatusCode = http.StatusBadRequest
		return response, err
	}


	//get data
	bodyBytes := apiCall( reqBody.UrlBase, reqBody.TokenAddress, reqBody.Action, reqBody.ApiToken )

	// init response body
	var responseBody Response
	json.Unmarshal(bodyBytes, &responseBody)
	out, _ := json.Marshal(responseBody)

	response.Body = string(out)
	response.StatusCode = http.StatusOK

	return response, nil
}


func apiCall( urlBase string, token string, action string, apiToken string ) []byte {

	url:= fmt.Sprintf("%s/api?module=stats&action=%s&contractaddress=%s&apikey=%s`",urlBase,action,token,apiToken)
	
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()


	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	return bodyBytes


}