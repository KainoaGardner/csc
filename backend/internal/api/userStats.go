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

func (h *Handler) registerUserStatRoutes(r chi.Router) {
	r.Get("/userStat/{userID}", h.getUserStats)
	r.Get("/userStat", h.getAllUserStats)
}

func (h *Handler) getUserStats(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	userStats, err := db.FindUserStatsFromUserID(h.client, h.dbConfig, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "User Stats", userStats)
}

// admin
func (h *Handler) getAllUserStats(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.dbConfig, h.jwt.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	userStats, err := db.ListAllUserStats(h.client, h.dbConfig)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	result := []types.UserStats{}
	for _, dbUserStat := range userStats {
		result = append(result, dbUserStat)
	}

	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("%d users found", len(result)), result)
}
