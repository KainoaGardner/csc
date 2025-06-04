package api

import (
	// "github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/utils"
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
	r.Post("/game/move", h.movePiece)

	//use ID for each game access
	// r.Post("/game/id/create", h.getBoard)
	// r.Post("/game/id/move", h.getBoard)
}

func (h *Handler) getBoard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func (h *Handler) movePiece(w http.ResponseWriter, r *http.Request) {
	// err := utils.ParseJSON(r, &postHandScore)
	// if err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// returnHandScore, err := hands.GetHandScore(&postHandScore)
	// if err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	utils.WriteJSON(w, http.StatusOK, "test")
}
