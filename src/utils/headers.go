package utils

import (
	"encoding/json"

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
	}

	res := &Response{
		StatusCode:      statusCode,
		Body:            string(objMarshalled),
		IsBase64Encoded: false,
	}

	res.WithCors()

	return res, nil
}

// SetHeaders define response headers
func (r *Response) SetHeaders(h map[string]string) {
	r.Headers = h
}

// WithCors add cors support to API Gatway response
func (r *Response) WithCors() {
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}

	// Warning: Upadate origon to production
	r.Headers["Access-Control-Allow-Origin"] = "localhost"
	r.Headers["Content-Type"] = "application/json"
}
