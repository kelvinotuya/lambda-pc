package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file (~/.aws/credentials).
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Fatal("Error creating AWS session: ", err)
	}

	svc := lambda.New(sess)
	result, err := svc.ListFunctions(&lambda.ListFunctionsInput{})
	if err != nil {
		log.Fatal("Error listing lambda functions: ", err)
	}

	// Get the list of lambda functions which have provisioned concurrency enabled
	var provisionedFn []string
	for _, fn := range result.Functions {
		provisionConc, err := svc.GetFunctionConcurrency(&lambda.GetFunctionConcurrencyInput{FunctionName: fn.FunctionName})
		if err != nil {
			log.Fatal("Error getting function concurrency: ", err)
		}
		if aws.Int64Value(provisionConc.ReservedConcurrentExecutions) > 0 {
			provisionedFn = append(provisionedFn, aws.StringValue(fn.FunctionName))
		}
	}

	fmt.Println("Lambda Functions with provisioned concurrency enabled:")
	for _, fn := range provisionedFn {
		fmt.Printf("\t%!s(MISSING)\n", fn)
	}
	os.Exit(0)
}
