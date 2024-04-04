package health

import (
	"github.com/go-chi/render"
	"net/http"
)

type healthResponse struct {
	Status bool `json:"status"`
}

func (h healthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	health := healthResponse{Status: true}
	render.Render(w, r, health)
}
