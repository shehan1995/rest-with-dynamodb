package book

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"rest-with-dynamodb/entities"
	"rest-with-dynamodb/internal"
	"rest-with-dynamodb/usecases/book"
)

// AddBookHandler Insert Books to DynamoDB
func AddBookHandler(w http.ResponseWriter, r *http.Request) {

	data := &entities.AddBookRequest{}
	//convert request body to the entity
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, internal.ErrBadRequest)
		return
	}

	bookUseCase := book.NewUseCase()
	//create book item to insert to db
	bookItem := entities.BookItem{
		ISBNNumber:       data.ISBNNumber,
		SortKey:          "unique_sort", //this is to keep isbn unique
		Name:             data.Name,
		Author:           data.Author,
		PublishedVersion: data.PublishedVersion,
		CreatedAt:        time.Now().String(),
	}

	//call use case
	err := bookUseCase.AddBook(r.Context(), bookItem)
	if err != nil {
		render.Render(w, r, internal.ErrInternalServerError)
		return
	}

	//success response
	resp := entities.AddBookResponse{
		Status:  "Success",
		Message: "Created",
	}

	render.Render(w, r, resp)
}

// GetBookHandler GET Books from DynamoDB by ISBN
func GetBookHandler(w http.ResponseWriter, r *http.Request) {

	isbn := chi.URLParam(r, "isbn")
	if isbn == "" {
		render.Render(w, r, internal.ErrBadRequest)
		return
	}

	bookUseCase := book.NewUseCase()

	//call use case
	bookItem, err := bookUseCase.GetBook(r.Context(), isbn)
	if err != nil {
		if err.Error() == "NotFound" {
			render.Render(w, r, internal.ErrNotFound)
			return
		}
		render.Render(w, r, internal.ErrInternalServerError)
		return
	}
	resp := entities.GetBookResponse{
		ISBNNumber:       bookItem.ISBNNumber,
		Name:             bookItem.Name,
		Author:           bookItem.Author,
		PublishedVersion: bookItem.PublishedVersion,
	}

	render.Render(w, r, resp)
}
