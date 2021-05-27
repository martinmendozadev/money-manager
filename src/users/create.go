package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/martinmendozadev/money-manager/src/models"
	"github.com/martinmendozadev/money-manager/src/utils"
)

// CreateUser -
func CreateUser(request utils.Request) (utils.Response, error) {
	// Unmarshal to access request object properties
	userString := request.Body
	userStruct := models.User{}
	err := json.Unmarshal([]byte(userString), &userStruct)
	if err != nil {
		log.Println("Error Unmarshal userString: ", err.Error())
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	userUUID := uuid.New().String()
	now := time.Now()
	user := models.User{
		ID:        userUUID,
		Email:     userStruct.Email,
		FirstName: userStruct.FirstName,
		LastName:  userStruct.LastName,
		CreatedAt: now.String(),
		UpdatedAt: now.String(),
	}

	tableName := os.Getenv("DYNAMODB_TABLE")
	dbClient := utils.NewDBConnection(tableName)

	// Search an user with the same email
	result, err := dbClient.FindUserByEmail(&user.Email)
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
	_, err = dbClient.SaveItem(av)
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
