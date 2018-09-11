package dynamoDB

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"fmt"
)

type Service interface {

}

type service struct {
	ddb *dynamodb.DynamoDB
}

func NewDynamoService() Service{
	s := &service{}
	s.ddb = s.DynamoDBSession()
	return s
}

func (s *service)DynamoDBSession() *dynamodb.DynamoDB{
	aws_access_key_id := ""
	aws_secret_access_key := ""
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		log.Println(err)
	}
	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	svc := dynamodb.New(session.New(), cfg)
	return svc
}

func (s *service)DynamoDBUpdate() {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String("0.5"),
			},
		},
		TableName: aws.String("receipts"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String("456456325345"),
			},
			"link": {
				S: aws.String("The Big New Movie"),
			},
			"data": {
				S: aws.String("Some kak data"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set info.rating = :r"),
	}

		return
	}
