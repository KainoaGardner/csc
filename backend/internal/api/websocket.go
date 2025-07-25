package api

import (
	"fmt"
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
	r.Get("/ws/{gameID}/{accessToken}", h.connectToGame)
}

// auth
func (h *Handler) connectToGame(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "gameID")
	accessToken := chi.URLParam(r, "accessToken")

	if gameID == "" || accessToken == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Missing gameID or accessToken"))
		return
	}

	claims, err := auth.ParseToken(h.config.JWT.AccessKey, accessToken)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if auth.CheckExpiredToken(claims) {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	websockets.AddPlayerToGame(gameID, claims.UserID, conn)

	go websockets.HandleMessages(gameID, claims.UserID, conn, h.client, h.config)
}
