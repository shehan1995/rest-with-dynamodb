package usecases

import (
	"context"
	"rest-with-dynamodb/entities"
)

type BookInterface interface {
	AddBook(ctx context.Context, bookItem entities.BookItem) error
	GetBook(ctx context.Context, isbn string) (entities.BookItem, error)
}
