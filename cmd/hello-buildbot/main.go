package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("TEST_ENV = %s", os.Getenv("TEST_ENV")))

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       b.String(),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
