package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dosorio55/gambitApi/awsgo"
	"github.com/dosorio55/gambitApi/db"
	"os"
	"strings"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InitialiceAws()

	if !ValidateParams() {
		panic("Missing parameters for lambda function")
	}

	var res *events.APIGatewayProxyResponse

	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	db.ReadSecret()

	headerResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headerResp,
	}

	return res, nil
}

func ValidateParams() bool {
	_, ok := os.LookupEnv("SecretName")
	if !ok {
		return false
	}

	// _, ok = os.LookupEnv("UserPoolId")
	// if !ok {
	// 	return false
	// }

	// _, ok = os.LookupEnv("Region")
	// if !ok {
	// 	return false
	// }

	_, ok = os.LookupEnv("UrlPrefix")
	if !ok {
		return false
	}

	return ok
}
