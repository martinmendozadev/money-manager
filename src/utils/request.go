package utils

import "github.com/aws/aws-lambda-go/events"

// Request wraps the base API Gateway request
type Request events.APIGatewayProxyResponse
