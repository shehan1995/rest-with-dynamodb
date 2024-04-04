package book

import (
	"context"

	"rest-with-dynamodb/entities"
	"rest-with-dynamodb/repositories"
	"rest-with-dynamodb/repositories/dynamodb"
)

type Usecase struct {
	DynamoDBRepo repositories.DynamoDBRepoInf
}

func NewUseCase() *Usecase {
	return &Usecase{DynamoDBRepo: dynamodb.DynamoDBRepo}
}

func (u *Usecase) AddBook(ctx context.Context, bookItem entities.BookItem) error {
	err := u.DynamoDBRepo.AddItem(ctx, bookItem)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) GetBook(ctx context.Context, isbn string) (book entities.BookItem, err error) {

	err = u.DynamoDBRepo.QueryItem(ctx, isbn, &book)
	if err != nil {
		return entities.BookItem{}, err
	}

	return book, nil
}
