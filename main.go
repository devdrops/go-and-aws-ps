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

func main() {
	// Reading values from environment variables
	// (see the ./.env file for more details).
	reg = os.Getenv("AWS_REGION")
	aki = os.Getenv("AWS_ACCESS_KEY_ID")
	sak = os.Getenv("AWS_SECRET_ACCESS_KEY")
	st = os.Getenv("AWS_SESSION_TOKEN")
	pn = os.Getenv("AWS_PARAMETER_NAME")

	// Creating a new AWS Config
	c := &aws.Config{
		Credentials: credentials.NewStaticCredentials(aki, sak, st),
		Region:      aws.String(reg),
	}

	// Creating a new AWS Session
	sess := session.Must(session.NewSession())

	// Creating a new AWS SSM Service Client
	sc := ssm.New(sess, c)

	// Reading a Parameter from AWS Parameter Store:
	// First, we create a GetParameterInput
	pi := &ssm.GetParameterInput{
		// The Parameter Name
		Name: aws.String(pn),
		// To get the Parameter value decrypted, when using SecureString
		WithDecryption: aws.Bool(true),
	}
	// Then, we request GetParameter
	out, err := sc.GetParameter(pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(out)
}
