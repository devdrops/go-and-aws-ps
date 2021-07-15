package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Env vars for configuration:
// Region, Access Key ID, Secret Access Key and Session Token.
// Also, the Parameter's Name that we're gonna read from AWS SSM.
var (
	reg, aki, sak, st, pn string
)

func getSSMClient() *ssm.SSM {
	// Reading values from environment variables
	// (see the ./.env file for more details).
	reg = os.Getenv("AWS_REGION")
	aki = os.Getenv("AWS_ACCESS_KEY_ID")
	sak = os.Getenv("AWS_SECRET_ACCESS_KEY")
	st = os.Getenv("AWS_SESSION_TOKEN")

	// Creating a new AWS Config
	c := &aws.Config{
		Credentials: credentials.NewStaticCredentials(aki, sak, st),
		Region:      aws.String(reg),
	}

	// Creating a new AWS Session
	sess := session.Must(session.NewSession())

	// Creating a new AWS SSM Service Client
	return ssm.New(sess, c)
}

func getParameterName() string {
	// Reading values from environment variables
	// (see the ./.env file for more details).
	return os.Getenv("AWS_PARAMETER_NAME")
}

// AWS SSM DeleteParameter Example
// This is a demonstration of how to delete a single Parameter from AWS SSM.
func DeleteParameterExample() {
	// AWS SSM Client
	sc := getSSMClient()

	// Deleting a Parameter from AWS Parameter Store:
	// First, we create a DeleteParameterInput
	pi := &ssm.DeleteParameterInput{
		// The Parameter Name
		Name: aws.String(getParameterName()),
	}
	// Then, we request GetParameter
	out, err := sc.DeleteParameter(pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("DeleteParameter")
	fmt.Println(out)
}

// AWS SSM GetParameter Example
// This is a demonstration of how to read a single Parameter from AWS SSM.
func GetParameterExample() {
	// AWS SSM Client
	sc := getSSMClient()

	// Reading a Parameter from AWS Parameter Store:
	// First, we create a GetParameterInput
	pi := &ssm.GetParameterInput{
		// The Parameter Name
		Name: aws.String(getParameterName()),
		// To get the Parameter value decrypted, when using SecureString
		WithDecryption: aws.Bool(true),
	}
	// Then, we request GetParameter
	out, err := sc.GetParameter(pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("GetParameter:", pn)
	fmt.Println(out)
}

// AWS SSM GetParameters Example
// This is a demonstration of how to read a list of Parameters from AWS SSM.
func GetParametersExample() {
	// AWS SSM Client
	sc := getSSMClient()

	// Reading a list of Parameters from AWS Parameter Store:
	// First, we create a list of Parameter names
	n := []*string{
		// A valid Parameter name
		aws.String(getParameterName()),
		// An invalid Parameter name
		aws.String("InvalidParamName"),
	}
	// Then, we create a GetParametersInput
	pi := &ssm.GetParametersInput{
		// The Parameter Names
		Names: n,
		// To get the Parameter values decrypted, when using SecureString
		WithDecryption: aws.Bool(true),
	}
	// And then, we request GetParameters
	out, err := sc.GetParameters(pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("GetParameters:", pn)
	fmt.Println(out)
}

// AWS SSM PutParameter Example
// This is a demonstration of how to insert or update a Parameter on AWS SSM.
func PutParameterExample() {
	// AWS SSM Client
	sc := getSSMClient()

	// Inserting or updating a Parameter on AWS Parameter Store:
	// First, we create a PutParameterInput
	pi := &ssm.PutParameterInput{
		// Parameter name
		Name: aws.String(getParameterName()),
		// Parameter value
		Value: aws.String("Cash Rules Everything Around Me"),
		// Overwrite if a Parameter already exists with the given Name
		Overwrite: aws.Bool(true),
		// Parameter type
		Type: aws.String("SecureString"),
	}
	// And then, we request PutParameter
	out, err := sc.PutParameter(pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("PutParameter:", pn)
	fmt.Println(out)
}

func main() {
	PutParameterExample()
	GetParameterExample()
	GetParametersExample()
	DeleteParameterExample()
}
