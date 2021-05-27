package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/martinmendozadev/money-manager/src/utils"
)

// CreateUser -
func CreateUser(request utils.Request) (utils.Response, error) {
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
		FirstName: userStruct.FirstName,
		LastName:  userStruct.LastName,
		CreatedAt: now.String(),
		UpdatedAt: now.String(),
	}

	tableName := os.Getenv("DYNAMODB_TABLE")

	// Search an user with the same email
	result, err := utils.FindUserByEmail(&user.Email, &tableName)
	if *result.Count > 0 {
		return utils.Response{StatusCode: http.StatusConflict}, err
	} else if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	// Marshal user to insert at DB
	av, err := utils.MarshalItem(user)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	// Insert Item in DynamoDB
	_, err = utils.InsertNewItem(av, &tableName)
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
	utils.Start(CreateUser)
}
