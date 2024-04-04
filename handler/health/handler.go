package health

import (
	"net/http"

	"github.com/go-chi/render"
)

type healthResponse struct {
	Status bool `json:"status"`
}

func (h healthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Handler health check endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	health := healthResponse{Status: true}
	render.Render(w, r, health)
}
