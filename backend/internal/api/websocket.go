package api

import (
	"github.com/KainoaGardner/csc/internal/auth"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/KainoaGardner/csc/internal/websockets"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) registerWebsocketRoutes(r chi.Router) {
	r.Get("/ws/{gameID}", h.connectToGame)
}

// auth
func (h *Handler) connectToGame(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "gameID")

	claims, statusCode, err := auth.CheckValidAuth(h.client, h.config.DB, h.config.JWT.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	websockets.AddPlayerToGame(gameID, claims.UserID, conn)

	go websockets.HandleMessages(gameID, claims.UserID, conn)
}
