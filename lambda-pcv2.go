package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func main() {
	// Load session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create lambda client
	svc := lambda.New(sess)

	// List functions
	result, err := svc.ListFunctions(&lambda.ListFunctionsInput{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Iterate over functions and print out the ones with provisioned concurrency
	for _, fun := range result.Functions {
		res, err := svc.GetFunction(&lambda.GetFunctionInput{FunctionName: fun.FunctionName})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if res.Concurrency.ReservedConcurrentExecutions != nil {
			fmt.Println(*res.Configuration.FunctionName)
		}
	}
}
