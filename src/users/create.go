package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"github.com/martinmendozadev/money-manager/src/utils"
)

// User struct
type User struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email"`
	FistName  string `json:"fistName"`
	LastName  string `json:"lastName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Handler function
func CreateUser(ctx context.Context, request events.APIGatewayProxyRequest) (utils.Response, error) { // nolint:gocritic
	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	userUUID := uuid.New().String()

	// Unmarshal to access object properties
	userString := request.Body
	userStruct := User{}
	err := json.Unmarshal([]byte(userString), &userStruct)
	if err != nil {
		fmt.Println("Error Unmarshal userString: ", err.Error())
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	if userStruct.Email == "" {
		return utils.Response{StatusCode: http.StatusBadRequest}, err
	}

	// Creating a timestamp
	now := time.Now()

	// Create new struct of type user
	user := User{
		ID:        userUUID,
		Email:     userStruct.Email,
		FistName:  userStruct.FistName,
		LastName:  userStruct.LastName,
		CreatedAt: now.String(),
		UpdatedAt: now.String(),
	}

	// Marshal to Dynamodb item
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		fmt.Println("Error marshaling user: ", err.Error())
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
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
		fmt.Println("Got an error inserting a User: ", err.Error())
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	// Build standard http response
	response, err := utils.NewResponse(http.StatusCreated, user)
	if err != nil {
		fmt.Println("Got error using utils: ", err.Error())
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return *response, nil
}

func main() {
	lambda.Start(CreateUser)
}
