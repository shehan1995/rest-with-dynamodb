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

	//load config
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config load error")
	}

	//create dynamodb handler
	dynamoDb, err := dynamodb.CreateDynamoHandle(conf.AWS)
	if err != nil {
		log.Fatal("DynamoDB Create handler error")
	}

	//create and start http server
	srv := server.NewServer(conf, dynamoDb)
	srv.Start(ctx)

}
