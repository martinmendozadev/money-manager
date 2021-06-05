package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/martinmendozadev/money-manager/src/models"
	"github.com/martinmendozadev/money-manager/src/utils"
)

// GetUser -
func GetUser(request utils.Request) (utils.Response, error) {
	tableName := os.Getenv("DYNAMODB_TABLE")
	dbClient := utils.NewDBConnection(tableName)

	// Getting id from path parameters
	//pathParamID := request.PathParameters["id"]
	pathParamID := "c1abfedc-02a1-4cd4-94e4-2eea47496b1d"

	result, err := dbClient.GetItemByID(&pathParamID)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	if len(result.Item) == 0 {
		return utils.Response{StatusCode: http.StatusNoContent}, err
	}

	user := models.User{}

	err = utils.UnMarshalItem(result.Item, &user)
	if err != nil {
		panic(fmt.Sprintf("Failed to UnmarshalMap result.Item: %v\n", err))
	}

	marshalledUser, err := json.Marshal(user)

	// Success response
	response, err := utils.NewResponse(http.StatusOK, marshalledUser)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return *response, nil
}

func main() {
	utils.Start(GetUser)
}
