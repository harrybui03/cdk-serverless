package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

type Repository struct {
	Db *dynamodb.DynamoDB
}

func GetDynamoDB() *dynamodb.DynamoDB {
	if os.Getenv("ENVIRONMENT") == "development" {
		sess := session.Must(session.NewSession(&aws.Config{
			Region:   aws.String("ap-southeast-1"),
			Endpoint: aws.String("http://dynamodb:8000"),
		}))
		return dynamodb.New(sess)
	}
	sess := session.Must(session.NewSession())
	return dynamodb.New(sess)
}
