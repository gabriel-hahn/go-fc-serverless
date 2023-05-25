package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getDatabaseSession() (*dynamodb.DynamoDB) {
	sess := session.Must(session.NewSession())
	return dynamodb.New(sess)
}

func Insert(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	svc := getDatabaseSession()
	return svc.PutItem(input)
}

func GetAll(tableName string)(*dynamodb.ScanOutput, error) {
	svc := getDatabaseSession()

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	return svc.Scan(input)
}
