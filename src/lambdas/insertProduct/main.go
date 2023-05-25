package main

import (
	"go-fc-serverless/src/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(services.InsertProduct)
}
