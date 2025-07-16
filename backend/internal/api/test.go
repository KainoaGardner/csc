package api

import (
	"github.com/KainoaGardner/csc/internal/auth"
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) registerTestRoutes(r chi.Router) {
	r.Get("/test/moveConvert", h.moveTest)
}

// admin
func (h *Handler) moveTest(w http.ResponseWriter, r *http.Request) {
	statusCode, err := auth.CheckAdminRequest(h.client, h.config.DB, h.config.JWT.AccessKey, r)
	if err != nil {
		utils.WriteError(w, statusCode, err)
		return
	}

	moves, err := engine.MovesTest()
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Failed", moves)
		return
	}

	utils.WriteResponse(w, http.StatusBadRequest, "Passed", moves)
}
