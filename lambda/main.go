package main

import (
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app := app.NewApp()
	lambda.Start(app.ApiHandler.RegisterUserHandler)
}
