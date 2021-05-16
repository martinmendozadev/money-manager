package utils

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DynamoDBClient -
func dynamoDBClient() *dynamodb.DynamoDB {
	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return svc
}

// MarshalItem to dynamodbAtributeValue
func MarshalItem(item interface{}) (map[string]*dynamodb.AttributeValue, error) {
	// Marshal to Dynamodb item
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err.Error()))
	}

	return av, err
}

// InsertNewItem -
func InsertNewItem(item map[string]*dynamodb.AttributeValue, tableName string) (*dynamodb.PutItemOutput, error) {
	svc := dynamoDBClient()

	log.Printf("Putting new Item: %v\n", item)

	// PutItem in DynamoDB
	ok, err := svc.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	})

	if err != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
	}

	return ok, err
}
