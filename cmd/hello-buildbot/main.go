package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("HEADERS:")
	fmt.Printf("%v\n", request.Headers)
	fmt.Printf("\nCONTEXT:\n")
	fmt.Printf("%v\n", request.RequestContext)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "hello AWS lambda, this is a change",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
