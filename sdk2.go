package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

// Env var for AWS Region and Parameter name
var (
	reg, pn string
)

func getSSMClient() *ssm.Client {
	// Loading AWS Region
	reg = os.Getenv("AWS_REGION")
	// First we create a new AWS Config reading the credentials from the
	// environment (check ./.env)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(reg))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// With the AWS Config in hands, we create the SSM Client
	return ssm.NewFromConfig(cfg)
}

func getParameterName() string {
	// Reading values from environment variables
	// (see the ./.env file for more details).
	return os.Getenv("AWS_PARAMETER_NAME")
}

// AWS SSM DeleteParameter Example
// This is a demonstration of how to delete a single Parameter from AWS SSM.
func DeleteParameterExample() {
	// Getting the SSM Client
	c := getSSMClient()

	// Creating the DeleteParameterInput
	pi := &ssm.DeleteParameterInput{
		// The Parameter name
		Name: aws.String(getParameterName()),
	}

	// Then, request the DeleteParameter
	res, err := c.DeleteParameter(context.TODO(), pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("DeleteParameter: ", pn)
	fmt.Println(res)
}

// AWS SSM GetParameter Example
// This is a demonstration of how to read a single Parameter from AWS SSM.
func GetParameterExample() {
	// Getting the SSM Client
	c := getSSMClient()

	// Creating the GetParameterInput
	pi := &ssm.GetParameterInput{
		// The Parameter name
		Name: aws.String(getParameterName()),
		// Designed when reading a Parameter of type SecureString
		WithDecryption: true,
	}

	// Then, request the GetParameter
	res, err := c.GetParameter(context.TODO(), pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("GetParameter: ", pn)
	fmt.Println(res)
}

// AWS SSM GetParameters Example
// This is a demonstration of how to read a list of Parameters from AWS SSM.
func GetParametersExample() {
	// Getting the SSM Client
	c := getSSMClient()

	// Reading a list of Parameters from AWS Parameter Store:
	// First, we create a list of Parameter names
	n := []string{
		// A valid Parameter name
		getParameterName(),
		// An invalid Parameter name
		"InvalidParamName",
	}
	// Then, we create a GetParametersInput
	pi := &ssm.GetParametersInput{
		// The Parameter Names
		Names: n,
		// Designed when reading a Parameter of type SecureString
		WithDecryption: true,
	}
	// Then, request the GetParameters
	res, err := c.GetParameters(context.TODO(), pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("GetParameters: ", pn)
	fmt.Println(res)
}

// AWS SSM PutParameter Example
// This is a demonstration of how to insert or update a Parameter on AWS SSM.
func PutParameterExample() {
	// Getting the SSM Client
	c := getSSMClient()

	// Creating the GetParameterInput
	pi := &ssm.PutParameterInput{
		// The Parameter name
		Name: aws.String(getParameterName()),
		// The Parameter value
		Value: aws.String("Cash Rules Everything Around Me"),
		// Instruction to overwrite if the Parameter already exists
		Overwrite: true,
		// Parameter type
		Type: types.ParameterTypeSecureString,
	}

	// Then, request the GetParameter
	res, err := c.PutParameter(context.TODO(), pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("PutParameter: ", pn)
	fmt.Println(res)
}

func main() {
	PutParameterExample()
	GetParameterExample()
	GetParametersExample()
	DeleteParameterExample()
}
