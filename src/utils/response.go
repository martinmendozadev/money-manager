package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// Response wraps the base API Gateway response
type Response events.APIGatewayProxyResponse

// NewResponse -
func NewResponse(code int, body interface{}) (*Response, error) {
	statusCode := code

	objMarshalled, err := json.Marshal(body)
	if err != nil {
		statusCode = 500
		log.Println("Got error marshaling object: ", err.Error())
	}

	res := &Response{
		StatusCode:      statusCode,
		Body:            string(objMarshalled),
		IsBase64Encoded: false,
	}

	res.setHeaders()

	return res, nil
}

// setHeaders add standard support headers to an API Gatway response
func (response *Response) setHeaders() {
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	if response.Headers == nil {
		response.Headers = map[string]string{}
	}

	response.Headers["Access-Control-Allow-Origin"] = allowedOrigins
	response.Headers["Content-Type"] = "application/json"
}
