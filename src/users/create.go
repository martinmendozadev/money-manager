package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

// CreateUser -
func CreateUser(ctx context.Context, request *events.APIGatewayProxyRequest) (utils.Response, error) {
	// Unmarshal to access request object properties
	userString := request.Body
	userStruct := User{}
	err := json.Unmarshal([]byte(userString), &userStruct)
	if err != nil {
		log.Println("Error Unmarshal userString: ", err.Error())
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	userUUID := uuid.New().String()
	now := time.Now()
	user := User{
		ID:        userUUID,
		Email:     userStruct.Email,
		FistName:  userStruct.FistName,
		LastName:  userStruct.LastName,
		CreatedAt: now.String(),
		UpdatedAt: now.String(),
	}

	tableName := os.Getenv("DYNAMODB_TABLE")

	// Marshal user to insert at DB
	av, err := utils.MarshalItem(user)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	// Insert Item in DynamoDB
	_, err = utils.InsertNewItem(av, tableName)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	// Success response
	response, err := utils.NewResponse(http.StatusCreated, user)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return *response, nil
}

func main() {
	lambda.Start(CreateUser)
}
