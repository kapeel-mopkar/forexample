package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct{
	ID float64
	Value string
}

type Response struct{
	Message string
	Ok bool
}

func Handler(request Request) (Response, error){
	return Response {
		Message: fmt.Sprintf("Process request ID %f", request.ID),
		Ok: true,
	}, nil
}

func main(){
	lambda.Start(Handler)
}