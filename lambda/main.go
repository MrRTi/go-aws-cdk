package main

import (
	"lambda-func/app"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return app.ApiHandler.RegisterUserHandler(request)
		default:
			return events.APIGatewayProxyResponse{
				Body: "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
