# AWS Lambda for Go

Lambda is a compute service that lets you run code without provisioning or managing servers. Lambda runs your code on a high-availability compute infrastructure and performs all of the administration of the compute resources, including server and operating system maintenance, capacity provisioning and automatic scaling, code monitoring and logging. With Lambda, you can run code for virtually any type of application or backend service. All you need to do is supply your code in one of the languages that Lambda supports.
I deploying a Rest API using AWS,API Gateway and lambda functions using golang.

# Getting Started
// main.go
package main

import (
	"github.com/aws/aws-lambda-go/lambda" <br>
)<br>

func hello() (string, error) {<br>
	return "Hello ƛ!", nil<br>
}<br>

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda<br>
	lambda.Start(hello)<br>
}
