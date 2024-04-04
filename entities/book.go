package entities

import (
	"net/http"
)

type AddBookRequest struct {
	ISBNNumber       string `json:"isbn_number"`
	Name             string `json:"name"`
	Author           string `json:"author"`
	PublishedVersion int    `json:"published_version"`
}

type AddBookResponse struct {
	Status  string
	Message string
}

type BookItem struct {
	ISBNNumber       string `json:"isbn"`
	SortKey          string `json:"sort"`
	Name             string `json:"name"`
	Author           string `json:"author"`
	PublishedVersion int    `json:"published_version"`
	CreatedAt        string `json:"created_at"`
}

type GetBookResponse struct {
	ISBNNumber       string `json:"isbn_number"`
	Name             string `json:"name"`
	Author           string `json:"author"`
	PublishedVersion int    `json:"published_version"`
}

func (b AddBookResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (b GetBookResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (b *AddBookRequest) Bind(r *http.Request) error {
	return nil
}
