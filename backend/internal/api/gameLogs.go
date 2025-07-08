package api

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) registerGameLogRoutes(r chi.Router) {
	r.Get("/log", h.getAllGameLogs)
	r.Get("/log/{gameLogID}", h.getGameLog)
	r.Delete("/log", h.deleteAllGameLogs)
}

func (h *Handler) getAllGameLogs(w http.ResponseWriter, r *http.Request) {
	games, err := db.ListAllGameLogs(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ids := []string{}
	for _, game := range games {
		idString := game.ID.Hex()
		ids = append(ids, idString)
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("%d game logs found", len(ids)), ids)
}

func (h *Handler) deleteAllGameLogs(w http.ResponseWriter, r *http.Request) {
	amount, err := db.DeleteAllGameLogs(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{"count": amount}

	utils.WriteResponse(w, http.StatusOK, "Gamelogs deleted", data)
}

func (h *Handler) getGameLog(w http.ResponseWriter, r *http.Request) {
	gameLogID := chi.URLParam(r, "gameLogID")
	gameLog, err := db.FindGameLog(h.client, h.dbConfig, gameLogID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	data := map[string]interface{}{
		"id":          gameLog.ID.Hex(),
		"gameID":      gameLog.GameID,
		"whiteID":     gameLog.WhiteID,
		"blackID":     gameLog.BlackID,
		"date":        gameLog.Date,
		"moveCount":   gameLog.MoveCount,
		"moves":       gameLog.Moves,
		"boardStates": gameLog.BoardStates,
		"winner":      gameLog.Winner,
		"reason":      gameLog.Reason,
	}

	utils.WriteResponse(w, http.StatusOK, "Board", data)
}
