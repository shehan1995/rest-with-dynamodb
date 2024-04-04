package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"rest-with-dynamodb/config"
)

type DynamoHandle struct {
	Session client.ConfigProvider
	Handle  *dynamodb.DynamoDB
}

// CreateDynamoHandle creates an aws session for dynamo db
func CreateDynamoHandle(conf config.AWSConfig) (handle *DynamoHandle, err error) {

	dCfg := aws.NewConfig().WithCredentials(
		credentials.NewStaticCredentials(
			conf.AccessKey, conf.SecretKey, "",
		)).WithRegion(conf.Region)

	if conf.Endpoint != "" {
		dCfg = dCfg.WithEndpoint(conf.Endpoint)
	}

	awsSession, err := session.NewSession(dCfg)
	if err != nil {
		return nil, err
	}

	handle = &DynamoHandle{
		Session: awsSession,
		Handle:  dynamodb.New(awsSession),
	}
	return
}
