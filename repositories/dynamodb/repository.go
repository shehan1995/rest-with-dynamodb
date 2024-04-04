package dynamodb

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"rest-with-dynamodb/repositories"
)

var DynamoDBRepo repositories.DynamoDBRepoInf

type Repository struct {
	handler *DynamoHandle
	table   *string
}

func NewRepository(handle *DynamoHandle, table string) *Repository {
	return &Repository{
		handler: handle,
		table:   &table,
	}
}

func (repo *Repository) AddItem(ctx context.Context, item interface{}) error {

	itemMap, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	in := &dynamodb.PutItemInput{
		TableName: repo.table,
		Item:      itemMap,
	}
	_, err = repo.handler.Handle.PutItem(in)
	if err != nil {
		return err
	}
	return err
}

// QueryItems queries items from DynamoDB
func (repo *Repository) QueryItem(ctx context.Context, key string, itemInf interface{}) error {
	// Define the input parameters
	input := &dynamodb.QueryInput{
		TableName: repo.table,
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":isbn": {S: &key},
		},
		KeyConditionExpression: aws.String("isbn = :isbn"),
	}

	// Perform the query
	resp, err := repo.handler.Handle.Query(input)
	if err != nil {
		return err
	}

	for _, item := range resp.Items {
		err = dynamodbattribute.UnmarshalMap(item, &itemInf)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("NotFound")
}
