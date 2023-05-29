package services

import (
	"go-fc-serverless/internal/data"
	"go-fc-serverless/internal/database"

	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func ListProducts(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := database.GetAll("ProductsVideo")
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	var products []data.Product
	for _, item := range result.Items {
		price, err := strconv.Atoi(*item["price"].N)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		products = append(products, data.Product{
			ID: *item["id"].S,
			Name: *item["name"].S,
			Price: price,
		})
	}

	body, err := json.Marshal(products)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string {
			"Content-type": "application/json",
		},
		Body: string(body),
	}, nil
}
