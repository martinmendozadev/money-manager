package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"encoding/json"
	"fmt"
	"os"
)

// User struct
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	FistName string `json:"fistName"`
	LastName string `json:"lastName"`
}

// Handler function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) { //nolint:gocritic
	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// New UUID for user ID
	userUUID := uuid.New().String()

	fmt.Println("Generated new user UUID:", userUUID)

	// Unmarshal to User to access object properties
	userString := request.Body
	userStruct := User{}
	err := json.Unmarshal([]byte(userString), &userStruct)
	if err != nil {
		fmt.Println("Error Unmarshal userString: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	if userStruct.Email == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Create new struct of type user
	user := User{
		ID:       userUUID,
		Email:    userStruct.Email,
		FistName: userStruct.FistName,
		LastName: userStruct.LastName,
	}

	// Marshal to dynamodb item
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		fmt.Println("Error marshaling user: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	tableName := os.Getenv("DYNAMODB_TABLE")

	// Build put user input
	fmt.Printf("Putting user: %v\n", av)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// PutItem request in DynamoDB
	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	// Marshal user to return
	userMarshalled, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Got Marshaling object: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	fmt.Println("Returning user: ", string(userMarshalled))

	// Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{Body: string(userMarshalled), StatusCode: 201}, nil
}

func main() {
	lambda.Start(Handler)
}
