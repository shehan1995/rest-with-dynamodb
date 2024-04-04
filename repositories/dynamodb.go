package repositories

import "context"

type DynamoDBRepoInf interface {
	AddItem(ctx context.Context, item interface{}) error
	QueryItem(ctx context.Context, key string, itemInf interface{}) error
}
