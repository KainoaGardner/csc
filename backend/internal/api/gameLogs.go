package api

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) registerGameLogRoutes(r chi.Router) {
	r.Get("/log", h.getAllGameLogs)
	r.Get("/log/{gameID}", h.getGameLog)
	r.Delete("/log", h.deleteAllGameLogs)
}

func (h *Handler) getAllGameLogs(w http.ResponseWriter, r *http.Request) {
	games, err := db.ListAllGames(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ids := []string{}
	for _, game := range games {
		idString := game.ID.Hex()
		ids = append(ids, idString)
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("%d games found", len(ids)), ids)
}

func (h *Handler) deleteAllGameLogs(w http.ResponseWriter, r *http.Request) {
	amount, err := db.DeleteAllGames(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"count": amount}

	utils.WriteResponse(w, http.StatusOK, "Games deleted", data)
}

func (h *Handler) getGameLog(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "gameID")
	game, err := db.FindGame(h.client, h.dbConfig, gameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	result, err := engine.ConvertBoardToString(*game)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"fen": result,
	}

	utils.WriteResponse(w, http.StatusOK, "Board", data)
}
