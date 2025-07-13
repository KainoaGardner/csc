package api

import (
	"fmt"
	"github.com/KainoaGardner/csc/internal/auth"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) registerGameLogRoutes(r chi.Router) {
	r.Get("/log", h.getAllGameLogs)
	r.Get("/log/{gameLogID}", h.getGameLog)
	r.Delete("/log", h.deleteAllGameLogs)
}

// admin
func (h *Handler) getAllGameLogs(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	gameLogs, err := db.ListAllGameLogs(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	result := []types.GameLog{}
	for _, gameLog := range gameLogs {
		result = append(result, gameLog)
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("%d game logs found", len(result)), result)
}

// admin
func (h *Handler) deleteAllGameLogs(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

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

	utils.WriteResponse(w, http.StatusOK, "Board", gameLog)
}
