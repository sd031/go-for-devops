package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, req Request) (Response, error) {
	return Response{
		Message: "Hello " + req.Name,
	}, nil
}

// main code block

func main() {
	lambda.Start(handler)
}
