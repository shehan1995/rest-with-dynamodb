package main

import (
	"context"
	"log"
	"rest-with-dynamodb/config"
	"rest-with-dynamodb/repositories/dynamodb"
	"rest-with-dynamodb/server"
)

func main() {
	ctx := context.Background()
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config load error")
	}

	dynamoDb, err := dynamodb.CreateDynamoHandle(conf.AWS)
	if err != nil {
		log.Fatal("DynamoDB Create handler error")
	}

	srv := server.NewServer(conf, dynamoDb)
	srv.Start(ctx)

}
