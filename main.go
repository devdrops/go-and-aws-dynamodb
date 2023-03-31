package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	reg, aki, sak, st string
)

func getClient() *dynamodb.DynamoDB {
	reg = os.Getenv("AWS_REGION")
	aki = os.Getenv("AWS_ACCESS_KEY_ID")
	sak = os.Getenv("AWS_SECRET_ACCESS_KEY")
	st = os.Getenv("AWS_SESSION_TOKEN")

	c := &aws.Config{
		Credentials: credentials.NewStaticCredentials(aki, sak, st),
		Region:      aws.String(reg),
	}

	sess := session.Must(session.NewSession())

	return dynamodb.New(sess, c)
}

func main() {
	fmt.Println("AAAAA")
}
