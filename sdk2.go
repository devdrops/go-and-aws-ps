package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

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
	cfg, err := config.LoadDefaultConfig(context.TODO())
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
	fmt.Println(Prettify(res))
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
	fmt.Println(Prettify(res))
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
	fmt.Println(Prettify(res))
}

// AWS SSM GetParametersByPath Example
// This is a demonstration of how to read a list of Parameters by it's Path
// hierarchy from AWS SSM.
func GetParametersByPathExample() {
	// Getting the SSM Client
	c := getSSMClient()

	// We can use a list of filters for this task
	filters := []types.ParameterStringFilter{
		types.ParameterStringFilter{
			// The Key we can use to determine the filter. It must
			// create a relationship with Values.
			Key: aws.String("Type"),
			// The values we can use to apply the filter. It must
			// match the relationship with Key
			Values: []string{"SecureString"},
		},
	}

	// Then, we create a GetParametersInput
	pi := &ssm.GetParametersByPathInput{
		// The Parameter Path
		Path: aws.String("/"),
		// The maximum number of results
		MaxResults: 2,
		// Read Parameters by it's Path hierarchy
		Recursive: true,
		// Filters
		ParameterFilters: filters,
		// Designed when reading Parameters of type SecureString
		WithDecryption: true,
	}
	// Then, request the GetParameters
	res, err := c.GetParametersByPath(context.TODO(), pi)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("GetParametersByPath:", pn)
	fmt.Println(Prettify(res))
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
	fmt.Println(Prettify(res))
}

func main() {
	PutParameterExample()
	GetParameterExample()
	GetParametersExample()
	GetParametersByPathExample()
	DeleteParameterExample()
}

/* Parts of code copied from the original AWS SDK, just to make Print better :) */

// Prettify returns the string representation of a value.
func Prettify(i interface{}) string {
	var buf bytes.Buffer
	prettify(reflect.ValueOf(i), 0, &buf)
	return buf.String()
}

// prettify will recursively walk value v to build a textual
// representation of the value.
func prettify(v reflect.Value, indent int, buf *bytes.Buffer) {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		strtype := v.Type().String()
		if strtype == "time.Time" {
			fmt.Fprintf(buf, "%s", v.Interface())
			break
		} else if strings.HasPrefix(strtype, "io.") {
			buf.WriteString("<buffer>")
			break
		}

		buf.WriteString("{\n")

		names := []string{}
		for i := 0; i < v.Type().NumField(); i++ {
			name := v.Type().Field(i).Name
			f := v.Field(i)
			if name[0:1] == strings.ToLower(name[0:1]) {
				continue // ignore unexported fields
			}
			if (f.Kind() == reflect.Ptr || f.Kind() == reflect.Slice || f.Kind() == reflect.Map) && f.IsNil() {
				continue // ignore unset fields
			}
			names = append(names, name)
		}

		for i, n := range names {
			val := v.FieldByName(n)
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(n + ": ")
			prettify(val, indent+2, buf)

			if i < len(names)-1 {
				buf.WriteString(",\n")
			}
		}

		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")
	case reflect.Slice:
		strtype := v.Type().String()
		if strtype == "[]uint8" {
			fmt.Fprintf(buf, "<binary> len %d", v.Len())
			break
		}

		nl, id, id2 := "", "", ""
		if v.Len() > 3 {
			nl, id, id2 = "\n", strings.Repeat(" ", indent), strings.Repeat(" ", indent+2)
		}
		buf.WriteString("[" + nl)
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(id2)
			prettify(v.Index(i), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString("," + nl)
			}
		}

		buf.WriteString(nl + id + "]")
	case reflect.Map:
		buf.WriteString("{\n")

		for i, k := range v.MapKeys() {
			buf.WriteString(strings.Repeat(" ", indent+2))
			buf.WriteString(k.String() + ": ")
			prettify(v.MapIndex(k), indent+2, buf)

			if i < v.Len()-1 {
				buf.WriteString(",\n")
			}
		}

		buf.WriteString("\n" + strings.Repeat(" ", indent) + "}")
	default:
		if !v.IsValid() {
			fmt.Fprint(buf, "<invalid value>")
			return
		}
		format := "%v"
		switch v.Interface().(type) {
		case string:
			format = "%q"
		case io.ReadSeeker, io.Reader:
			format = "buffer(%p)"
		}
		fmt.Fprintf(buf, format, v.Interface())
	}
}
