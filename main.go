package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(req events.APIGatewayV2HTTPRequest)(events.APIGatewayV2HTTPResponse,error){
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body: "This isn a test project",
		
	},nil //return error nil


}