package api

import (
	// "github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/game/board", h.getBoard)
}

func (h *Handler) getBoard(w http.ResponseWriter, r *http.Request) {
}
