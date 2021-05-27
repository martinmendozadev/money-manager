package utils

import "github.com/aws/aws-lambda-go/lambda"

// Start lambda handler
func Start(handler interface{}) {
	lambda.Start(handler)
}
