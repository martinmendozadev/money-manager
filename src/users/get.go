package main

import (
	"net/http"
	"os"

	"github.com/martinmendozadev/money-manager/src/auth0"
	"github.com/martinmendozadev/money-manager/src/models"
	"github.com/martinmendozadev/money-manager/src/utils"
)

// GetUser -
func GetUser(request utils.Request) (utils.Response, error) {
	tableName := os.Getenv("DYNAMODB_TABLE")
	dbClient := utils.NewDBConnection(tableName)

	// Getting id from path parameters
	pathParamID := request.PathParameters["id"]

	// Found user by ID
	result, err := dbClient.GetItemByID(&pathParamID)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	// In case not found any user
	if result.Item == nil {
		return utils.Response{StatusCode: http.StatusNoContent}, err
	}

	user := models.User{}

	// UnmarshallMap result.item into user
	err = utils.UnmarshalItem(result.Item, &user)
	if err != nil {
		return utils.Response{StatusCode: http.StatusInternalServerError}, err
	}

	// Success response
	response, err := utils.NewResponse(http.StatusOK, user)

	return *response, nil
}

func main() { // nolint:typecheck
	auth0.Authorization(utils.Start(GetUser))
}
