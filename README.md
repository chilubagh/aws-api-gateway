# AWS Lambda for Go

Lambda is a compute service that lets you run code without provisioning or managing servers. Lambda runs your code on a high-availability compute infrastructure and performs all of the administration of the compute resources, including server and operating system maintenance, capacity provisioning and automatic scaling, code monitoring and logging. With Lambda, you can run code for virtually any type of application or backend service. All you need to do is supply your code in one of the languages that Lambda supports.
I deploying a Rest API using AWS,API Gateway and lambda functions using golang.

# Getting Started
```// main.go
package main

import (
	"github.com/aws/aws-lambda-go/lambda" 
)

func hello() (string, error) {
	return "Hello ƛ!", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
```

# Building your function(Linux and macOS)
Preparing a binary to deploy to AWS Lambda requires that it is compiled for Linux and placed into a .zip file.<br>
Remember to build your handler executable for Linux!<br>
GOOS=linux GOARCH=amd64 go build -o main main.go<br>
zip main.zip main<br>

#Function Deployment
The function was deployed using then official [AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/deploying-lambda-apps.html)


