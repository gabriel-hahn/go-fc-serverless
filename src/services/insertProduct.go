package services

import (
	"go-fc-serverless/src/data"
	"go-fc-serverless/src/database"

	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createProductInput(p data.Product) (*dynamodb.PutItemInput) {
	return &dynamodb.PutItemInput{
		TableName: aws.String("ProductsVideo"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(p.ID),
			},
			"name": {
				S: aws.String(p.Name),
			},
			"price": {
				N: aws.String(strconv.Itoa(p.Price)),
			},
		},
	}
}

func InsertProduct(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var p data.Product

	err := json.Unmarshal([]byte(request.Body), &p)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	p.GenerateID()

	input := createProductInput(p)
	
	_, err = database.Insert(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	body, err := json.Marshal(p)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers: map[string]string {
			"Content-type": "application/json",
		},
		Body: string(body),
	}, nil
}
