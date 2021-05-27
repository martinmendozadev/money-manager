package utils

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

// SaveItem into dynamodb
func SaveItem(item map[string]*dynamodb.AttributeValue, tableName *string) (*dynamodb.PutItemOutput, error) {
	svc := dynamoDBClient()

	log.Printf("Putting new Item: %v\n", item)

	// PutItem in DynamoDB
	ok, err := svc.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(*tableName),
	})

	if err != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
	}

	return ok, err
}

// FindUserByEmail -
func FindUserByEmail(email, tableName *string) (*dynamodb.ScanOutput, error) {
	svc := dynamoDBClient()

	// Get all items with that email
	filt := expression.Name("email").Equal(expression.Value(*email))

	// Get back the id
	proj := expression.NamesList(expression.Name("id"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(*tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	return result, err
}
