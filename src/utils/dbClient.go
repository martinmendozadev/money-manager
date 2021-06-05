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

// DBClient connection struct
type DBClient struct {
	tableName string
}

var svc = dynamoDBClient()

// NewDBConnection -
func NewDBConnection(tableName string) *DBClient {
	return &DBClient{
		tableName: tableName,
	}
}

// create a new session to conect with Dynamodb
func dynamoDBClient() *dynamodb.DynamoDB {
	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	return dynamodb.New(sess)
}

// MarshalItem as a dynamodbAtributeValue
func MarshalItem(item interface{}) (map[string]*dynamodb.AttributeValue, error) {
	// Marshal to Dynamodb item
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err.Error()))
	}

	return av, err
}

// UnmarshalItem -
func UnmarshalItem(result map[string]*dynamodb.AttributeValue, item interface{}) error {
	err := dynamodbattribute.UnmarshalMap(result, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return err
}

// SaveItem into Dynamodb
func (db *DBClient) SaveItem(item map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	ok, err := svc.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(db.tableName),
	})

	if err != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
	}

	return ok, err
}

// GetItemByID -
func (db *DBClient) GetItemByID(id *string) (*dynamodb.GetItemOutput, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(*id),
			},
		},
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	return result, err
}

// FindUserByEmail at DynamoDB
func (db *DBClient) FindUserByEmail(email *string) (*dynamodb.ScanOutput, error) {
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
		TableName:                 aws.String(db.tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	return result, err
}
