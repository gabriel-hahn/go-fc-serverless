package main

import (
	"go-fc-serverless/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(services.ListProducts)
}
